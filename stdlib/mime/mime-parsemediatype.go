package main

import (
	"fmt"
	"mime"
)

func main() {

	for i, e := range []string{
		// #0: mediaType=text/plain, param=map[charset:UTF-8], err=<nil>
		"text/plain;charset=UTF-8",
		// メディアタイプは小文字に正規化される
		// #1: mediaType=text/plain, param=map[charset:UTF-8 hello:yo wao:hi], err=<nil>
		"Text/PLAIN ; charset=UTF-8 ; wao = hi;hello=yo",
		// #2: mediaType=text/plain, param=map[], err=<nil>
		"text/plain",
		// #3: mediaType=text, param=map[], err=<nil>
		"text",
	} {
		mt, param, err := mime.ParseMediaType(e)
		fmt.Printf("#%d: mediaType=%s, param=%s, err=%v\n", i, mt, param, err)
	}
}
