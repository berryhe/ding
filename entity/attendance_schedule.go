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
