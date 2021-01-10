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

// Package schedule api interface
package schedule

import "github.com/Berry961103/ding"

const (
	apiListSchedule            = "/topapi/attendance/listschedule"
	apiScheduleListByDay       = "/topapi/attendance/schedule/listbyday"
	apiScheduleListByUsers     = "/topapi/attendance/schedule/listbyusers"
	apiScheduleGroupASync      = "/topapi/attendance/group/schedule/async"
	apiScheduleResultListByIDs = "/topapi/attendance/schedule/result/listbyids"
)

// AttendanceListSchedule 查询企业考勤排班详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/query-attendance-scheduling-details
// POST https://oapi.dingtalk.com/topapi/attendance/listschedule?access_token=ACCESS_TOKEN
func AttendanceListSchedule(dCtx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dCtx.HTTPPost(apiListSchedule, payload, ding.DefaultPostDecodeStr)
}

// AttendanceScheduleListByDay 查询成员排班信息
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/query-member-scheduling-information
// POST https://oapi.dingtalk.com/topapi/attendance/schedule/listbyday?access_token=ACCESS_TOKEN
func AttendanceScheduleListByDay(dCtx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dCtx.HTTPPost(apiScheduleListByDay, payload, ding.DefaultPostDecodeStr)
}

// AttendanceScheduleListByUsers 批量查询人员排班信息
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/batch-query-of-personnel-scheduling-information
// POST POST https://oapi.dingtalk.com/topapi/attendance/schedule/listbyusers?access_token=ACCESS_TOKEN
func AttendanceScheduleListByUsers(dCtx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dCtx.HTTPPost(apiScheduleListByUsers, payload, ding.DefaultPostDecodeStr)
}

// AttendanceScheduleGroupASync 排班考勤组排班
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/class-scheduling-attendance-group
// POST https://oapi.dingtalk.com/topapi/attendance/group/schedule/async?access_token=ACCESS_TOKEN
func AttendanceScheduleGroupASync(dCtx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dCtx.HTTPPost(apiScheduleGroupASync, payload, ding.DefaultPostDecodeStr)
}

// AttendanceScheduleResultListByIDs 查询排班打卡结果
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/query-the-punch-in-result-of-shift-scheduling
// POST https://oapi.dingtalk.com/topapi/attendance/schedule/result/listbyids?access_token=ACCESS_TOKEN
func AttendanceScheduleResultListByIDs(dCtx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dCtx.HTTPPost(apiScheduleResultListByIDs, payload, ding.DefaultPostDecodeStr)
}
