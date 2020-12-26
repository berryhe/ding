// Package process 智能工作流/官方
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
// See: https://ding-doc.dingtalk.com/doc#/serverapi2/cmct1a
// POST https://oapi.dingtalk.com/topapi/processinstance/create?access_token=ACCESS_TOKEN
func Create(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, payload, ding.DefaultPostDecodeStr)
}

// ListIds 批量获取审批实例id
// See: https://ding-doc.dingtalk.com/doc#/serverapi2/hh8lx5
// POST https://oapi.dingtalk.com/topapi/processinstance/listids?access_token=ACCESS_TOKEN
func ListIds(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListIds, payload, ding.DefaultPostDecodeStr)
}

// Get 获取审批实例详情
// See: https://ding-doc.dingtalk.com/doc#/serverapi2/xgqkvx
// POST https://oapi.dingtalk.com/topapi/processinstance/get?access_token=ACCESS_TOKEN
func Get(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, payload, ding.DefaultPostDecodeStr)
}

// GetTodoNum 获取用户待审批数量
// See: https://ding-doc.dingtalk.com/doc#/serverapi2/ui5305
// POST https://oapi.dingtalk.com/topapi/process/gettodonum?access_token=ACCESS_TOKEN
func GetTodoNum(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetTodoNum, payload, ding.DefaultPostDecodeStr)
}

// ListByUserID 获取用户可见的审批模板
// See: https://ding-doc.dingtalk.com/doc#/serverapi2/tcwmha
// POST https://oapi.dingtalk.com/topapi/process/listbyuserid?access_token=ACCESS_TOKEN
func ListByUserID(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListByUserID, payload, ding.DefaultPostDecodeStr)
}

// CspaceInfo 获取审批钉盘空间信息
// See: https://ding-doc.dingtalk.com/doc#/serverapi2/xq6ml3
// POST https://oapi.dingtalk.com/topapi/processinstance/cspace/info?access_token=ACCESS_TOKEN
func CspaceInfo(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCspaceInfo, payload, ding.DefaultPostDecodeStr)
}

// CspacePreview 预览审批附件
// See: https://ding-doc.dingtalk.com/doc#/serverapi2/sg687u
// POST https://oapi.dingtalk.com/topapi/processinstance/cspace/preview?access_token=ACCESS_TOKEN
func CspacePreview(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCspacePreview, payload, ding.DefaultPostDecodeStr)
}
