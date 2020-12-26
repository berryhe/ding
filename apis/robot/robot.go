package robot

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Berry961103/ding"
	"github.com/Berry961103/ding/entity"
)

const (
	robotAPI = "/robot/send"
)

// MakeDingRobotTextEntity Send text struct
func MakeDingRobotTextEntity(msg string, phone []string) entity.SendGrouptype {

	return entity.SendGrouptype{
		MsgType: "text",
		Text: entity.TextType{
			Content: msg,
		},
		At: entity.AtType{
			AtMobiles: phone,
		},
		IsAtAll: true,
	}
}

// SendDingGroupRobot 发送消息到钉钉群机器人
func SendDingGroupRobot(ctx *ding.App, sendGroup entity.SendGrouptype) error {

	reqData, err := json.Marshal(sendGroup)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s?access_token=%s", ding.DingdingServerURL, robotAPI, ctx.Config.RebotToken)
	_, err = ctx.Client.RobotHTTPPost(url, bytes.NewReader(reqData), ding.DefaultPostDecodeStr)

	return err
}
