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

// Package leave 假勤审批-> 查询请假状态
package leave

import (
	"encoding/json"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/apis/attendance/checkin"
	"github.com/Berry961103/ding/entity"
)

// AttendanceGetLeaveStatus 查询请假状态
func AttendanceGetLeaveStatus(dctx *ding.DingCtx, alr entity.AttendanceLeaveRequest) (resp entity.AttendanceLeaveResp, err error) {
	playload, err := json.Marshal(alr)
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
