package main

import (
	"fmt"
	"net/url"
)

func print(index int, u *url.URL, err error) {
	if err == nil {
		fmt.Printf("url[%d]=%s, scheme=%s, host=%s, path=%s, query=%v\n", index, u, u.Scheme, u.Host, u.Path, u.Query())
	} else {
		fmt.Printf("url[%d]=%s, scheme=%s, host=%s, path=%s, err=%v\n", index, u, u.Scheme, u.Host, u.Path, err)
	}
}

func main() {
	// スキームやauthority,pathのないURLをParseするとどうなるか確認するもの

	// url[0]=param=xx, scheme=, host=, path=param=xx, query=map[]
	u, err := url.Parse("param=xx")
	print(0, u, err)

	// url[1]=?param=xx, scheme=, host=, path=, query=map[param:[xx]]
	u, err = url.Parse("?param=xx")
	print(1, u, err)

	// url[2]=, scheme=, host=, path=, query=map[]
	u, err = url.Parse("")
	print(2, u, err)

	// url[3]=https://www.example.jp, scheme=https, host=www.example.jp, path=, query=map[]
	u, err = url.Parse("https://www.example.jp")
	print(3, u, err)

	// url[4]=https://www.example.jp/, scheme=https, host=www.example.jp, path=/, query=map[]
	u, err = url.Parse("https://www.example.jp/")
	print(4, u, err)

	// url[5]=https://www.example.jp/path/to, scheme=https, host=www.example.jp, path=/path/to, query=map[]
	u, err = url.Parse("https://www.example.jp/path/to")
	print(5, u, err)

	// url[6]=/withouthost?param=x, scheme=, host=, path=/withouthost, query=map[param:[x]]
	u, err = url.Parse("/withouthost?param=x")
	print(6, u, err)
}
