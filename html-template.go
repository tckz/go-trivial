package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("message").Parse(`you & i {{.Param}}
`)
	if err != nil {
		panic(err)
	}

	// htmlエスケープは置換箇所が対象。テンプレートの中にHTMLエスケープされうる文字があっても関係なし
	// you & i hell &amp; o
	err = t.Execute(os.Stdout, map[string]interface{}{"Param": "hell & o"})
	if err != nil {
		panic(err)
	}
}
