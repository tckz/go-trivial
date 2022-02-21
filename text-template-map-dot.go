package main

import (
	"os"
	"text/template"
)

func main() {
	// keyにdotやhyphenが含まれる場合は組込み関数の「index」で回避できる
	// 関数の引数に他の関数の返り値を使う場合はカッコで入れ子にできる
	t, err := template.New("message").Parse(`
map with key includes dot {{index (index . "param-hyphen") "key.dot"}}
map with key includes dot {{(index . "param-hyphen").key}}
`)
	if err != nil {
		panic(err)
	}

	// map with key includes dot value2
	// map with key includes dot value1
	err = t.Execute(os.Stdout, map[string]interface{}{
		"param-hyphen": map[string]interface{}{
			"key":     "value1",
			"key.dot": "value2",
		},
	})
	if err != nil {
		panic(err)
	}
}
