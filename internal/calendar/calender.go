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
