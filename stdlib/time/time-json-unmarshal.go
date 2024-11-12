package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/samber/lo"
)

func main() {
	type Rec struct {
		Time *time.Time `json:"time"`
	}
	var r Rec
	// time.TimeのデフォルトのUnmarshalでミリ秒扱える
	lo.Must0(json.Unmarshal([]byte(`{
		"time": "2019-12-31T20:01:02.123+09:00"
}`), &r))
	fmt.Printf("%v\n", r.Time)
}
