// Package checkin 考情打卡 interface
package checkin

import "github.com/Berry961103/ding"

const (
	apiAttendanceList       = "/attendance/list"
	apiAttendanceListRecord = "/attendance/listRecord"
)

// AttendanceList 获取打卡结果
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/get-punch-results
// POST https://oapi.dingtalk.com/attendance/list?access_token=ACCESS_TOKEN
func AttendanceList(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiAttendanceList, payload, ding.DefaultPostDecodeStr)
}

// AttendanceListRecord 获取打卡详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/attendance-clock-in-record-is-open
// POST https://oapi.dingtalk.com/attendance/listrecord?access_token=ACCESS_TOKEN
func AttendanceListRecord(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiAttendanceList, payload, ding.DefaultPostDecodeStr)
}
