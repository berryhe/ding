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

// Package process 智能工作流/官方 interface
package process

import (
	"github.com/Berry961103/ding"
)

const (
	apiCreate        = "/topapi/processinstance/create"
	apiListIds       = "/topapi/processinstance/listids"
	apiGet           = "/topapi/processinstance/get"
	apiGetTodoNum    = "/topapi/process/gettodonum"
	apiListByUserID  = "/topapi/process/listbyuserid"
	apiCspaceInfo    = "/topapi/processinstance/cspace/info"
	apiCspacePreview = "/topapi/processinstance/cspace/preview"
)

// Create 发起审批实例
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/initiate-DingCtxroval
// POST https://oapi.dingtalk.com/topapi/processinstance/create?access_token=ACCESS_TOKEN
func Create(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiCreate, payload, ding.DefaultPostDecodeStr)
}

// ListIds 批量获取审批实例id
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-an-DingCtxroval-list-of-instance-ids
// POST https://oapi.dingtalk.com/topapi/processinstance/listids?access_token=ACCESS_TOKEN
func ListIds(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiListIds, payload, ding.DefaultPostDecodeStr)
}

// Get 获取审批实例详情
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-details-about-a-specified-DingCtxroval-instance
// POST https://oapi.dingtalk.com/topapi/processinstance/get?access_token=ACCESS_TOKEN
func Get(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiGet, payload, ding.DefaultPostDecodeStr)
}

// GetTodoNum 获取用户待审批数量
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-number-of-requests-for-DingCtxroval-by-a-specified
// POST https://oapi.dingtalk.com/topapi/process/gettodonum?access_token=ACCESS_TOKEN
func GetTodoNum(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiGetTodoNum, payload, ding.DefaultPostDecodeStr)
}

// ListByUserID 获取用户可见的审批模板
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-a-list-of-DingCtxroval-forms-visible-to-the-specified
// POST https://oapi.dingtalk.com/topapi/process/listbyuserid?access_token=ACCESS_TOKEN
func ListByUserID(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiListByUserID, payload, ding.DefaultPostDecodeStr)
}

// CspaceInfo 获取审批钉盘空间信息
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-information-about-DingCtxroval-nail-disk
// POST https://oapi.dingtalk.com/topapi/processinstance/cspace/info?access_token=ACCESS_TOKEN
func CspaceInfo(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiCspaceInfo, payload, ding.DefaultPostDecodeStr)
}

// CspacePreview 预览审批附件
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/preview-authorization-attachment
// POST POST https://oapi.dingtalk.com/topapi/processinstance/cspace/preview?access_token=ACCESS_TOKEN
func CspacePreview(dCtx *ding.DCtx, payload []byte) (resp []byte, err error) {
	return dCtx.HTTPPost(apiCspacePreview, payload, ding.DefaultPostDecodeStr)
}
