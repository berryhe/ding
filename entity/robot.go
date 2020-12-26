// Package 钉钉群机器人消息结构体

package entity

type (
	// SendGrouptype text类型消息主体
	SendGrouptype struct {
		MsgType string   `json:"msgtype"`
		Text    TextType `json:"text"`
		At      AtType   `json:"at"`
		IsAtAll bool     `json:"isAtAll"`
	}

	// TextType text内容
	TextType struct {
		Content string `json:"content"`
	}

	// AtType @人的手机号
	AtType struct {
		AtMobiles []string `json:"atMobiles"`
	}
)

// SendGroupResp 返回消息体
type SendGroupResp struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}
