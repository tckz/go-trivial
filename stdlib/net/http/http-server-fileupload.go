package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
)

// ファイルアップロードを受け付けて、受け取ったファイルの情報をstdoutに出力するサーバー
//
// 起動:
//   go run http-server-fileupload.go [-addr :8080] [-max-memory 32768]
//
// multipart/form-data (curl -F):
//   curl -F file=@foo.bin -F other=@bar.txt http://localhost:8080/upload
//
// 生のボディをそのまま送る (curl -T / --data-binary):
//   curl -T foo.bin http://localhost:8080/upload
//   curl --data-binary @foo.bin -H 'Content-Type: application/octet-stream' http://localhost:8080/upload

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Int64Var(&maxMemory, "max-memory", 32<<10, "ParseMultipartForm: この量を超えたファイルパートは一時ファイルに書かれる (bytes)")
	flag.Parse()

	m := http.NewServeMux()
	m.HandleFunc("/upload", handleUpload)
	m.HandleFunc("/upload/", handleUpload) // curl -T は /upload/<filename> に PUT してくる
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "POST/PUT to /upload\n")
	})

	log.Printf("listening on %s", *addr)
	if err := http.ListenAndServe(*addr, m); err != nil {
		log.Fatalf("*** http.ListenAndServe: %v", err)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	ct := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(ct)
	if err != nil && ct != "" {
		http.Error(w, fmt.Sprintf("bad Content-Type: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("=== %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)

	if strings.HasPrefix(mediaType, "multipart/") {
		if err := handleMultipart(w, r); err != nil {
			log.Printf("*** multipart: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// multipartでなければボディ全体を1ファイルとみなす
	info, err := readFile(r.Body)
	if err != nil {
		log.Printf("*** read body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// curl -T の場合はファイル名がパスに乗ってくることがある(/upload/foo.bin)
	info.FileName = strings.TrimPrefix(strings.TrimPrefix(r.URL.Path, "/upload"), "/")
	info.ContentType = ct
	info.ContentLength = r.ContentLength
	info.print()

	fmt.Fprintf(w, "ok: %d bytes\n", info.Size)
}

// maxMemory は ParseMultipartForm に渡すメモリ上限。
// ファイルパートの合計がこれを超えると、超えた分は os.CreateTemp で一時ファイルに書かれる。
var maxMemory int64

func handleMultipart(w http.ResponseWriter, r *http.Request) error {
	// ParseMultipartForm はボディ全体を読み切って r.MultipartForm に展開する。
	// ファイル以外のフィールドは r.MultipartForm.Value (と r.PostForm / r.Form) に、
	// ファイルは r.MultipartForm.File に入る。
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		return fmt.Errorf("ParseMultipartForm: %w", err)
	}
	// 一時ファイルに書かれた分を消す。呼ばないと一時ファイルが残る。
	// とAIは書いているが実際は必要ない
	//defer r.MultipartForm.RemoveAll()

	for name, values := range r.MultipartForm.Value {
		fmt.Printf("  value[%s]: %q\n", name, values)
	}

	var total int64
	for name, fhs := range r.MultipartForm.File {
		for i, fh := range fhs {
			f, err := fh.Open()
			if err != nil {
				return fmt.Errorf("file[%s][%d] Open: %w", name, i, err)
			}
			additionalInfo := ""
			if fp, ok := f.(*os.File); ok {
				additionalInfo = fmt.Sprintf(", tmpfile=%s)", fp.Name())
			}
			// fh.Open() が返すのはメモリ上なら *sectionReadCloser、一時ファイルなら *os.File
			fmt.Printf("  file[%s][%d]: backed by %T%s\n", name, i, f, additionalInfo)

			info, err := readFile(f)
			f.Close()
			if err != nil {
				return fmt.Errorf("file[%s][%d] read: %w", name, i, err)
			}
			info.FieldName = name
			info.FileName = fh.Filename
			info.ContentType = fh.Header.Get("Content-Type")
			// FileHeader.Size はパース時に確定しているので読まなくても分かる
			info.ContentLength = fh.Size
			info.print()

			total += info.Size
		}
	}

	fmt.Fprintf(w, "ok: %d bytes\n", total)
	return nil
}

type fileInfo struct {
	FieldName     string
	FileName      string
	ContentType   string
	ContentLength int64
	Size          int64
	SHA256        string
}

// readFile はrを最後まで読んでサイズとハッシュを求める。内容はどこにも保存しない。
func readFile(r io.Reader) (*fileInfo, error) {
	h := sha256.New()
	n, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	return &fileInfo{
		Size:   n,
		SHA256: hex.EncodeToString(h.Sum(nil)),
	}, nil
}

func (fi *fileInfo) print() {
	w := os.Stdout
	if fi.FieldName != "" {
		fmt.Fprintf(w, "  field:          %s\n", fi.FieldName)
	}
	fmt.Fprintf(w, "  filename:       %q\n", fi.FileName)
	fmt.Fprintf(w, "  content-type:   %s\n", fi.ContentType)
	if fi.ContentLength >= 0 {
		fmt.Fprintf(w, "  content-length: %d\n", fi.ContentLength)
	}
	fmt.Fprintf(w, "  size:           %d bytes\n", fi.Size)
	fmt.Fprintf(w, "  sha256:         %s\n", fi.SHA256)
	fmt.Fprintln(w)
}
