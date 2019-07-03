package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyString string
type MyType struct {
	Name   string   `json:"name"`
	MyName MyString `json:"my_name"`
}

func main() {
	m := &MyType{}
	json.Unmarshal([]byte(`{
		"name": "valname",
		"my_name": "val_myname"
	}`), m)

	// basicなstringを元にしたtypeに対するmarshal/unmarshalが可能か-> 可能
	fmt.Fprintf(os.Stderr, "MyName=%s\n", m.MyName)

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "Marshal=%s\n", string(b))
}
