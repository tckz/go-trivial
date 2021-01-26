package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonBigInt struct {
	BigValue int64 `json:"big_value"`
}

func main() {

	// goの中ではbigint(2^63-1)のMarshalが指数表記にならないし、そのままbigintにUnmarshalできる

	var v JsonBigInt
	err := json.Unmarshal([]byte(`
	{
		"big_value": 9223372036854775807
	}
	`), &v)
	if err != nil {
		panic(err)
	}
	// v={BigValue:9223372036854775807}
	fmt.Fprintf(os.Stderr, "v=%+v\n", v)

	b, err := json.Marshal(&v)
	if err != nil {
		panic(err)
	}
	// marshal={"big_value":9223372036854775807}
	fmt.Fprintf(os.Stderr, "marshal=%s\n", string(b))
}
