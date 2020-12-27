// Package schedule api interface
package schedule

import "github.com/Berry961103/ding"

const (
	apiListSchedule      = "/topapi/attendance/listschedule"
	apiScheduleListbyday = "/topapi/attendance/schedule/listbyday"
)

// AttendanceListSchedule 查询企业考勤排班详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/query-attendance-scheduling-details
// POST https://oapi.dingtalk.com/topapi/attendance/listschedule?access_token=ACCESS_TOKEN
func AttendanceListSchedule(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiListSchedule, payload, ding.DefaultPostDecodeStr)
}

// AttendanceScheduleListByDay 查询成员排班信息
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/query-member-scheduling-information
// POST https://oapi.dingtalk.com/topapi/attendance/schedule/listbyday?access_token=ACCESS_TOKEN
func AttendanceScheduleListByDay(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiScheduleListbyday, payload, ding.DefaultPostDecodeStr)
}
