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

package calendar

import (
	"testing"
	"time"

	"github.com/berryhe/ding"
	"github.com/berryhe/ding/entity"
)

//AgentId：
//AppKey：
//AppSecret：
func TestCreate(t *testing.T) {
	config := ding.Config{

		AgentID:      1234567890,
		AppKey:       "dkasdasfdsgdfh",
		AppSecretKey: "AppSecretKey",
	}
	appCtx := ding.NewDCtx(config)

	calendarCreate := entity.CalendarCreateRequest{
		AgentID: "12321214",

		Event: entity.Event{
			Summary:    "测试会议",
			CalendarID: "primary",

			Reminder: entity.Reminder{
				Method:  "app",
				Minutes: 5,
			},
			Attendees: []entity.AttendeesType{
				entity.AttendeesType{
					UserID: "016262302",
				},
				entity.AttendeesType{
					UserID: "2312fdsf",
				},
				entity.AttendeesType{
					UserID: "asdasdvdv",
				},
				entity.AttendeesType{
					UserID: "sadasd21321",
				},
			},
			Organizer: entity.Organizer{
				Userid: "asdasd4213",
			},
			Start: entity.Time{
				Timezone:  "Asia/Shanghai",
				Timestamp: time.Now().Add(10 * time.Minute).Unix(),
			},
			End: entity.Time{
				Timezone:  "Asia/Shanghai",
				Timestamp: time.Now().Add(30 * time.Minute).Unix(),
			},
			Location: entity.Location{
				Place: "测试",
			},
		},
	}
	resp, err := Create(appCtx, calendarCreate)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(resp)
}
