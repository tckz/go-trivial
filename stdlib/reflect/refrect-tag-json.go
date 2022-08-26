package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	type JsonMapType struct {
		Field1              string `json:"-"`
		Field2              string `json:"field2"`
		Field3              string `json:"field3,omitempty"`
		FieldWithoutJSONTag string
	}

	v := JsonMapType{}
	t := reflect.TypeOf(v)

	/*
		[0]Field1, tag=-, name=-
		[1]Field2, tag=field2, name=field2
		[2]Field3, tag=field3,omitempty, name=field3
		[3]FieldWithoutJSONTag, tag=, name=
	*/
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		name := strings.SplitN(tag, ",", 2)[0]
		fmt.Fprintf(os.Stderr, "[%d]%s, tag=%s, name=%s\n", i, f.Name, tag, name)
	}
}
