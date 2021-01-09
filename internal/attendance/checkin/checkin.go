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

// Package checkin 考勤打卡 impl
package checkin

import (
	"encoding/json"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/apis/attendance/checkin"
	"github.com/Berry961103/ding/entity"
)

// LoopAttendanceCheckinList 如果查询的dingids数量超过50，可帮助循环查询所有
func LoopAttendanceCheckinList(dctx *ding.DingCtx, dingIDs []string, workDateFrom, workDateTo string, isI18N bool) ([]entity.AttendanceCheckinListResp, error) {
	arc := entity.AttendanceCheckinListRequest{
		WorkDateFrom: workDateFrom,
		WorkDateTo:   workDateTo,
		UserIDList:   dingIDs,
		Limit:        50,
		Offset:       0,
		IsI18N:       isI18N,
	}

	var resps []entity.AttendanceCheckinListResp

	resp, err := AttendanceCheckinList(dctx, arc)
	if err != nil {
		return nil, err
	}
	resps = append(resps, resp)

	for resp.HasMore {
		arc.Offset += arc.Limit
		resp, err := AttendanceCheckinList(dctx, arc)
		if err != nil {
			return nil, err
		}

		resps = append(resps, resp)
	}
	return resps, nil
}

// AttendanceCheckinList 获取打卡结果
func AttendanceCheckinList(dctx *ding.DingCtx, acr entity.AttendanceCheckinListRequest) (resp entity.AttendanceCheckinListResp, err error) {
	playload, err := json.Marshal(acr)
	if err != nil {
		return resp, err
	}

	data, err := checkin.AttendanceList(dctx, playload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return
	}

	return

}

// AttendanceListRecord 获取打卡详情
func AttendanceListRecord(dctx *ding.DingCtx, alr entity.AttendanceListRecordRequest) (resp entity.AttendanceListRecordResp, err error) {
	playload, err := json.Marshal(alr)
	if err != nil {
		return resp, err
	}

	data, err := checkin.AttendanceListRecord(dctx, playload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return
	}

	return
}

// AttendanceRecordUpload 上传打卡记录
func AttendanceRecordUpload(dctx *ding.DingCtx, aru entity.AttendanceRecordUploadRequest) (resp entity.AttendanceRecordUploadResp, err error) {
	playload, err := json.Marshal(aru)
	if err != nil {
		return resp, err
	}

	data, err := checkin.AttendanceUploadRecord(dctx, playload)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return
	}

	return
}
