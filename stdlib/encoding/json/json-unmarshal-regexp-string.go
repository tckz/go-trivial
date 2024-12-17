package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/samber/lo"
)

func main() {
	type MyType struct {
		Regexp *regexp.Regexp `json:"regexp"`
	}

	// RegexpはUnmarshalTextを実装しているのでJSON StringからUnmarshalできる
	var m MyType
	lo.Must0(json.Unmarshal([]byte(`{"regexp": "^this is regexp"}`), &m))

	fmt.Fprintf(os.Stderr, "%s\n", m.Regexp)
}
