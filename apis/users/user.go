// Package users 用户管理2.0
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
func GetUserInfo(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUser, payload, ding.DefaultPostDecodeStr)
}

// CountUser 获取员工人数
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-number-of-employees-v2
// POST https://oapi.dingtalk.com/topapi/user/count?access_token=ACCESS_TOKEN
func CountUser(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCountUser, payload, ding.DefaultPostDecodeStr)
}

// ListAdmin 获取管理员列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-a-list-of-administrators-v2
// POST https://oapi.dingtalk.com/topapi/user/listadmin?access_token=ACCESS_TOKEN
func ListAdmin(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiListAdmin, payload, ding.DefaultPostDecodeStr)
}
