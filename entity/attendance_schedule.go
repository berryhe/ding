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

// Package entity 考情排班结构体
package entity

// 考情排班
type (
	// AttendanceListScheduleRequest 考情排班详情请求
	AttendanceListScheduleRequest struct {
		WorkDate string `json:"workDate"`
		Offset   int    `json:"offset"`
		Size     int    `json:"size"`
	}
	// AttendanceListScheduleResp 响应体
	AttendanceListScheduleResp struct {
		Errcode   int                  `json:"errcode"`
		Result    AttendanceListResult `json:"result"`
		RequestID string               `json:"request_id"`
	}
	// AttendanceListResult 考情排班
	AttendanceListResult struct {
		HasMore   bool        `json:"has_more"`
		Schedules []Schedules `json:"schedules"`
	}
	// Schedules 考情排班
	Schedules struct {
		CheckType      string `json:"check_type"`
		ClassID        int    `json:"class_id"`
		ClassSettingID int    `json:"class_setting_id"`
		GroupID        int    `json:"group_id"`
		PlanCheckTime  string `json:"plan_check_time"`
		PlanID         int64  `json:"plan_id"`
		Userid         string `json:"userid"`
	}
)

// 成员排班信息
type (
	// AttendanceScheduleListByDayRequest 查询成员排班请求
	AttendanceScheduleListByDayRequest struct {
		OpUserID string `json:"op_user_id"`
		DateTime int64  `json:"date_time"`
		UserID   string `json:"user_id"`
	}
	// AttendanceScheduleListByDayResp 查询成员排班返回
	AttendanceScheduleListByDayResp struct {
		Errcode   int                     `json:"errcode"`
		Result    []AttendanceByDayResult `json:"result"`
		Success   bool                    `json:"success"`
		RequestID string                  `json:"request_id"`
	}
	// AttendanceByDayResult 成员排班信息
	AttendanceByDayResult struct {
		CheckDateTime  string `json:"check_date_time"`
		CheckStatus    string `json:"check_status"`
		CheckType      string `json:"check_type"`
		ClassID        int    `json:"class_id"`
		ClassName      string `json:"class_name"`
		ClassSettingID int    `json:"class_setting_id"`
		CorpID         string `json:"corp_id"`
		Features       string `json:"features"`
		GmtCreate      string `json:"gmt_create"`
		GmtModified    string `json:"gmt_modified"`
		GroupID        int    `json:"group_id"`
		ID             int64  `json:"id"`
		IsRest         string `json:"is_rest"`
		RealPlanTime   string `json:"real_plan_time"`
		UserID         string `json:"user_id"`
		WorkDate       string `json:"work_date"`
	}
)
