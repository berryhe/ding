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

// Package robot implements top
package robot

import (
	"bytes"
	"encoding/json"

	"github.com/berryhe/ding"
	"github.com/berryhe/ding/entity"
)

const (
	robotAPI = "/robot/send"
)

// DingRobotText Send ding robot text
func DingRobotText(dCtx *ding.DCtx, msg string, phone []string) error {

	sg := entity.SendGrouptype{
		MsgType: "text",
		Text: entity.TextType{
			Content: msg,
		},
		At: entity.AtType{
			AtMobiles: phone,
		},
		IsAtAll: true,
	}

	reqData, err := json.Marshal(sg)
	if err != nil {
		return err
	}

	_, err = dCtx.RobotHTTPPost(robotAPI, bytes.NewReader(reqData), ding.DefaultPostDecodeStr)
	return err
}
