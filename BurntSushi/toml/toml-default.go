package main

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/samber/lo"
)

//go:embed some.toml
var someBytes []byte

func main() {
	type Config struct {
		DefaultValue string `toml:"default_value" default:"wao"`
	}

	dec := toml.NewDecoder(bytes.NewReader(someBytes))
	var c Config
	_ = lo.Must(dec.Decode(&c))

	// DefaultValue=
	// defaultタグは反映されない
	fmt.Printf("DefaultValue=%v\n", c.DefaultValue)
}
