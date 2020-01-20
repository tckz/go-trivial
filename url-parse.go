package main

import (
	"fmt"
	"net/url"
)

func print(index int, u *url.URL, err error) {
	if err == nil {
		fmt.Printf("url[%d]=%v, path=%s, query=%v\n", index, u, u.Path, u.Query())
	} else {
		fmt.Printf("url[%d]=%v, path=%s, err=%v\n", index, u, u.Path, err)
	}
}

func main() {
	// スキームやauthority,pathのないURLをParseするとどうなるか確認するもの

	u, err := url.Parse("param=xx")
	print(0, u, err)

	u, err = url.Parse("?param=xx")
	print(1, u, err)

	u, err = url.Parse("")
	print(2, u, err)

	u, err = url.Parse("https://www.example.jp")
	print(3, u, err)

	u, err = url.Parse("https://www.example.jp/")
	print(4, u, err)

	u, err = url.Parse("https://www.example.jp/path/to")
	print(5, u, err)

}
