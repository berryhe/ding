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
