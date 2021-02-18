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

import (
	"encoding/json"
	"fmt"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/entity"
)

const (
	apiCreateDepartment     = "/topapi/v2/department/create"
	apiUpdateDepartment     = "/topapi/v2/department/update"
	apiDeleteDepartment     = "/topapi/v2/department/delete"
	apiGetGetDepartment     = "/topapi/v2/department/get"
	apiListGetDepartmentSub = "/topapi/v2/department/listsub"
	apiListDepartmentSubID  = "/topapi/v2/department/listsubid"
	apiListParentByUser     = "/topapi/v2/department/listparentbyuser"
	apiListParentByDept     = "/topapi/v2/department/listparentbydept"
)

// CreateDepartment 创建部门
// See https://ding-doc.dingtalk.com/document/app/create-a-department-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/create?access_token=ACCESS_TOKEN
func CreateDepartment(dCtx *ding.DCtx, createDepRequest entity.CreateDepartmentRequest) (resp entity.CreateDepartmentResp, err error) {
	reqData, err := json.Marshal(createDepRequest)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiCreateDepartment, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// UpdateDepartment 更新部门
// See https://ding-doc.dingtalk.com/document/app/update-a-department-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/update?access_token=ACCESS_TOKEN
func UpdateDepartment(dCtx *ding.DCtx, updateDepRequest entity.UpdateDepartmentRequest) (resp entity.UpdateDepartmentResp, err error) {
	reqData, err := json.Marshal(updateDepRequest)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiUpdateDepartment, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// DeleteDepartment 删除部门
// See https://ding-doc.dingtalk.com/document/app/delete-a-department-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/delete?access_token=ACCESS_TOKEN
func DeleteDepartment(dCtx *ding.DCtx, depID int) error {
	reqData := []byte(fmt.Sprintf(`{"dep_id:%d"}`, depID))

	_, err := dCtx.HTTPPost(apiDeleteDepartment, reqData)
	if err != nil {
		return err
	}

	return nil
}

// GetDepartmentSelf  获取部门详情 企业内部
func GetDepartmentSelf(dCtx *ding.DCtx, depID string) (resp entity.GetDepSelfResp, err error) {
	respData, err := getDepartment(dCtx, depID)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}
	return
}

// GetDepartmentTP  获取部门详情 第三方
func GetDepartmentTP(dCtx *ding.DCtx, depID string) (resp entity.GetDepTPResp, err error) {
	respData, err := getDepartment(dCtx, depID)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}
	return
}

// GetDepartment 获取部门详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-department-details-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/get?access_token=ACCESS_TOKEN
func getDepartment(dCtx *ding.DCtx, depID string) ([]byte, error) {
	reqData := []byte(fmt.Sprintf(`{"id:"%s","lang":"zh_CN"}`, depID))

	return dCtx.HTTPPost(apiGetGetDepartment, reqData)

}

// ListDepartmentSubID 获取子部门ID列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-a-sub-department-id-list-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/listsubid?access_token=ACCESS_TOKEN
func ListDepartmentSubID(dCtx *ding.DCtx, deptID int) (resp entity.ListSubIDsResp, err error) {

	if deptID == 0 {
		deptID = 1
	}

	reqData := []byte(fmt.Sprintf(`{"dept_id:"%d"}`, deptID))

	respData, err := dCtx.HTTPPost(apiListDepartmentSubID, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// DeptUserIDByListParent 获取指定用户的所有父部门列表
// See https://ding-doc.dingtalk.com/document/app/queries-the-list-of-all-parent-departments-of-a-user
// POST https://oapi.dingtalk.com/topapi/v2/department/listparentbyuser?access_token=ACCESS_TOKEN
func DeptUserIDByListParent(dCtx *ding.DCtx, dUserID string) (resp entity.DeptUserIDByListParentResp, err error) {
	reqData := []byte(fmt.Sprintf(`{"userid:"%s"}`, dUserID))

	respData, err := dCtx.HTTPPost(apiListParentByUser, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// ListParentByDept 获取指定部门的所有父部门列表
// See https://ding-doc.dingtalk.com/document/app/query-the-list-of-all-parent-departments-of-a-department
// POST https://oapi.dingtalk.com/topapi/v2/department/listparentbydept?access_token=ACCESS_TOKEN
func ListParentByDept(dCtx *ding.DCtx, deptID int) (resp entity.ListParentByDeptResp, err error) {

	if deptID == 0 {
		deptID = 1
	}

	reqData := []byte(fmt.Sprintf(`{"dept_id:"%d"}`, deptID))

	respData, err := dCtx.HTTPPost(apiListParentByDept, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// ListDepartmentSub 获取部门列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-department-list-v2
// POST https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=ACCESS_TOKEN
func ListDepartmentSub(dCtx *ding.DCtx, deptID int) (resp entity.ListDepartmentSubResp, err error) {
	if deptID == 0 {
		deptID = 1
	}

	reqData := []byte(fmt.Sprintf(`{"dept_id:"%d","lang":"zh_CN"}`, deptID))

	respData, err := dCtx.HTTPPost(apiListGetDepartmentSub, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}
