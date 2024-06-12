package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/samber/lo"
)

// 未知イベントおよび部分的未対応イベントのUnmarshal

func um(s string) {
	ev := lo.Must(webhook.UnmarshalEvent([]byte(s)))

	fmt.Printf("### %T\n", ev)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(ev)
}

func main() {

	// 完全な未知イベント
	// webhook.UnknownEvent
	/*
		{
		  "EventInterface": null,
		  "Type": "anotherevent",
		  "Raw": {
		    "anotherevent": {
		      "somevalue": "12878018534390"
		    },
		    "mode": "",
		    "replyToken": "somesome9999",
		    "source": {
		      "type": "user",
		      "userId": "Uxxxxxx"
		    },
		    "timestamp": 1603072703358,
		    "type": "anotherevent"
		  }
		}
	*/

	um(`
		{
		  "replyToken": "somesome9999",
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
`)

	// 部分的未対応イベント
	// webhook.ThingsEvent
	// ThingsEventは認識されるがtype=connectが認識されずUnknownThingsContentになりRawに入る
	/*
		{
		  "type": "things",
		  "source": {
		    "type": "user",
		    "userId": "uXXX"
		  },
		  "timestamp": 1547817848122,
		  "mode": "active",
		  "webhookEventId": "",
		  "deliveryContext": null,
		  "replyToken": "somesome",
		  "things": {
		    "ThingsContentInterface": null,
		    "Type": "connect",
		    "Raw": {
		      "connectResult": {
		        "resultCode": "success",
		        "connectionId": "XXXX",
		        "type": "onetime"
		      },
		      "deviceId": "tXXX",
		      "type": "connect",
		      "unknownProp": "unknown"
		    }
		  }
		}
	*/

	um(`
		{
			"type": "things",
			"replyToken": "somesome",
			"source": {
				"userId": "uXXX",
				"type": "user"
			},
			"mode": "active",
			"timestamp": 1547817848122,
			"things": {
				"type": "connect",
				"deviceId": "tXXX",
				"unknownProp": "unknown",
				"connectResult": {
					"resultCode": "success",
					"connectionId": "XXXX",
					"type": "onetime"
				}
			}
		}
`)

	// 既知イベント
	// webhook.MessageEvent
	/*
		{
		  "type": "message",
		  "source": {
		    "type": "user",
		    "userId": "uXXX"
		  },
		  "timestamp": 1547817848122,
		  "mode": "active",
		  "webhookEventId": "",
		  "deliveryContext": null,
		  "replyToken": "somesome",
		  "message": {
		    "type": "text",
		    "id": "",
		    "text": "Hello, world!",
		    "quoteToken": ""
		  }
		}
	*/
	um(`
		{
			"type": "message",
			"replyToken": "somesome",
			"source": {
				"userId": "uXXX",
				"type": "user"
			},
			"mode": "active",
			"timestamp": 1547817848122,
			"message": {
				"type": "text",
				"text": "Hello, world!"
			}
		}
`)
}
