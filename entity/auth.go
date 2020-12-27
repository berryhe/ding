// Package entity 获取通讯录权限范围结构体
package entity

type (
	// AuthScopesResp 获取通讯录权限范围返回
	AuthScopesResp struct {
		Errcode        int           `json:"errcode"`
		ConditionField []string      `json:"condition_field"`
		AuthUserField  []string      `json:"auth_user_field"`
		AuthOrgScopes  AuthOrgScopes `json:"auth_org_scopes"`
		Errmsg         string        `json:"errmsg"`
	}
	// AuthOrgScopes 权限范围返回
	AuthOrgScopes struct {
		AuthedUser []string `json:"authed_user"`
		AuthedDept []int    `json:"authed_dept"`
	}
)
