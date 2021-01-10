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

// Package calendar implements top
package calendar

import (
	"encoding/json"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/entity"
)

const (
	apiCreate         = "/topapi/calendar/v2/event/create"
	apiUpdate         = "/topapi/calendar/v2/event/update"
	apiCancel         = "/topapi/calendar/v2/event/cancel"
	apiAttendeeUpdate = "/topapi/calendar/v2/attendee/update"
)

// Create 钉日程的创建
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/create-schedule
// POST https://oapi.dingtalk.com/topapi/calendar/v2/event/update?access_token=ACCESS_TOKEN
func Create(ctx *ding.DCtx, calendarCreate entity.CalendarCreateRequest) (calenderCreateResp entity.CalendarCreateResp, err error) {
	playload, err := json.Marshal(calendarCreate)
	if err != nil {
		return
	}

	resp, err := ctx.HTTPPost(apiCreate, playload, ding.DefaultPostDecodeStr)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calenderCreateResp); err != nil {
		return
	}
	return
}

// Update 钉日程的修改
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/create-schedule
// POST https://oapi.dingtalk.com/topapi/calendar/v2/event/update?access_token=ACCESS_TOKEN
func Update(ctx *ding.DCtx, calendarUpdate entity.CalendarUpdateRequest) (calendarUpdateResp entity.CalendarUpdateResp, err error) {
	playload, err := json.Marshal(calendarUpdate)
	if err != nil {
		return
	}

	resp, err := ctx.HTTPPost(apiUpdate, playload, ding.DefaultPostDecodeStr)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calendarUpdateResp); err != nil {
		return
	}
	return
}

// AttendeeUpdate 修改日程参与者
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/modify-schedule-participant
// POST POST https://oapi.dingtalk.com/topapi/calendar/v2/attendee/update?access_token=ACCESS_TOKEN
func AttendeeUpdate(ctx *ding.DCtx, calendarAtUpdate entity.CalendarAttendeeUpdateRequest) (calendarUpdateAtResp entity.CalendarAttendeeUpdateResp, err error) {
	playload, err := json.Marshal(calendarAtUpdate)
	if err != nil {
		return
	}

	resp, err := ctx.HTTPPost(apiAttendeeUpdate, playload, ding.DefaultPostDecodeStr)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calendarUpdateAtResp); err != nil {
		return
	}
	return
}

// Cancel 钉日程的取消
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/cancel-schedule
// POST POST https://oapi.dingtalk.com/topapi/calendar/v2/event/cancel?access_token=ACCESS_TOKEN
func Cancel(ctx *ding.DCtx, calendarCancel entity.CalendarCancelRequest) (calendarCancelResp entity.CalendarCancelResp, err error) {
	playload, err := json.Marshal(calendarCancel)
	if err != nil {
		return
	}

	resp, err := ctx.HTTPPost(apiCancel, playload, ding.DefaultPostDecodeStr)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calendarCancelResp); err != nil {
		return
	}
	return
}
