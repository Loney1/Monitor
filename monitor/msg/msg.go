package msg

import (
	"encoding/json"
	"fmt"
)

var str =  `{
    "msg_type": "interactive",
    "card": {
        "config": {
                "wide_screen_mode": true,
                "enable_forward": true
        },
        "elements": [{
                "tag": "div",
                "text": {
                        "content": "message body",
                        "tag": "lark_md"
                }
        }],
        "header": {
                "title": {
                        "content": "研发部：Jenkins构建记录",
                        "tag": "plain_text"
                }
        }
    }
}`

func Init_Msg() (*FeiShuMsg, error) {
	msg := FeiShuMsg{}
	err := json.Unmarshal([]byte(str), &msg)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &msg, nil
}

type FeiShuMsg struct {
	MsgType string `json:"msg_type"`
	Card	struct {
		Config struct{
			WideScreenMode bool	`json:"wide_screen_mode"`
			EnableForward  bool	`json:"enable_forward"`
		} `json:"config"`
		Elements []struct{
			Tag  string `json:"tag"`
			Text struct{
				Content string `json:"content"`
				Tag 	string `json:"tag"`
			} `json:"text,omitempty"`

			Actions []struct{
				Tag  string `json:"tag"`
				Text struct{
					Content string `json:"content"`
					Tag 	string `json:"tag"`
				} `json:"text"`
				URL		string `json:url`
				Type 	string `json:"type"`
				Value 	struct{
				} `json:"value"`
			} `json:"actions,omitempty"`
		} `json:"elements"`
		Header struct{
			Title struct{
				Content string `json:"content"`
				Tag 	string `json:"tag"`
			} `json:"title"`
		} `json:"header"`
	} `json:"card"`
}
