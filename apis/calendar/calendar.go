// Package calendar 日程
package calendar

import (
	"github.com/Berry961103/ding"
)

const (
	apiCreate         = "/topapi/calendar/v2/event/create"
	apiUpdate         = "/topapi/calendar/v2/event/update"
	apiCancel         = "/topapi/calendar/v2/event/cancel"
	apiAttendeeUpdate = "/topapi/calendar/v2/attendee/update"
)

// Create 创建日程
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/create-schedule
// POST https://oapi.dingtalk.com/topapi/calendar/v2/event/update?access_token=ACCESS_TOKEN
func Create(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, payload, ding.DefaultPostDecodeStr)
}

// Update 修改日程
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/modify-schedule
// POST https://oapi.dingtalk.com/topapi/calendar/v2/event/update?access_token=ACCESS_TOKEN
func Update(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, payload, ding.DefaultPostDecodeStr)
}

// AttendeeUpdate 修改日程参与者
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/modify-schedule-participant
// POST POST https://oapi.dingtalk.com/topapi/calendar/v2/attendee/update?access_token=ACCESS_TOKEN
func AttendeeUpdate(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAttendeeUpdate, payload, ding.DefaultPostDecodeStr)
}

// Cancel 取消日程
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/cancel-schedule
// POST POST https://oapi.dingtalk.com/topapi/calendar/v2/event/cancel?access_token=ACCESS_TOKEN
func Cancel(ctx *ding.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancel, payload, ding.DefaultPostDecodeStr)
}
