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

package entity

// 查询请假状态 实体
type (
	// AttendanceLeaveRequest 查询请假状态请求
	AttendanceLeaveRequest struct {
		StartTime  int64  `json:"start_time"`
		Offset     int    `json:"offset"`
		Size       int    `json:"size"`
		EndTime    int64  `json:"end_time"`
		UseridList string `json:"userid_list"`
	}

	// AttendanceLeaveResp 查询请假状态返回
	AttendanceLeaveResp struct {
		Errcode   int         `json:"errcode"`
		Result    LeaveResult `json:"result"`
		Success   bool        `json:"success"`
		RequestID string      `json:"request_id"`
	}

	// LeaveResult 请假状态返回
	LeaveResult struct {
		HasMore     bool          `json:"has_more"`
		LeaveStatus []LeaveStatus `json:"leave_status"`
	}

	// LeaveStatus 请假状态
	LeaveStatus struct {
		DurationPercent int    `json:"duration_percent"`
		DurationUnit    string `json:"duration_unit"`
		EndTime         int64  `json:"end_time"`
		StartTime       int64  `json:"start_time"`
		Userid          string `json:"userid"`
	}
)
