package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

func main() {
	t, err := template.New("some").Parse(`
myNumber={{.myNumber}}です
myNumberWithFormat={{printf "%.0f" .myNumber}}です
myNumber2={{.myNumber2}}です
`)
	if err != nil {
		panic(err)
	}

	var v map[string]interface{}
	json.Unmarshal([]byte(`{"myNumber": 1000000, "myNumber2": 123}`), &v)
	fmt.Fprintf(os.Stdout, "Unmarshal: %#v\n", v)

	br := &bytes.Buffer{}
	t.Execute(br, v)
	fmt.Fprintf(os.Stdout, "Templated: %s\n", br.String())

	// 数値をinterface{}にUnmarshalすることでfloat64になり
	// float64をhtml/templateがテキスト化することで指数表記になる。

	// 文字参照部分はhtml/templateを使っているため。
	/*
	   myNumber=1.234567e&#43;06です
	   myNumber2=123です
	*/

}
