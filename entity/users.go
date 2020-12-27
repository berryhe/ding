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
		OnlyActive string `json:"only_active"`
	}
	// CountUserResp 获取员工人数返回
	CountUserResp struct {
		Errcode   int             `json:"errcode"`
		Errmsg    string          `json:"errmsg"`
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
