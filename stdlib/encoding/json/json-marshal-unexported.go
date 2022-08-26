package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type HoldUnexported struct {
	name     string
	someTime time.Time
}

// jsonのプロパティ名に合わせたのかunexportedなフィールドばかりのstructをjson.Marshalしてるコードがあったので確認

func main() {
	m := &HoldUnexported{
		name:     "myname",
		someTime: time.Now(),
	}

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	// Marshal={}
	fmt.Fprintf(os.Stderr, "Marshal=%s\n", string(b))
}
