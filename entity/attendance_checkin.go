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

// Package entity 考情打卡结构体
package entity

// RecordCheckResult AttendanceCheckinListResp 考情打卡结果
type (
	// AttendanceCheckinListRequest 考情打卡请求
	AttendanceCheckinListRequest struct {
		WorkDateFrom string   `json:"workDateFrom"`
		Offset       int      `json:"offset"`
		UserIDList   []string `json:"userIdList"`
		Limit        int      `json:"limit"`
		IsI18N       bool     `json:"isI18n"`
		WorkDateTo   string   `json:"workDateTo"`
	}
	// AttendanceCheckinListResp 考情打卡返回
	AttendanceCheckinListResp struct {
		Errcode      int                 `json:"errcode"`
		Recordresult []RecordCheckResult `json:"recordresult"`
		Record       string              `json:"_record"`
		HasMore      bool                `json:"hasMore"`
		Errmsg       string              `json:"errmsg"`
	}
	RecordCheckResult struct {
		CheckType      string `json:"checkType"`
		CorpID         string `json:"corpId"`
		LocationResult string `json:"locationResult"`
		BaseCheckTime  int64  `json:"baseCheckTime"`
		GroupID        int    `json:"groupId"`
		TimeResult     string `json:"timeResult"`
		UserID         string `json:"userId"`
		RecordID       int64  `json:"recordId"`
		WorkDate       int64  `json:"workDate"`
		SourceType     string `json:"sourceType"`
		UserCheckTime  int64  `json:"userCheckTime"`
		PlanID         int64  `json:"planId"`
		ID             int64  `json:"id"`
	}
)

// ALRecordresult 考情打卡详情
type (
	// AttendanceListRecordRequest 考情打卡详情请求
	AttendanceListRecordRequest struct {
		CheckDateFrom string   `json:"checkDateFrom"`
		UserIDs       []string `json:"userIds"`
		IsI18N        string   `json:"isI18n"`
		CheckDateTo   string   `json:"checkDateTo"`
	}
	// AttendanceListRecordResp 考情打卡详情 返回
	AttendanceListRecordResp struct {
		Errcode      int              `json:"errcode"`
		Errmsg       string           `json:"errmsg"`
		Recordresult []ALRecordresult `json:"recordresult"`
	}
	ALRecordresult struct {
		GmtModified       int64   `json:"gmtModified"`
		BaseCheckTime     int64   `json:"baseCheckTime,omitempty"`
		GroupID           int     `json:"groupId,omitempty"`
		TimeResult        string  `json:"timeResult,omitempty"`
		DeviceID          string  `json:"deviceId"`
		ApproveID         int64   `json:"approveId,omitempty"`
		UserAccuracy      int     `json:"userAccuracy"`
		ClassID           int     `json:"classId,omitempty"`
		WorkDate          int64   `json:"workDate"`
		BizID             string  `json:"bizId"`
		PlanID            int64   `json:"planId,omitempty"`
		ID                int64   `json:"id"`
		CheckType         string  `json:"checkType,omitempty"`
		PlanCheckTime     int64   `json:"planCheckTime,omitempty"`
		CorpID            string  `json:"corpId"`
		LocationResult    string  `json:"locationResult,omitempty"`
		UserLongitude     float64 `json:"userLongitude"`
		IsLegal           string  `json:"isLegal,omitempty"`
		ProcInstID        string  `json:"procInstId,omitempty"`
		GmtCreate         int64   `json:"gmtCreate"`
		UserID            string  `json:"userId"`
		OutsideRemark     string  `json:"outsideRemark,omitempty"`
		UserAddress       string  `json:"userAddress"`
		UserLatitude      float64 `json:"userLatitude"`
		SourceType        string  `json:"sourceType"`
		UserCheckTime     int64   `json:"userCheckTime"`
		LocationMethod    string  `json:"locationMethod"`
		InvalidRecordType string  `json:"invalidRecordType,omitempty"`
		InvalidRecordMsg  string  `json:"invalidRecordMsg,omitempty"`
	}
)

// 上传打卡记录 请求与响应实体
type (
	// AttendanceRecordUploadRequest 上传打卡记录 请求
	AttendanceRecordUploadRequest struct {
		DeviceName    string `json:"device_name"`
		DeviceID      string `json:"device_id"`
		UserCheckTime int64  `json:"user_check_time"`
		PhotoURL      string `json:"photo_url"`
		Userid        string `json:"userid"`
	}

	// AttendanceRecordUploadResp 上传打卡记录 响应
	AttendanceRecordUploadResp struct {
		Errcode   int    `json:"errcode"`
		Success   bool   `json:"success"`
		Errmsg    string `json:"errmsg"`
		RequestID string `json:"request_id"`
	}
)
