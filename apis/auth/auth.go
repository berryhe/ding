// Package auth 身份验证 interface
package auth

import (
	"net/url"

	"github.com/Berry961103/ding"
)

const (
	apiAuthScopes = "/auth/scopes"
)

// GetAuthScopes 获取通讯录权限范围
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/address-book-permissions
// GET https://oapi.dingtalk.com/auth/scopes?access_token=ACCESS_TOKEN
func GetAuthScopes(ctx *ding.App, params url.Values) (resp []byte, err error) {
	return ctx.HTTPGet(apiAuthScopes + "?" + params.Encode())
}
