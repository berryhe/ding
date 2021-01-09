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
	"github.com/Berry961103/ding/apis/calendar"
	"github.com/Berry961103/ding/entity"
)

// Create 钉日程的创建
func Create(ctx *ding.App, calendarCreate entity.CalendarCreateRequest) (calenderCreateResp entity.CalendarCreateResp, err error) {
	playload, err := json.Marshal(calendarCreate)
	if err != nil {
		return
	}

	resp, err := calendar.Create(ctx, playload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calenderCreateResp); err != nil {
		return
	}
	return
}

// Update 钉日程的修改
func Update(ctx *ding.App, calendarUpdate entity.CalendarUpdateRequest) (calendarUpdateResp entity.CalendarUpdateResp, err error) {
	playload, err := json.Marshal(calendarUpdate)
	if err != nil {
		return
	}

	resp, err := calendar.Update(ctx, playload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calendarUpdateResp); err != nil {
		return
	}
	return
}

// AttendeeUpdate 修改日程参与者
func AttendeeUpdate(ctx *ding.App, calendarAtUpdate entity.CalendarAttendeeUpdateRequest) (calendarUpdateAtResp entity.CalendarAttendeeUpdateResp, err error) {
	playload, err := json.Marshal(calendarAtUpdate)
	if err != nil {
		return
	}

	resp, err := calendar.AttendeeUpdate(ctx, playload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calendarUpdateAtResp); err != nil {
		return
	}
	return
}

// Cancel 钉日程的取消
func Cancel(ctx *ding.App, calendarCancel entity.CalendarCancelRequest) (calendarCancelResp entity.CalendarCancelRespo, err error) {
	playload, err := json.Marshal(calendarCancel)
	if err != nil {
		return
	}

	resp, err := calendar.Cancel(ctx, playload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(resp, &calendarCancelResp); err != nil {
		return
	}
	return
}
