package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"time"
)

/*
httpサーバーへリクエストを投げた外部クライアントが応答を待たずに切断した場合。

リクエストのハンドラから別の外部http APIを叩いていてContextを正しく引き回していると、
この外部http API呼び出しもcontext canceledになる。

という状態の再現。

このコードでは外部クライアントとリクエストハンドラがcancel関数を共有し協調して実現しているが
ユーザーのブラウザ（特にモバイル）が応答をまたずにいなくなるケースは十分発生しうる
*/

func main() {
	var testURL string

	ctx, cancelExternalClient := context.WithCancel(context.Background())
	defer cancelExternalClient()

	m := &http.ServeMux{}
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "root handler begin\n")
		defer func() {
			fmt.Fprintf(os.Stderr, "root handler done\n")
		}()

		ctx := r.Context()

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		u, err := url.Parse(testURL)
		if err != nil {
			panic(err)
		}

		u.Path = "/someapi"
		fmt.Fprintf(os.Stderr, "someapi URL: %s\n", u)
		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			panic(err)
		}
		req = req.WithContext(ctx)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			contextCanceled := false
			if ue, ok := err.(*url.Error); ok {
				contextCanceled = ue.Err == context.Canceled
			}
			fmt.Fprintf(os.Stderr, "api-client: %T, %v, contextCanceled?=%t\n", err, err, contextCanceled)
			return
		}
		defer res.Body.Close()
	})

	// 外部http APIにあたるもの
	m.HandleFunc("/someapi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "someapi begin\n")
		defer func() {
			fmt.Fprintf(os.Stderr, "someapi done\n")
		}()

		// 外部クライアントの切断模擬
		cancelExternalClient()
		time.Sleep(100 * time.Millisecond)
	})

	testServer := httptest.NewServer(m)
	defer testServer.Close()

	testURL = testServer.URL

	req, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		panic(err)
	}

	req = req.WithContext(ctx)

	fmt.Fprintf(os.Stderr, "external-client.Do\n")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "external-client: %v\n", err)
		return
	}
	defer res.Body.Close()

	fmt.Fprintf(os.Stderr, "external-client done\n")
}
