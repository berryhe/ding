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

// Package calendar 日程 interface
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
func Create(ctx *ding.DingCtx, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiCreate, payload, ding.DefaultPostDecodeStr)
}

// Update 修改日程
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/modify-schedule
// POST https://oapi.dingtalk.com/topapi/calendar/v2/event/update?access_token=ACCESS_TOKEN
func Update(ctx *ding.DingCtx, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiUpdate, payload, ding.DefaultPostDecodeStr)
}

// AttendeeUpdate 修改日程参与者
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/modify-schedule-participant
// POST POST https://oapi.dingtalk.com/topapi/calendar/v2/attendee/update?access_token=ACCESS_TOKEN
func AttendeeUpdate(ctx *ding.DingCtx, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiAttendeeUpdate, payload, ding.DefaultPostDecodeStr)
}

// Cancel 取消日程
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/cancel-schedule
// POST POST https://oapi.dingtalk.com/topapi/calendar/v2/event/cancel?access_token=ACCESS_TOKEN
func Cancel(ctx *ding.DingCtx, payload []byte) (resp []byte, err error) {
	return ctx.HTTPPost(apiCancel, payload, ding.DefaultPostDecodeStr)
}
