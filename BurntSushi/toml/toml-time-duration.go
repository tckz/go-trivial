package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/samber/lo"
)

//go:embed some.toml
var someBytes []byte

func main() {
	type Config struct {
		Duration time.Duration `toml:"duration"`
	}

	dec := toml.NewDecoder(bytes.NewReader(someBytes))
	var c Config
	_ = lo.Must(dec.Decode(&c))

	// Duration=1h0m0s
	// 2022年あたりのversionからtime.Durationにマップできる
	fmt.Printf("Duration=%v\n", c.Duration)
}
