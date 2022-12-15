package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 埋め込み構造体をmarshalしたときプロパティとしては見えてこない

func main() {
	type Embedded struct {
		EmbeddedText string
	}
	type Outer struct {
		Embedded

		OuterText string
	}
	m := &Outer{
		Embedded: Embedded{
			EmbeddedText: "embedded",
		},
		OuterText: "outer",
	}

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	// Marshal={"EmbeddedText":"embedded","OuterText":"outer"}
	fmt.Fprintf(os.Stderr, "Marshal=%s\n", string(b))
}
