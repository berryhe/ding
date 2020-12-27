// Package process 智能工作流/自有 interface
package process

import "github.com/Berry961103/ding"

const (
	apiGetProcessCode = "/topapi/process/get_by_name"
)

// SelfGetByName 获取模板code
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-template-code
// POST https://oapi.dingtalk.com/topapi/process/get_by_name?access_token=ACCESS_TOKEN
func SelfGetByName(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiGetProcessCode, payload, ding.DefaultPostDecodeStr)
}
