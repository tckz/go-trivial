package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	s := `
	some text1
	var x = 'xxxxx';
	some text2
	var x = 'yyyyy';
	some text3
`

	// 最大マッチ回数を指定
	//   負数なら無制限
	// マッチした箇所ごとに[]string。複数マッチしうるので関数の戻りとしては[][]string
	// 0: マッチ全体
	// 1以降: グループに対するcapture

	// [0]=[]string{"var x = 'yyyyy';", "yyyyy"}
	// [1]=[]string{"var x = 'xxxxx';", "xxxxx"}
	reg := regexp.MustCompile(`var x = '(.*)';`)
	for i, e := range reg.FindAllStringSubmatch(s, -1) {
		fmt.Fprintf(os.Stderr, "[%d]=%#v\n", i, e)
	}

	// 最初のマッチを返す
	// ret=[]string{"var x = 'xxxxx';", "xxxxx"}
	ret := reg.FindStringSubmatch(s)
	fmt.Fprintf(os.Stderr, "ret=%#v\n", ret)
}
