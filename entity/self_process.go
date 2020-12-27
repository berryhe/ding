// Package entity  智能工作流/自有
package entity

// 获取模板code
type (
	// SelfGetByNameRequest 获取模板code 请求
	SelfGetByNameRequest struct {
		Name string `json:"name"`
	}

	// SelfGetByNameResp 获取模板code  返回
	SelfGetByNameResp struct {
		Errcode     int    `json:"errcode"`
		ProcessCode string `json:"process_code"`
		RequestID   string `json:"request_id"`
	}
)
