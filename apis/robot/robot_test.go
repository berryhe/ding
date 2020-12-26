package robot

import (
	"testing"

	"github.com/Berry961103/ding"
)

func TestRobot(t *testing.T) {
	config := ding.AppConfig{
		RebotToken: "",
	}
	appCtx := ding.NewApp(config)

	g := MakeDingRobotTextEntity("berry", []string{"12345678"})
	err := SendDingGroupRobot(appCtx, g)
	if err != nil {
		t.Fatal(err)
	}
}
