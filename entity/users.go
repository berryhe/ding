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

// Package entity users 请求体
package entity

// 获取用户详情
type (
	// UserInfoRequest 获取用户详情请求
	UserInfoRequest struct {
		Language string `json:"language"`
		UserID   string `json:"userid"`
	}
	// UserInfoResp 获取用户详情返回
	UserInfoResp struct {
		Errcode   int            `json:"errcode"`
		Result    UserInfoResult `json:"result"`
		Errmsg    string         `json:"errmsg"`
		RequestID string         `json:"request_id"`
	}
	// UserInfoResult 用户详情主体
	UserInfoResult struct {
		Extension     string          `json:"extension"`
		Unionid       string          `json:"unionid"`
		Boss          bool            `json:"boss"`
		UnionEmpExt   UnionEmpExt     `json:"unionEmpExt"`
		RoleList      []RoleList      `json:"role_list"`
		Admin         bool            `json:"admin"`
		Remark        string          `json:"remark"`
		Title         string          `json:"title"`
		HiredDate     int64           `json:"hired_date"`
		Userid        string          `json:"userid"`
		WorkPlace     string          `json:"work_place"`
		DeptOrderList []DeptOrderList `json:"dept_order_list"`
		RealAuthed    bool            `json:"real_authed"`
		DeptIDList    []int           `json:"dept_id_list"`
		JobNumber     string          `json:"job_number"`
		Email         string          `json:"email"`
		LeaderInDept  []LeaderInDept  `json:"leader_in_dept"`
		Mobile        string          `json:"mobile"`
		Active        bool            `json:"active"`
		Telephone     string          `json:"telephone"`
		Avatar        string          `json:"avatar"`
		HideMobile    bool            `json:"hide_mobile"`
		Senior        bool            `json:"senior"`
		Name          string          `json:"name"`
		StateCode     string          `json:"state_code"`
	}

	// UnionEmpExt 当用户来自于关联组织时的关联信息。
	UnionEmpExt struct {
		CorpID          string            `json:"corpId"`
		Userid          string            `json:"userid"`
		UnionEmpMapList []UnionEmpMapList `json:"unionEmpMapList"`
	}
	// UnionEmpMapList 关联映射关系
	UnionEmpMapList struct {
		CorpID string `json:"corpId"`
		Userid string `json:"userid"`
	}
	// RoleList 角色列表。
	RoleList struct {
		GroupName string `json:"group_name"`
		ID        int    `json:"id"`
		Name      string `json:"name"`
	}
	// DeptOrderList 员工在对应的部门中的排序
	DeptOrderList struct {
		DeptID int   `json:"dept_id"`
		Order  int64 `json:"order"`
	}
	// LeaderInDept 员工在对应的部门中是否领导
	LeaderInDept struct {
		DeptID int  `json:"dept_id"`
		Leader bool `json:"leader"`
	}
)

// 获取员工人数
type (
	// CountUserRequest 获取员工人数请求
	CountUserRequest struct {
		OnlyActive bool `json:"only_active"`
	}
	// CountUserResp 获取员工人数返回
	CountUserResp struct {
		ErrCode   int             `json:"errcode"`
		ErrMsg    string          `json:"errmsg"`
		Result    CountUserResult `json:"result"`
		RequestID string          `json:"request_id"`
	}
	// CountUserResult 返回消息主体内容
	CountUserResult struct {
		Count int `json:"count"`
	}
)

// 获取管理员
type (
	// UserAdminListResp 获取钉钉管理员
	UserAdminListResp struct {
		Errcode   int                   `json:"errcode"`
		Errmsg    string                `json:"errmsg"`
		Result    []UserAdminListResult `json:"result"`
		RequestID string                `json:"request_id"`
	}
	// UserAdminListResult 管理员信息主体
	UserAdminListResult struct {
		SysLevel int    `json:"sys_level"`
		Userid   string `json:"userid"`
	}
)

// 创建用户
type (
	CreateUserRequest struct {
		Extension     string              `json:"extension"`
		Mobile        string              `json:"mobile"`
		Remark        string              `json:"remark"`
		Telephone     string              `json:"telephone"`
		HideMobile    string              `json:"hide_mobile"`
		HiredDate     string              `json:"hired_date"`
		Title         string              `json:"title"`
		Userid        string              `json:"userid"`
		DeptTitleList []UserDeptTitleList `json:"dept_title_list"`
		WorkPlace     string              `json:"work_place"`
		DeptOrderList []UserDeptOrderList `json:"dept_order_list"`
		SeniorMode    string              `json:"senior_mode"`
		OrgEmail      string              `json:"org_email"`
		Name          string              `json:"name"`
		DeptIDList    string              `json:"dept_id_list"`
		JobNumber     string              `json:"job_number"`
		Email         string              `json:"email"`
	}
	UserDeptTitleList struct {
		DeptID int    `json:"dept_id"`
		Title  string `json:"title"`
	}
	UserDeptOrderList struct {
		DeptID int `json:"dept_id"`
		Order  int `json:"order"`
	}

	CreateUserResp struct {
		ErrCode   int          `json:"errcode"`
		Result    CreateResult `json:"result"`
		RequestID string       `json:"request_id"`
	}
	CreateResult struct {
		Userid string `json:"userid"`
	}
)

// 更新用户
type (
	UserUpdateRequest struct {
		Extension     string              `json:"extension"`
		Telephone     string              `json:"telephone"`
		Remark        string              `json:"remark"`
		HideMobile    string              `json:"hide_mobile"`
		Title         string              `json:"title"`
		HiredDate     string              `json:"hired_date"`
		Userid        string              `json:"userid"`
		DeptTitleList []UserDeptTitleList `json:"dept_title_list"`
		WorkPlace     string              `json:"work_place"`
		DeptOrderList []UserDeptOrderList `json:"dept_order_list"`
		SeniorMode    string              `json:"senior_mode"`
		OrgEmail      string              `json:"org_email"`
		Name          string              `json:"name"`
		DeptIDList    string              `json:"dept_id_list"`
		JobNumber     string              `json:"job_number"`
		Email         string              `json:"email"`
	}

	UserUpdateResp struct {
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		RequestID string `json:"request_id"`
	}
)

// 获取管理员列表
type (
	AdminListResp struct {
		ErrCode   int           `json:"errcode"`
		ErrMsg    string        `json:"errmsg"`
		Result    []AdminResult `json:"result"`
		RequestID string        `json:"request_id"`
	}
	AdminResult struct {
		SysLevel int    `json:"sys_level"`
		Userid   string `json:"userid"`
	}
)

// 获取未登录钉钉的员工列表
type (
	GetInactiveRequest struct {
		IsActive  bool   `json:"is_active"`
		Offset    int    `json:"offset"`
		Size      int    `json:"size"`
		DeptIds   []int  `json:"dept_ids"`
		QueryDate string `json:"query_date"`
	}

	GetInactiveResp struct {
		ErrCode   int               `json:"errcode"`
		Result    GetInactiveResult `json:"result"`
		RequestID string            `json:"request_id"`
	}

	GetInactiveResult struct {
		HasMore bool     `json:"has_more"`
		List    []string `json:"list"`
	}
)

// AdminScopeResp 管理员权限
type AdminScopeResp struct {
	DeptIDs   []int  `json:"dept_ids"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	RequestID string `json:"request_id"`
}

// 获取部门用户详情
type (
	DepUserListRequest struct {
		DeptID     int    `json:"dept_id"`
		Cursor     int    `json:"cursor"`
		Size       int    `json:"size"`
		OrderField string `json:"order_field"`
		Language   string `json:"language"`
	}

	DepUserListResp struct {
		ErrCode   int           `json:"errcode"`
		ErrMsg    string        `json:"errmsg"`
		Result    DepUserResult `json:"result"`
		RequestID string        `json:"request_id"`
	}
	DepUserList struct {
		Active           bool   `json:"active"`
		Admin            bool   `json:"admin"`
		Avatar           string `json:"avatar"`
		Boss             bool   `json:"boss"`
		DeptIDList       []int  `json:"dept_id_list"`
		DeptOrder        int    `json:"dept_order"`
		ExclusiveAccount bool   `json:"exclusiveAccount"`
		Extension        string `json:"extension"`
		HideMobile       bool   `json:"hide_mobile"`
		Leader           bool   `json:"leader"`
		Mobile           string `json:"mobile"`
		Name             string `json:"name"`
		Remark           string `json:"remark"`
		StateCode        string `json:"state_code"`
		Telephone        string `json:"telephone"`
		UnionID          string `json:"unionid"`
		UserID           string `json:"userid"`
		WorkPlace        string `json:"work_place"`
	}
	DepUserResult struct {
		HasMore    bool          `json:"has_more"`
		List       []DepUserList `json:"list"`
		NextCursor int           `json:"next_cursor"`
	}
)

// 获取部门用户
type (
	UserListSimpleRequest struct {
		Cursor             int    `json:"cursor"`
		ContainAccessLimit bool   `json:"contain_access_limit"`
		Size               int    `json:"size"`
		OrderField         string `json:"order_field"`
		Language           string `json:"language"`
		DeptID             string `json:"dept_id"`
	}

	UserListSimpleResp struct {
		ErrCode   int                  `json:"errcode"`
		ErrMsg    string               `json:"errmsg"`
		Result    UserListSimpleResult `json:"result"`
		RequestID string               `json:"request_id"`
	}
	UserListSimpleList struct {
		Name   string `json:"name"`
		Userid string `json:"userid"`
	}
	UserListSimpleResult struct {
		HasMore    bool                 `json:"has_more"`
		NextCursor int                  `json:"next_cursor"`
		List       []UserListSimpleList `json:"list"`
	}
)

// 获取部门用户userid列表
type (
	ListUsersResp struct {
		ErrCode   int             `json:"errcode"`
		ErrMsg    string          `json:"errmsg"`
		Result    ListUsersResult `json:"result"`
		RequestID string          `json:"request_id"`
	}
	ListUsersResult struct {
		UseridList []string `json:"userid_list"`
	}
)

// 根据手机号获取userid
type (
	UserMobileByIDResp struct {
		ErrCode   string               `json:"errcode"`
		ErrMsg    string               `json:"errmsg"`
		Result    UserMobileByIDResult `json:"result"`
		RequestID string               `json:"request_id"`
	}
	UserMobileByIDResult struct {
		UserID string `json:"userid"`
	}
)

// 根据unionID获取用户信息
type (
	UnionIDByUserIDResp struct {
		ErrCode   string              `json:"errcode"`
		ErrMsg    string              `json:"errmsg"`
		Result    UnionIDByUserResult `json:"result"`
		RequestID string              `json:"request_id"`
	}
	UnionIDByUserResult struct {
		ContactType string `json:"contact_type"`
		UserID      string `json:"userid"`
	}
)
