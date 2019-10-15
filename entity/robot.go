package entity

type SendGrouptype struct {
	MsgType string   `json:"msgtype"`
	Text    TextType `json:"text"`
	At      AtType   `json:"at"`
	IsAtAll bool     `json:"isAtAll"`
}
type TextType struct {
	Content string `json:"content"`
}
type AtType struct {
	AtMobiles []string `json:"atMobiles"`
}
type SendGroupResp struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}
