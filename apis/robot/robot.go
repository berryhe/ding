package robot

import (
	"bytes"

	"github.com/Berry961103/ding"
)

const (
	robotAPI = "/robot/send"
)

// SendDingGroupRobot 发送消息到钉钉群机器人
func SendDingGroupRobot(ctx *ding.App, payload []byte) error {

	_, err := ctx.RobotHTTPPost(robotAPI, bytes.NewReader(payload), ding.DefaultPostDecodeStr)

	return err
}
