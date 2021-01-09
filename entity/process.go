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

// Package entity process 工作流
package entity

import "time"

//  发起审批
type (
	// ProcessCreateRequest 发起审批请求
	ProcessCreateRequest struct {
		FormComponentValues FormComponentValues `json:"form_component_values"`
		AgentID             int                 `json:"agent_id"`
		ProcessCode         string              `json:"process_code"`
		CcPosition          string              `json:"cc_position"`
		Approvers           string              `json:"approvers"`
		CcList              string              `json:"cc_list"`
		DeptID              int                 `json:"dept_id"`
		Approvers2          ApproversV2         `json:"approvers_v2"`
		OriginatorUserID    string              `json:"originator_user_id"`
	}

	// FormComponentValues 审批流表单参数，最大列表长度20。仅支持下表列举的表单控件。
	FormComponentValues struct {
		Name     string `json:"name"`
		Value    string `json:"value"`
		ExtValue string `json:"ext_value"`
	}

	// ApproversV2 审批人列表。支持会签/或签，优先级高于approvers变量
	ApproversV2 struct {
		TaskActionType string `json:"task_action_type"`
		UserIds        string `json:"user_ids"`
	}

	// ProcessCreateResp 发起审批返回
	ProcessCreateResp struct {
		Errcode           int    `json:"errcode"`
		ProcessInstanceID string `json:"process_instance_id"`
		Errmsg            string `json:"errmsg"`
		RequestID         string `json:"request_id"`
	}
)

// 终止审批
type (
	// TerminateProcessRequest 终止审批请求
	TerminateProcessRequest struct {
		TerminateRequest TerminateProcess `json:"request"`
	}
	// TerminateProcess result
	TerminateProcess struct {
		IsSystem          bool   `json:"is_system"`
		ProcessInstanceID string `json:"process_instance_id"`
		OperatingUserid   string `json:"operating_userid"`
		Remark            string `json:"remark"`
	}

	// TerminateProcessResp 终止审批返回
	TerminateProcessResp struct {
		Result    bool   `json:"result"`
		Errcode   int    `json:"errcode"`
		Success   bool   `json:"success"`
		Errmsg    string `json:"errmsg"`
		RequestID string `json:"request_id"`
	}
)

// 批量获取审批实例id
type (

	// ProcessListIDRequest 批量获取审批实例id请求
	ProcessListIDRequest struct {
		EndTime     int64  `json:"end_time"`
		Cursor      int    `json:"cursor"`
		StartTime   int64  `json:"start_time"`
		Size        int    `json:"size"`
		ProcessCode string `json:"process_code"`
		UseridList  string `json:"userid_list"`
	}

	// ProcessListIDResp 批量获取审批实例id返回
	ProcessListIDResp struct {
		Errcode   int                 `json:"errcode"`
		Result    ProcessListIDResult `json:"result"`
		Errmsg    string              `json:"errmsg"`
		RequestID string              `json:"request_id"`
	}

	// ProcessListIDResult 返回消息主体
	ProcessListIDResult struct {
		List []string `json:"list"`
	}
)

// 获取审批实例详情
type (
	// UserProcessRequest 获取审批实例详情请求
	UserProcessRequest struct {
		ProcessInstanceID string `json:"process_instance_id"`
	}

	// UserProcessResp 获取审批实例详情响应
	UserProcessResp struct {
		Errcode         int             `json:"errcode"`
		ProcessInstance ProcessInstance `json:"process_instance"`
		RequestID       string          `json:"request_id"`
	}

	// UPFormComponentValues 表单详情列表。
	UPFormComponentValues struct {
		ComponentType string `json:"component_type"`
		ID            string `json:"id"`
		Name          string `json:"name"`
		Value         string `json:"value"`
	}
	// OperationRecords 操作记录列表。
	OperationRecords struct {
		Date            string `json:"date"`
		OperationResult string `json:"operation_result"`
		OperationType   string `json:"operation_type"`
		Userid          string `json:"userid"`
	}

	// ProcessInstance 审批实例详情
	ProcessInstance struct {
		AttachedProcessInstanceIds []string                `json:"attached_process_instance_ids"`
		BizAction                  string                  `json:"biz_action"`
		BusinessID                 string                  `json:"business_id"`
		CreateTime                 string                  `json:"create_time"`
		FinishTime                 string                  `json:"finish_time"`
		FormComponentValues        []UPFormComponentValues `json:"form_component_values"`
		OperationRecords           []OperationRecords      `json:"operation_records"`
		OriginatorDeptID           string                  `json:"originator_dept_id"`
		OriginatorDeptName         string                  `json:"originator_dept_name"`
		OriginatorUserid           string                  `json:"originator_userid"`
		// 	审批结果：
		// agree：同意
		// refuse：拒绝
		Result string `json:"result"`
		// 审批状态：
		// NEW：新创建
		// RUNNING：审批中
		// TERMINATED：被终止
		// COMPLETED：完成
		// CANCELED：取消
		Status string      `json:"status"`
		Tasks  []TaskTopVo `json:"tasks"`
		Title  string      `json:"title"`
	}

	// TaskTopVo 任务列表。
	TaskTopVo struct {
		UserID string `json:"userid"`
		// 		任务状态：
		// NEW：未启动
		// RUNNING：处理中
		// PAUSED：暂停
		// CANCELED：取消
		// COMPLETED：完成
		// TERMINATED：终止
		TaskStatus string    `json:"task_status"`
		TaskResult string    `json:"task_result"`
		CreateTime time.Time `json:"create_time"`
		FinishTime time.Time `json:"finish_time"`
		TaskID     string    `json:"task_id"`
		URL        string    `json:"url"`
	}
)

// 获取用户待审批数量
type (

	// ProcessTodoNumRequest 获取用户待审批数量请求
	ProcessTodoNumRequest struct {
		Userid string `json:"userid"`
	}

	// ProcessTodoNumResp  获取用户待审批数量响应
	ProcessTodoNumResp struct {
		Count     int    `json:"count"`
		Errcode   int    `json:"errcode"`
		Errmsg    string `json:"errmsg"`
		RequestID string `json:"request_id"`
	}
)

type (
	// ProcessListByUserIDRequest 获取指定用户可见的审批表单列表请求
	ProcessListByUserIDRequest struct {
		UserID string `json:"userid"`
		Offset int    `json:"offset"`
		Size   string `json:"size"`
	}

	// ProcessListByUserIDResp 获取指定用户可见的审批表单列表返回
	ProcessListByUserIDResp struct {
		Errcode   int               `json:"errcode"`
		Errmsg    string            `json:"errmsg"`
		Result    ProcessListResult `json:"result"`
		RequestID string            `json:"request_id"`
	}

	// ProcessList 可见表单列表。
	ProcessList struct {
		IconURL     string `json:"icon_url"`
		Name        string `json:"name"`
		ProcessCode string `json:"process_code"`
		URL         string `json:"url"`
	}

	// ProcessListResult 可见表单列表 查询结果。
	ProcessListResult struct {
		NextCursor  int           `json:"next_cursor"`
		ProcessList []ProcessList `json:"process_list"`
	}
)

// 获取审批钉盘空间信息
type (
	// ProcessCspaceInfoRequest 获取审批钉盘空间信息请求
	ProcessCspaceInfoRequest struct {
		UserID string `json:"user_id"`
	}

	// ProcessCspaceInfoResp 获取审批钉盘空间信息返回
	ProcessCspaceInfoResp struct {
		Errcode   int                 `json:"errcode"`
		Result    ProcessCspaceResult `json:"result"`
		Success   bool                `json:"success"`
		Errmsg    string              `json:"errmsg"`
		RequestID string              `json:"request_id"`
	}

	// ProcessCspaceResult 返回结果。
	ProcessCspaceResult struct {
		SpaceID int64 `json:"space_id"`
	}
)

type (
	// ProcessCspacePreviewRequest 授权预览审批附件 请求
	ProcessCspacePreviewRequest struct {
		Request PCPreviewRequest `json:"request"`
	}

	// PCPreviewRequest 授权预览审批附件请求主体
	PCPreviewRequest struct {
		Agentid           int    `json:"agentid"`
		FileID            string `json:"file_id"`
		ProcessInstanceID string `json:"process_instance_id"`
		Userid            string `json:"userid"`
	}

	// ProcessCspacePreviewResp 授权预览审批附件 返回
	ProcessCspacePreviewResp struct {
		Result    PCRespResult `json:"result"`
		Errcode   int          `json:"errcode"`
		RequestID string       `json:"request_id"`
	}

	// PCRespResult 授权预览审批附件返回主体
	PCRespResult struct {
		SpaceID string `json:"space_id"`
	}
)
