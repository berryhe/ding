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
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/initiate-approval
// POST https://oapi.dingtalk.com/topapi/processinstance/create?access_token=ACCESS_TOKEN
func Create(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiCreate, payload, ding.DefaultPostDecodeStr)
}

// ListIds 批量获取审批实例id
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-an-approval-list-of-instance-ids
// POST https://oapi.dingtalk.com/topapi/processinstance/listids?access_token=ACCESS_TOKEN
func ListIds(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiListIds, payload, ding.DefaultPostDecodeStr)
}

// Get 获取审批实例详情
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-details-about-a-specified-approval-instance
// POST https://oapi.dingtalk.com/topapi/processinstance/get?access_token=ACCESS_TOKEN
func Get(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiGet, payload, ding.DefaultPostDecodeStr)
}

// GetTodoNum 获取用户待审批数量
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-number-of-requests-for-approval-by-a-specified
// POST https://oapi.dingtalk.com/topapi/process/gettodonum?access_token=ACCESS_TOKEN
func GetTodoNum(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiGetTodoNum, payload, ding.DefaultPostDecodeStr)
}

// ListByUserID 获取用户可见的审批模板
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-a-list-of-approval-forms-visible-to-the-specified
// POST https://oapi.dingtalk.com/topapi/process/listbyuserid?access_token=ACCESS_TOKEN
func ListByUserID(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiListByUserID, payload, ding.DefaultPostDecodeStr)
}

// CspaceInfo 获取审批钉盘空间信息
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-the-information-about-approval-nail-disk
// POST https://oapi.dingtalk.com/topapi/processinstance/cspace/info?access_token=ACCESS_TOKEN
func CspaceInfo(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiCspaceInfo, payload, ding.DefaultPostDecodeStr)
}

// CspacePreview 预览审批附件
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/preview-authorization-attachment
// POST POST https://oapi.dingtalk.com/topapi/processinstance/cspace/preview?access_token=ACCESS_TOKEN
func CspacePreview(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiCspacePreview, payload, ding.DefaultPostDecodeStr)
}
