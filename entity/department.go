package entity

// 创建部门
type (
	// 创建部门请求
	CreateDepartmentRequest struct {
		Name             string `json:"name"`
		ParentID         int    `json:"parent_id"`
		HideDept         bool   `json:"hide_dept"`
		DeptPermits      string `json:"dept_permits"`
		UserPermits      string `json:"user_permits"`
		OuterDept        bool   `json:"outer_dept"`
		CreateDeptGroup  bool   `json:"create_dept_group"`
		OuterPermitUsers string `json:"outer_permit_users"`
		OuterPermitDepts string `json:"outer_permit_depts"`
		Order            int    `json:"order"`
	}

	// CreateDepartmentResp 创建部门返回
	CreateDepartmentResp struct {
		ErrCode   int                    `json:"errcode"`
		ErrMsg    string                 `json:"errmsg"`
		Result    CreateDepartmentResult `json:"result"`
		RequestID string                 `json:"request_id"`
	}

	// CreateDepartmentResult 创建部门返回主体
	CreateDepartmentResult struct {
		DeptID int `json:"dept_id"`
	}
)

// 更新部门
type (
	UpdateDepartmentRequest struct {
		DeptPermits            string `json:"dept_permits"`
		OuterPermitUsers       string `json:"outer_permit_users"`
		GroupContainOuterDept  bool   `json:"group_contain_outer_dept"`
		OrgDeptOwner           string `json:"org_dept_owner"`
		OuterDept              string `json:"outer_dept"`
		Language               string `json:"language"`
		SourceIdentifier       string `json:"source_identifier"`
		GroupContainSubDept    bool   `json:"group_contain_sub_dept"`
		AutoAddUser            bool   `json:"auto_add_user"`
		DeptManagerUseridList  string `json:"dept_manager_userid_list"`
		ParentID               string `json:"parent_id"`
		HideDept               bool   `json:"hide_dept"`
		Name                   string `json:"name"`
		UserPermits            string `json:"user_permits"`
		OuterPermitDepts       string `json:"outer_permit_depts"`
		GroupContainHiddenDept bool   `json:"group_contain_hidden_dept"`
		OuterDeptOnlySelf      bool   `json:"outer_dept_only_self"`
		DeptID                 string `json:"dept_id"`
		CreateDeptGroup        bool   `json:"create_dept_group"`
		Order                  string `json:"order"`
	}

	UpdateDepartmentResp struct {
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		RequestID string `json:"request_id"`
	}
)

// 获取部门详情
type (
	// GetDepSelfResp 企业内部返回
	GetDepSelfResp struct {
		ErrCode               int    `json:"errcode"`
		SourceIdentifier      string `json:"sourceIdentifier"`
		UserPermits           string `json:"userPermits"`
		OrgDeptOwner          string `json:"orgDeptOwner"`
		OuterDept             bool   `json:"outerDept"`
		ErrMsg                string `json:"errmsg"`
		DeptManagerUseridList string `json:"deptManagerUseridList"`
		ParentID              int    `json:"parentid"`
		GroupContainSubDept   bool   `json:"groupContainSubDept"`
		OuterPermitUsers      string `json:"outerPermitUsers"`
		OuterPermitDepts      string `json:"outerPermitDepts"`
		CreateDeptGroup       bool   `json:"createDeptGroup"`
		Name                  string `json:"name"`
		DeptGroupChatID       string `json:"deptGroupChatId"`
		ID                    int    `json:"id"`
		AutoAddUser           bool   `json:"autoAddUser"`
		DeptHiding            bool   `json:"deptHiding"`
		DeptPermits           string `json:"deptPermits"`
		Order                 int    `json:"order"`
	}

	// GetDepTPResp 第三方返回
	GetDepTPResp struct {
		ErrCode               int    `json:"errcode"`
		SourceIdentifier      string `json:"sourceIdentifier"`
		UserPermits           string `json:"userPermits"`
		OrgDeptOwner          string `json:"orgDeptOwner"`
		OuterDept             bool   `json:"outerDept"`
		ErrMsg                string `json:"errmsg"`
		DeptManagerUseridList string `json:"deptManagerUseridList"`
		GroupContainSubDept   bool   `json:"groupContainSubDept"`
		OuterPermitUsers      string `json:"outerPermitUsers"`
		OuterPermitDepts      string `json:"outerPermitDepts"`
		DeptPerimits          string `json:"deptPerimits"`
		CreateDeptGroup       bool   `json:"createDeptGroup"`
		Name                  string `json:"name"`
		DeptGroupChatID       string `json:"deptGroupChatId"`
		ID                    int    `json:"id"`
		AutoAddUser           bool   `json:"autoAddUser"`
		DeptHiding            bool   `json:"deptHiding"`
		DeptPermits           string `json:"deptPermits"`
		Order                 int    `json:"order"`
	}
)

type (
	ListSubIDsResp struct {
		ErrCode   int          `json:"errcode"`
		ErrMsg    string       `json:"errmsg"`
		Result    SubIDsResult `json:"result"`
		RequestID string       `json:"request_id"`
	}
	SubIDsResult struct {
		DeptIDList []string `json:"dept_id_list"`
	}
)

type (
	DeptUserIDByListParentResp struct {
		ErrCode   int          `json:"errcode"`
		ErrMsg    string       `json:"errmsg"`
		Result    ParentResult `json:"result"`
		RequestID string       `json:"request_id"`
	}
	ParentList struct {
		ParentDeptIDList []int `json:"parent_dept_id_list"`
	}
	ParentResult struct {
		ParentList []ParentList `json:"parent_list"`
	}
)
type (
	ListParentByDeptResp struct {
		ErrCode   int        `json:"errcode"`
		ErrMsg    string     `json:"errmsg"`
		Result    ParentList `json:"result"`
		RequestID string     `json:"request_id"`
	}
)

type (
	ListDepartmentSubResp struct {
		ErrCode   int                       `json:"errcode"`
		ErrMsg    string                    `json:"errmsg"`
		Result    []ListDepartmentSubResult `json:"result"`
		RequestID string                    `json:"request_id"`
	}
	ListDepartmentSubResult struct {
		AutoAddUser     bool   `json:"auto_add_user"`
		CreateDeptGroup bool   `json:"create_dept_group"`
		DeptID          int    `json:"dept_id"`
		Name            string `json:"name"`
		ParentID        int    `json:"parent_id"`
	}
)
