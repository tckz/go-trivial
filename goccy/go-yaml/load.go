package main

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/samber/lo"
)

func main() {
	fn := os.Args[1]

	b := lo.Must(os.ReadFile(fn))
	var v any
	err := yaml.UnmarshalWithOptions(b, &v, yaml.UseOrderedMap())
	if err != nil {
		os.Stderr.WriteString(yaml.FormatError(err, true, true))
		os.Exit(1)
	}
	lo.Must0(yaml.NewEncoder(os.Stdout, yaml.UseLiteralStyleIfMultiline(true)).Encode(v))
}
