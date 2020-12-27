package robot

import (
	"testing"

	"github.com/Berry961103/ding"
)

func TestRobotText(t *testing.T) {
	config := ding.AppConfig{
		RobotToken: "",
	}
	appCtx := ding.NewApp(config)

	err := DingRobotText(appCtx, "berry", []string{"12345678"})
	if err != nil {
		t.Fatal(err)
	}
}
