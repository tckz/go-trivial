package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
)

// fieldに独自typeかつ独自Unmarshalを実装することで大きい数値のUnmarshalを可能にする

type MyNumber struct {
	num *big.Int
}

type MyFormat struct {
	Number MyNumber `json:"my_number"`
}

func (m MyNumber) String() string {
	return m.num.String()
}

func (m *MyNumber) UnmarshalJSON(data []byte) error {
	m.num = big.NewInt(0)
	m.num.SetString(string(data), 10)
	fmt.Fprintf(os.Stderr, "val=%s\n", string(data))
	return nil
}

func main() {
	j := `{
"my_number": 123456780123456789901234567890
}
`
	m := MyFormat{}
	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", m.Number)
}
