package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MyType struct {
	Name     string
	SomeTime time.Time
}

// time.Timeをjson.Marshalするとどうなるか
// RFC3339Nanoが適用される。time.Time.MarshalJSON()
// "2020-09-11T13:16:28.988983366+09:00"

func main() {
	m := &MyType{
		Name:     "myname",
		SomeTime: time.Now(),
	}

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	// Marshal={"Name":"myname","SomeTime":"2020-09-11T13:16:28.988983366+09:00"}
	fmt.Fprintf(os.Stderr, "Marshal=%s\n", string(b))
}
