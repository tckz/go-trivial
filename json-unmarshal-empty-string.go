package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type MyType struct {
		ID string `json:"id"`
	}

	// 空文字をstructにUnmarshalすると、
	// unexpected end of JSON input
	var m MyType
	err := json.Unmarshal([]byte(""), &m)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", m.ID)
}
