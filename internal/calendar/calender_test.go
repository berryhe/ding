package calendar

import (
	"testing"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/entity"
)

func TestCreate(t *testing.T) {
	config := ding.AppConfig{
		AgentID:   251419465,
		AppKey:    "ding7ujwcr5ldqpaj8gv",
		AppSecret: "BRkNLDSg_6f8AzrYDSHUMqmsNxxp6mwc4eoFwfFEsvey2Ju51S-ibUUI3wlLRMlZ",
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
