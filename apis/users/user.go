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

// Package users 用户管理2.0 interface
package users

import "github.com/Berry961103/ding"

const (
	apiGetUser   = "/topapi/v2/user/get"
	apiCountUser = "/topapi/user/count"
	apiListAdmin = "/topapi/user/listadmin"
)

// GetUserInfo 获取用户详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-user-details-v2
// POST https://oapi.dingtalk.com/topapi/v2/user/get?access_token=ACCESS_TOKEN
func GetUserInfo(dctx *ding.DingCtx, payload []byte) (resp []byte, err error) {
	return dctx.HTTPPost(apiGetUser, payload, ding.DefaultPostDecodeStr)
}

// CountUser 获取员工人数
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-number-of-employees-v2
// POST https://oapi.dingtalk.com/topapi/user/count?access_token=ACCESS_TOKEN
func CountUser(dctx *ding.DingCtx, payload []byte) (resp []byte, err error) {
	return dctx.HTTPPost(apiCountUser, payload, ding.DefaultPostDecodeStr)
}

// ListAdmin 获取管理员列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-a-list-of-administrators-v2
// POST https://oapi.dingtalk.com/topapi/user/listadmin?access_token=ACCESS_TOKEN
func ListAdmin(dctx *ding.DingCtx, payload []byte) (resp []byte, err error) {
	return dctx.HTTPPost(apiListAdmin, payload, ding.DefaultPostDecodeStr)
}
