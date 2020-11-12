package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}

	body := []byte(`
{
	"events": [
		{
		  "replyToken": "5ce4d7a548214faa9e8ba863c84a82f6",
		  "type": "anotherevent",
		  "mode": "",
		  "timestamp": 1603072703358,
		  "source": {
			"type": "user",
			"userId": "Uxxxxxx"
		  },
		  "anotherevent": {
			"somevalue": "12878018534390"
		  }
		}
	]
}
`)
	body = []byte(`
{
	"events": [
{
	"type": "things",
	"replyToken": "0f3779fba3b349968c5d07db31eab56f",
	"source": {
		"userId": "uXXX",
		"type": "user"
	},
	"mode": "active",
	"timestamp": 1547817848122,
	"things": {
		"type": "connect",
		"deviceId": "tXXX",
		"connectResult": {
			"resultCode": "success",
			"connectionId": "XXXX",
			"type": "onetime"
		}
	}
}
]
}
`)
	if err := json.Unmarshal(body, request); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", request.Events[0])

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(&request)
}
