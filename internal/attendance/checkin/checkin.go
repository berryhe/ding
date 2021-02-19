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
	"errors"

	"github.com/berryhe/ding/internal"

	"github.com/berryhe/ding"
	"github.com/berryhe/ding/entity"
)

const (
	apiAttendanceList         = "/attendance/list"
	apiAttendanceListRecord   = "/attendance/listrecord"
	apiAttendanceUploadRecord = "/topapi/attendance/record/upload"
)

// LoopAttendanceCheckinList 如果查询的dingids数量超过50，可帮助循环查询所有
func LoopAttendanceCheckinList(dCtx *ding.DCtx, dingIDs []string, workDateFrom, workDateTo string, isI18N bool) ([]entity.RecordCheckResult, error) {

	var (
		resps    []entity.RecordCheckResult
		resp     entity.AttendanceCheckinListResp
		loopFunc func(entity.AttendanceCheckinListRequest)
		err      error
	)

	arc := entity.AttendanceCheckinListRequest{
		WorkDateFrom: workDateFrom,
		WorkDateTo:   workDateTo,
		UserIDList:   dingIDs,
		Limit:        50,
		Offset:       0,
		IsI18N:       isI18N,
	}

	loopFunc = func(arc entity.AttendanceCheckinListRequest) {
		resp, err = AttendanceCheckinList(dCtx, arc)
		if err != nil {
			return
		}

		resps = append(resps, resp.Recordresult...)
		if resp.HasMore {
			arc.Offset += arc.Limit
			loopFunc(arc)
		}
	}

	// 根据limit限制个数分组，每一组arc.Limit个
	dIDGroups := internal.StrArrGroupAlg(dingIDs, arc.Limit)

	for _, d := range dIDGroups {
		arc.Offset = 0
		arc.UserIDList = d

		loopFunc(arc)

		if err != nil {
			return resps, err
		}

	}

	return resps, err
}

// AttendanceCheckinList 获取打卡结果
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/get-punch-results
// POST https://oapi.dingtalk.com/attendance/list?access_token=ACCESS_TOKEN
func AttendanceCheckinList(dCtx *ding.DCtx, acr entity.AttendanceCheckinListRequest) (resp entity.AttendanceCheckinListResp, err error) {

	if acr.Limit > 50 || acr.Offset < 0 {
		err = errors.New("offset 小于0 或 Limit大于50")
		return
	}

	playLoad, err := json.Marshal(acr)
	if err != nil {
		return resp, err
	}

	data, err := dCtx.HTTPPost(apiAttendanceList, playLoad)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return
	}

	return

}

// AttendanceListRecord 获取打卡详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/attendance-clock-in-record-is-open
// POST https://oapi.dingtalk.com/attendance/listrecord?access_token=ACCESS_TOKEN
func AttendanceListRecord(dCtx *ding.DCtx, alr entity.AttendanceListRecordRequest) (resp entity.AttendanceListRecordResp, err error) {
	playLoad, err := json.Marshal(alr)
	if err != nil {
		return resp, err
	}

	data, err := dCtx.HTTPPost(apiAttendanceListRecord, playLoad)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return
	}

	return
}

// AttendanceRecordUpload 上传打卡记录
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/upload-punch-records
// POST https://oapi.dingtalk.com/topapi/attendance/record/upload?access_token=ACCESS_TOKEN
func AttendanceRecordUpload(dCtx *ding.DCtx, aru entity.AttendanceRecordUploadRequest) (resp entity.AttendanceRecordUploadResp, err error) {
	playLoad, err := json.Marshal(aru)
	if err != nil {
		return resp, err
	}

	data, err := dCtx.HTTPPost(apiAttendanceUploadRecord, playLoad)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &resp); err != nil {
		return
	}

	return
}
