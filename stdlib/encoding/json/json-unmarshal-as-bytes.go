package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSONの一部を加工しない[]byteで受け取る

var _ json.Unmarshaler = (*JSONValue)(nil)

type JSONValue struct {
	json []byte
}

func (v *JSONValue) UnmarshalJSON(data []byte) error {
	v.json = data
	return nil
}

func ParseObjects(b []byte) ([]*JSONValue, error) {
	var objs = struct {
		Objects []*JSONValue `json:"objects"`
	}{}

	err := json.Unmarshal(b, &objs)
	if err != nil {
		return nil, err
	}
	return objs.Objects, nil
}

func main() {
	for i, e := range []string{
		// [0]err=<nil>, len(objs)=2, nil?=false
		// [0][0]obj={"a":1}, nil?=false
		// [0][1]obj={"b":2}, nil?=false
		`{"objects":[{"a":1},{"b":2}]}`,
		// [1]err=<nil>, len(objs)=5, nil?=false
		// [1][0]obj="a", nil?=false
		// [1][1]obj="b", nil?=false
		// [1][2]obj=1, nil?=false
		// [1][3]obj=2, nil?=false
		// [1][4]obj=3, nil?=false
		`{"objects":["a", "b", 1, 2, 3]}`,
		// [2]err=<nil>, len(objs)=1, nil?=false
		// [2][0]obj=, nil?=true
		`{"objects":[null]}`,
		// 該当部分がjson要素として壊れていればエラーになる
		// [3]err=invalid character 'b' looking for beginning of value, len(objs)=0, nil?=true
		`{"objects":[bad json]}`,
		// [4]err=<nil>, len(objs)=2, nil?=false
		// [4][0]obj=true, nil?=false
		// [4][1]obj=false, nil?=false
		`{"objects":[true, false]}`,
		// [5]err=<nil>, len(objs)=2, nil?=false
		// [5][0]obj=[1,2], nil?=false
		// [5][1]obj=[3,4], nil?=false
		`{"objects":[[1,2], [3,4]]}`,
		// おまけ
		// [6]err=<nil>, len(objs)=0, nil?=false
		`{"objects":[]}`,
		// おまけ
		// [7]err=<nil>, len(objs)=0, nil?=true
		`{"wao":1}`,
		// おまけ
		// [8]err=json: cannot unmarshal string into Go value of type struct { Objects []*main.JSONValue "json:\"objects\"" }, len(objs)=0, nil?=true
		`"valid JSON but not object"`,
	} {
		objs, err := ParseObjects([]byte(e))
		fmt.Fprintf(os.Stderr, "[%d]err=%v, len(objs)=%d, nil?=%t\n", i, err, len(objs), objs == nil)
		for j, obj := range objs {
			t := ""
			isNil := obj == nil
			if !isNil {
				t = string(obj.json)
			}
			fmt.Fprintf(os.Stderr, "[%d][%d]obj=%s, nil?=%t\n", i, j, t, isNil)
		}
	}
}
