package trivial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/goccy/go-yaml"
)

func OutYaml(w io.Writer, srcs ...any) error {
	// yaml.Marshal which compliant with encoding/yaml with types without yaml tag such as GetScheduleOutput outputs keys as lowercase.
	// To avoid it, we marshal it to JSON and decode it again.

	enc := yaml.NewEncoder(w, yaml.UseLiteralStyleIfMultiline(true))
	for _, src := range srcs {
		js, err := json.Marshal(src)
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}

		dec := json.NewDecoder(bytes.NewReader(js))
		// 数値がfloat64になって後続のyaml化で「0.0」になるのを防ぐ。
		dec.UseNumber()
		var v any
		err = dec.Decode(&v)
		if err != nil {
			return fmt.Errorf("json.Decode: %w", err)
		}

		if err = enc.Encode(v); err != nil {
			return fmt.Errorf("yaml.Encode: %w", err)
		}
	}
	return nil
}
