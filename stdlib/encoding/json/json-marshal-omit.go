package main

import (
	"encoding/json"
	"os"
)

type SomeInnerType struct {
	InnerText string `json:"inner_text,omitempty"`
}

type SomeType1 struct {
	Text        string        `json:"text"`
	SomeString  string        `json:"some_string,omitempty"`
	SomeInt     int           `json:"some_int,omitempty"`
	Inner       SomeInnerType `json:"inner,omitempty"`
	StringSlice []string      `json:"string_slice,omitempty"`
}

type SomeType2 struct {
	Text       string         `json:"text"`
	SomeString *string        `json:"some_string,omitempty"`
	SomeInt    *int           `json:"some_int,omitempty"`
	Inner      *SomeInnerType `json:"inner,omitempty"`
}

// omitemptyがついている属性の値がその型のdefault値であれば省略
//
// ただし、型がstructの場合は当該structの全属性がdefault値であってもomitされない

func main() {
	// {"text":"Marshal SomeType1","inner":{}}
	enc := json.NewEncoder(os.Stderr)
	v1 := SomeType1{
		Text: "Marshal SomeType1",
	}
	enc.Encode(&v1)

	// {"text":"Marshal SomeType1 with empty slice","inner":{}}
	v1 = SomeType1{
		Text:        "Marshal SomeType1 with empty slice",
		StringSlice: []string{},
	}
	enc.Encode(&v1)

	// {"text":"Marshal SomeType2"}
	v2 := SomeType2{
		Text: "Marshal SomeType2",
	}
	enc.Encode(&v2)

	// {"text":"Marshal SomeType2 with default value","some_string":"","some_int":0,"inner":{}}
	var i int
	var s string
	v3 := SomeType2{
		Text:       "Marshal SomeType2 with default value",
		SomeString: &s,
		SomeInt:    &i,
		Inner:      &SomeInnerType{},
	}
	enc.Encode(&v3)
}
