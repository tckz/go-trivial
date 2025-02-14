package main

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/mattn/go-isatty"
	"github.com/samber/lo"
)

func main() {
	fn := os.Args[1]

	b := lo.Must(os.ReadFile(fn))
	var v any
	if err := yaml.UnmarshalWithOptions(b, &v, yaml.UseOrderedMap()); err != nil {
		w := os.Stderr
		w.WriteString(yaml.FormatError(err, isatty.IsTerminal(w.Fd()), true))
		os.Exit(1)
	}
	lo.Must0(yaml.NewEncoder(os.Stdout, yaml.UseLiteralStyleIfMultiline(true)).Encode(v))
}
