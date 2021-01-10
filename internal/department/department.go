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

// Package department 部门管理2.0 interface
package department

import "github.com/Berry961103/ding"

const (
	apiGetGetDepartment     = "/topapi/v2/department/get"
	apiListGetDepartmentSub = "/topapi/v2/department/listsub"
	apiListDepartmentSubID  = "/topapi/v2/department/listsubid"
)

// GetDepartment 获取部门详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-department-details-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/get?access_token=ACCESS_TOKEN
func GetDepartment(dctx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dctx.HTTPPost(apiGetGetDepartment, payload, ding.DefaultPostDecodeStr)
}

// ListDepartMentSub 获取部门列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-department-list-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=ACCESS_TOKEN
func ListDepartMentSub(dctx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dctx.HTTPPost(apiListGetDepartmentSub, payload, ding.DefaultPostDecodeStr)
}

// ListDepartmentSubID 获取子部门ID列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-a-sub-department-id-list-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/listsubid?access_token=ACCESS_TOKEN
func ListDepartmentSubID(dctx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dctx.HTTPPost(apiListDepartmentSubID, payload, ding.DefaultPostDecodeStr)
}
