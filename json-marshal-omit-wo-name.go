package main

import (
	"encoding/json"
	"os"
)

// jsonタグにプロパティ名を書かずにomitemptyだけ書いて効果があるか->ある

func main() {

	type SomeType struct {
		Text  string
		Empty string
		Value string `json:",omitempty"`
	}

	// {"Text":"Marshal SomeType","Empty":""}
	enc := json.NewEncoder(os.Stderr)
	v1 := SomeType{
		Text: "Marshal SomeType",
	}
	enc.Encode(&v1)
}
