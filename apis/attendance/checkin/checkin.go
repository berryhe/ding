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

// Package checkin 考情打卡 interface
package checkin

import "github.com/Berry961103/ding"

const (
	apiAttendanceList         = "/attendance/list"
	apiAttendanceListRecord   = "/attendance/listRecord"
	apiAttendanceUploadRecord = "/topapi/attendance/record/upload"
)

// AttendanceList 获取打卡结果
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/get-punch-results
// POST https://oapi.dingtalk.com/attendance/list?access_token=ACCESS_TOKEN
func AttendanceList(dctx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dctx.HTTPPost(apiAttendanceList, payload, ding.DefaultPostDecodeStr)
}

// AttendanceListRecord 获取打卡详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/attendance-clock-in-record-is-open
// POST https://oapi.dingtalk.com/attendance/listrecord?access_token=ACCESS_TOKEN
func AttendanceListRecord(dctx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dctx.HTTPPost(apiAttendanceList, payload, ding.DefaultPostDecodeStr)
}

// AttendanceUploadRecord 上传打卡记录
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/upload-punch-records
// POST https://oapi.dingtalk.com/topapi/attendance/record/upload?access_token=ACCESS_TOKEN
func AttendanceUploadRecord(dctx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dctx.HTTPPost(apiAttendanceUploadRecord, payload, ding.DefaultPostDecodeStr)
}
