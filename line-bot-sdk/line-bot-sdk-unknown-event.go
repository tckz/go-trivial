package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// 未知イベントおよび部分的未対応イベントのUnmarshal

func um(s string) {
	b := []byte(s)
	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}

	if err := json.Unmarshal(b, request); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", request.Events[0])

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(&request)
}

func main() {

	// 完全な未知イベント
	/*
		{
		  "events": [
		    {
		      "replyToken": "907694c2cc034aa3aa6a21ffc3224a3d",
		      "type": "anotherevent",
		      "mode": "",
		      "timestamp": 1603072703358,
		      "source": {
		        "type": "user",
		        "userId": "Uxxxxxx"
		      }
		    }
		  ]
		}
	*/

	um(`
{
	"events": [
		{
		  "replyToken": "907694c2cc034aa3aa6a21ffc3224a3d",
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

	// 部分的未対応イベント
	/*
		{
		  "events": [
		    {
		      "replyToken": "af34dadc26d1446ab4916f977e6f8e38",
		      "type": "things",
		      "mode": "active",
		      "timestamp": 1547817848122,
		      "source": {
		        "type": "user",
		        "userId": "uXXX"
		      },
		      "things": {
		        "deviceId": "tXXX",
		        "type": "connect"
		      }
		    }
		  ]
		}
	*/

	um(`
{
	"events": [
		{
			"type": "things",
			"replyToken": "af34dadc26d1446ab4916f977e6f8e38",
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
}
