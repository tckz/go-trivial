package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	Text string `json:"text"`
}

type SomeType struct {
	Code     int       `json:"code"`
	Response *Response `json:"response"`
}

func main() {

	// typeの一部に別typeのポインタがある場合に、unmarshalでインスタンスを生成してくれるか？->生成してくれる
	{
		var v SomeType
		err := json.Unmarshal([]byte(`{
"code": 123,
"response": {
	"text": "hello"
}
}`), &v)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(os.Stderr, "code=%d, response.nil?=%t, response.text=%s\n",
			v.Code, v.Response == nil, v.Response.Text)
	}

	// typeの一部に別typeのポインタがある場合に、json側に対応する値がないとnilのまま
	{
		var v SomeType
		err := json.Unmarshal([]byte(`{
"code": 456	
}`), &v)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(os.Stderr, "code=%d, response.nil?=%t\n",
			v.Code, v.Response == nil)
	}
}
