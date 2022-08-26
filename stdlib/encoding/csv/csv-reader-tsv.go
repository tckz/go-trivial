package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fn := os.Args[1]

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, rec := range records {
		fmt.Printf("[%d]len=%d, %#v\n", i, len(rec), rec)
	}
}
