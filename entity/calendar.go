// Package entity 日程
package entity

// 日程的创建
type (
	// CalendarCreateRequest 日程创建请求
	CalendarCreateRequest struct {
		AgentID string `json:"agentid"`
		Event   Event  `json:"event"`
	}

	// CalendarCreateResp 日程创建返回
	CalendarCreateResp struct {
		Errcode   int                  `json:"errcode"`
		Errmsg    string               `json:"errmsg"`
		Result    CalendarCreateResult `json:"result"`
		Success   bool                 `json:"success"`
		RequestID string               `json:"request_id"`
	}
	// CalendarCreateResult 创建日程的消息主体
	CalendarCreateResult struct {
		Attendees        Attendees `json:"attendees"`
		CalendarID       string    `json:"calendar_id"`
		Description      string    `json:"description"`
		End              Time      `json:"end"`
		EventID          string    `json:"event_id"`
		Location         Location  `json:"location"`
		NotificationType string    `json:"notification_type"`
		Organizer        Organizer `json:"organizer"`
		Reminder         Reminder  `json:"reminder"`
		Start            Time      `json:"start"`
		Summary          string    `json:"summary"`
	}
)

// 日程的修改
type (
	// CalendarUpdateRequest 日程更新返回
	CalendarUpdateRequest struct {
		AgentID string `json:"agentid"`
		Event   Event  `json:"event"`
	}
	// Event 消息主体
	Event struct {
		EventID     string      `json:"event_id"`
		Summary     string      `json:"summary"`
		CalendarID  string      `json:"calendar_id"`
		Description string      `json:"description"`
		Attendees   []Attendees `json:"attendees"`
		Reminder    Reminder    `json:"reminder"`
		Organizer   Organizer   `json:"organizer"`
		Start       Time        `json:"start"`
		End         Time        `json:"end"`
		Location    Location    `json:"location"`
	}
	// Reminder 会议开始前提醒
	Reminder struct {
		Method  string `json:"method"`
		Minutes string `json:"minutes"`
	}
	// Time 开始与结束时间
	Time struct {
		Timezone  string `json:"timezone"`
		Timestamp int    `json:"timestamp"`
	}
	// Organizer 日程组织者信息。
	Organizer struct {
		Userid string `json:"userid"`
	}
	// Location 地址信息
	Location struct {
		Latitude  string `json:"latitude"`
		Place     string `json:"place"`
		Longitude string `json:"longitude"`
	}
	// Attendees 日程参与者，参与者最大人数为100，包括组织者
	Attendees struct {
		UserID         string `json:"userid"`
		AttendeeStatus string `json:"attendee_status"`
	}
	// CalendarUpdateResp 日程更新返回
	CalendarUpdateResp struct {
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		Success   bool   `json:"success"`
		RequestID string `json:"request_id"`
	}
)

type (
	// CalendarAttendeeUpdateRequest 更新日程参与者请求
	CalendarAttendeeUpdateRequest struct {
		AgentID    int         `json:"agentid"`
		EventID    string      `json:"event_id"`
		CalendarID string      `json:"calendar_id"`
		Attendees  []Attendees `json:"attendees"`
	}
	// CalendarAttendeeUpdateResp 更新日程参与者返回
	CalendarAttendeeUpdateResp struct {
		Errcode   int    `json:"errcode"`
		Errmsg    string `json:"errmsg"`
		Success   bool   `json:"success"`
		RequestID string `json:"request_id"`
	}
)

type (
	// CalendarCancelRequest 日程取消结请求
	CalendarCancelRequest struct {
		AgentID    int    `json:"agentid"`
		EventID    string `json:"event_id"`
		CalendarID string `json:"calendar_id"`
	}
	// CalendarCancelRespo 日程取消返回
	CalendarCancelRespo struct {
		Errcode   int    `json:"errcode"`
		Errmsg    string `json:"errmsg"`
		Success   bool   `json:"success"`
		RequestID string `json:"request_id"`
	}
)
