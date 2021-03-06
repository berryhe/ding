// MIT License

// Copyright (c) 2019 Berryhe

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
