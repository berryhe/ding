// Package robot implements top
package robot

import (
	"encoding/json"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/apis/robot"
	"github.com/Berry961103/ding/entity"
)

// DingRobotText Send ding robot text
func DingRobotText(ctx *ding.App, msg string, phone []string) error {

	sg := entity.SendGrouptype{
		MsgType: "text",
		Text: entity.TextType{
			Content: msg,
		},
		At: entity.AtType{
			AtMobiles: phone,
		},
		IsAtAll: true,
	}

	reqData, err := json.Marshal(sg)
	if err != nil {
		return err
	}

	return robot.SendDingGroupRobot(ctx, reqData)
}
