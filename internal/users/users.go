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

// Package users 用户操作2.0
package users

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/entity"
)

const (
	apiCreateUser        = "/topapi/v2/user/create"
	apiUpdateUser        = "/topapi/v2/user/update"
	apiDeleteUser        = "/topapi/v2/user/delete"
	apiGetUser           = "/topapi/v2/user/get"
	apiCountUser         = "/topapi/user/count"
	apiListAdmin         = "/topapi/user/listadmin"
	apiGetInActiveUser   = "/topapi/inactive/user/v2/get"
	apiGetAdminScope     = "/topapi/user/get_admin_scope"
	apiDepListUser       = "/topapi/v2/user/list"
	apiDepListUserSimple = "/topapi/user/listsimple"
	apiDepListUserIDs    = "/topapi/user/listid"
	apiGetUserMobileByID = "/topapi/v2/user/getbymobile"
	apiUnionIDByUserID   = "/topapi/user/getbyunionid"
)

// UserCreate 创建用户
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/create-a-user-v2
// POST https://oapi.dingtalk.com/topapi/v2/user/create?access_token=ACCESS_TOKEN
func CreateUser(dCtx ding.DCtx, createUser entity.CreateUserRequest) (resp entity.CountUserResp, err error) {
	reqData, err := json.Marshal(createUser)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiCreateUser, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// UpdateUser 更新用户信息
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/update-user-information-v2
// POST https://oapi.dingtalk.com/topapi/v2/user/update?access_token=ACCESS_TOKEN
func UpdateUser(dCtx ding.DCtx, userUpdateRequest entity.UserUpdateRequest) (resp entity.UserUpdateResp, err error) {
	playLoad, err := json.Marshal(userUpdateRequest)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiUpdateUser, playLoad)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// DeleteUser 删除用户
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/delete-a-user-v2
// POST https://oapi.dingtalk.com/topapi/v2/user/delete?access_token=ACCESS_TOKEN
func DeleteUser(dCtx ding.DCtx, userID string) error {
	playLoad := []byte(fmt.Sprintf(`{"userid":"%s"}`, userID))

	_, err := dCtx.HTTPPost(apiDeleteUser, playLoad)

	return err
}

// GetUserInfo 获取用户详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-user-details-v2
// POST https://oapi.dingtalk.com/topapi/v2/user/get?access_token=ACCESS_TOKEN
func GetUserInfo(dCtx *ding.DCtx, userInfo entity.UserInfoRequest) (resp entity.UserInfoResp, err error) {
	data, err := json.Marshal(userInfo)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiGetUser, data)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// CountUser 获取员工人数
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-number-of-employees-v2
// POST https://oapi.dingtalk.com/topapi/user/count?access_token=ACCESS_TOKEN
// 是否包含未激活钉钉人数：
// 		false：包含未激活钉钉的人员数量。
// 		true：只包含激活钉钉的人员数量。
func CountUser(dCtx *ding.DCtx, onlyActive bool) (int, error) {

	uCount := entity.CountUserRequest{
		OnlyActive: onlyActive,
	}

	reqData, err := json.Marshal(uCount)
	if err != nil {
		return 0, err
	}

	resp, err := dCtx.HTTPPost(apiCountUser, reqData)
	if err == nil {
		return 0, err
	}

	var countUserResp entity.CountUserResp

	if err = json.Unmarshal(resp, &countUserResp); err != nil {
		return 0, err
	}

	return countUserResp.Result.Count, nil
}

// ListAdmin 获取管理员列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-a-list-of-administrators-v2
// POST https://oapi.dingtalk.com/topapi/user/listadmin?access_token=ACCESS_TOKEN
func ListAdmin(dCtx *ding.DCtx) (resp entity.AdminListResp, err error) {

	respData, err := dCtx.HTTPPost(apiListAdmin, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return

}

// LoopGetInActiveUser 帮助循环获取未登录员工列表
func LoopGetInActiveUser(dCtx ding.DCtx, deptIDs []int, queryData string, isActive bool) ([]entity.GetInactiveResp, error) {
	ir := entity.GetInactiveRequest{
		IsActive:  isActive,
		DeptIds:   deptIDs,
		Offset:    0,
		Size:      100,
		QueryDate: queryData,
	}

	var (
		resps []entity.GetInactiveResp
		dfs   func(entity.GetInactiveRequest)
	)

	dfs = func(ir entity.GetInactiveRequest) {
		resp, err := GetInActiveUser(dCtx, ir)
		if err != nil {
			return
		}

		resps = append(resps, resp)
		if resp.Result.HasMore {
			ir.Offset += ir.Size
			dfs(ir)
		}
	}

	dfs(ir)

	return resps, nil
}

// GetInActiveUser 获取未登录钉钉的员工列表
// deptIDs 不传默认整个企业
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-employees-not-logged-v2
// POST https://oapi.dingtalk.com/topapi/inactive/user/v2/get?access_token=ACCESS_TOKEN
func GetInActiveUser(dCtx ding.DCtx, inactiveReq entity.GetInactiveRequest) (resp entity.GetInactiveResp, err error) {

	if inactiveReq.Size > 100 || inactiveReq.Offset < 0 {
		err = errors.New("offset 小于0 或 Size大于100")
		return
	}

	reqData, err := json.Marshal(inactiveReq)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiGetInActiveUser, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// GetAdminScope 获取管理员通讯录权限范围
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-the-administrator-addressbook-permission-v2
// POST https://oapi.dingtalk.com/topapi/user/get_admin_scope?access_token=ACCESS_TOKEN
func GetAdminScope(dCtx ding.DCtx, userID string) (resp entity.AdminScopeResp, err error) {
	reqData := []byte(fmt.Sprintf(`{"userid":"%s"}`, userID))

	respData, err := dCtx.HTTPPost(apiGetAdminScope, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// LoopListUsers 循环帮助获取部门用户详情所有数据
func LoopListUsers(dCtx *ding.DCtx, depUserList entity.DepUserListRequest) (resps []entity.DepUserListResp, err error) {
	var dfs func(entity.DepUserListRequest)

	dfs = func(depUserList entity.DepUserListRequest) {
		resp, err := ListUsers(dCtx, depUserList)
		if err != nil {
			return
		}
		resps = append(resps, resp)

		if resp.Result.HasMore {
			depUserList.Cursor = resp.Result.NextCursor
			dfs(depUserList)
		}
	}

	dfs(depUserList)
	return
}

// 获取部门用户详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-department-user-details-v2
// POST https://oapi.dingtalk.com/topapi/v2/user/list?access_token=ACCESS_TOKEN
func ListUsers(dCtx *ding.DCtx, depUserList entity.DepUserListRequest) (resp entity.DepUserListResp, err error) {
	reqData, err := json.Marshal(depUserList)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiDepListUser, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}
	return

}

// LoopListUserSimple 循环帮助获取部门用户
func LoopListUserSimple(dCtx *ding.DCtx, userListSimple entity.UserListSimpleRequest) (resps []entity.UserListSimpleResp, err error) {
	var dfs func(entity.UserListSimpleRequest)

	dfs = func(userListSimple entity.UserListSimpleRequest) {
		resp, err := ListUserSimple(dCtx, userListSimple)
		if err != nil {
			return
		}
		resps = append(resps, resp)

		if resp.Result.HasMore {
			userListSimple.Cursor = resp.Result.NextCursor
			dfs(userListSimple)
		}
	}

	dfs(userListSimple)
	return
}

// ListUserSimple 获取部门用户
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/obtains-users-of-department-v2
// POST https://oapi.dingtalk.com/topapi/user/listsimple?access_token=ACCESS_TOKEN
func ListUserSimple(dCtx *ding.DCtx, userListSimple entity.UserListSimpleRequest) (resp entity.UserListSimpleResp, err error) {
	reqData, err := json.Marshal(userListSimple)
	if err != nil {
		return
	}

	respData, err := dCtx.HTTPPost(apiDepListUserSimple, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}

// ListUserIDs 获取部门用户userid列表
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/query-the-userid-list-of-a-department-v2
// POST https://oapi.dingtalk.com/topapi/user/listid?access_token=ACCESS_TOKEN
func ListUserIDs(dCtx *ding.DCtx, depID int) (resp entity.ListUsersResp, err error) {
	reqData := []byte(fmt.Sprintf(`{ "dept_id":"%d"}`, depID))

	respData, err := dCtx.HTTPPost(apiDepListUserIDs, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}
	return
}

// GetUserMobileByUserID 根据手机号获取userid
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/query-users-by-phone-number-v2
// POST https://oapi.dingtalk.com/topapi/v2/user/getbymobile?access_token=ACCESS_TOKEN
func GetUserMobileByUserID(dCtx *ding.DCtx, userMobile string) (string, error) {
	reqData := []byte(fmt.Sprintf(`{ "mobile":"%s"}`, userMobile))

	respData, err := dCtx.HTTPPost(apiGetUserMobileByID, reqData)
	if err != nil {
		return "", err
	}

	var resp entity.UserMobileByIDResp
	if err = json.Unmarshal(respData, &resp); err != nil {
		return "", err
	}

	return resp.Result.UserID, nil
}

// GetUnionIDByUser 根据unionid获取用户信息
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/retrieve-user-information-by-unionid-v2
// POST https://oapi.dingtalk.com/topapi/user/getbyunionid?access_token=ACCESS_TOKEN
func GetUnionIDByUser(dCtx *ding.DCtx, unionID string) (resp entity.UnionIDByUserIDResp, err error) {
	reqData := []byte(fmt.Sprintf(`{ "unionid":"%s"}`, unionID))
	respData, err := dCtx.HTTPPost(apiUnionIDByUserID, reqData)
	if err != nil {
		return
	}

	if err = json.Unmarshal(respData, &resp); err != nil {
		return
	}

	return
}
