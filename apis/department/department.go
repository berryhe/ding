// Package department 部门管理2.0 interface
package department

import "github.com/Berry961103/ding"

const (
	apiGetGetDepartment     = "/topapi/v2/department/get"
	apiListGetDepartmentSub = "/topapi/v2/department/listsub"
	apiListDepartmentSubID  = "/topapi/v2/department/listsubid"
)

// ListDepartMentSub 获取部门列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-department-list-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=ACCESS_TOKEN
func ListDepartMentSub(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiListGetDepartmentSub, payload, ding.DefaultPostDecodeStr)
}

// GetDepartment 获取部门详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-department-details-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/get?access_token=ACCESS_TOKEN
func GetDepartment(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiGetGetDepartment, payload, ding.DefaultPostDecodeStr)
}

// ListDepartmentSubID 获取子部门ID列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-a-sub-department-id-list-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/listsubid?access_token=ACCESS_TOKEN
func ListDepartmentSubID(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiListDepartmentSubID, payload, ding.DefaultPostDecodeStr)
}
