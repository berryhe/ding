package calendar

import (
	"testing"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/entity"
)

func TestCreate(t *testing.T) {
	config := ding.AppConfig{
		AgentID:   12345678,
		AppKey:    "dingappkey",
		AppSecret: "dingappSecretKey",
	}
	appCtx := ding.NewApp(config)

	calendarCreate := entity.CalendarCreateRequest{}
	resp, err := Create(appCtx, calendarCreate)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(resp)
}
