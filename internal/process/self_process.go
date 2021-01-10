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

// Package process 智能工作流/自有 interface
package process

import "github.com/Berry961103/ding"

const (
	apiGetProcessCode = "/topapi/process/get_by_name"
)

// SelfGetByName 获取模板code
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-template-code
// POST https://oapi.dingtalk.com/topapi/process/get_by_name?access_token=ACCESS_TOKEN
func SelfGetByName(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiGetProcessCode, payload, ding.DefaultPostDecodeStr)
}
