package main

import (
	"os"
	"text/template"
)

func main() {
	t, err := template.New("message").Option("missingkey=zero").Parse(`
missing key value={{.missingKey}}
`)
	if err != nil {
		panic(err)
	}

	// missingkey=zeroの挙動
	// https://github.com/golang/go/issues/24963#issuecomment-400086111

	/*
		値がinterface{}の場合、ゼロ値＝nilなので"<no value>"に
		missing key value=<no value>
	*/
	err = t.Execute(os.Stdout, map[string]interface{}{})
	if err != nil {
		panic(err)
	}

	/*
		値がstringの場合、ゼロ値＝空文字なので""に
		missing key value=
	*/
	err = t.Execute(os.Stdout, map[string]string{})
	if err != nil {
		panic(err)
	}
}
