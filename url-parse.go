package main

import (
	"fmt"
	"net/url"
)

func print(index int, u *url.URL, err error) {
	if err == nil {
		fmt.Printf("url[%d]=%v, query=%v\n", index, u, u.Query())
	} else {
		fmt.Printf("url[%d]=%v, err=%v\n", index, u, err)
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

}
