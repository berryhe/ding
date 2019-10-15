package ding_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Berry961103/ding"
)

// SendDingRobot Send text messages to ding robot
func SendDingRobot(msg, token string, phone []string) error {
	t := ding.TextType{
		Content: msg,
	}
	a := ding.AtType{
		AtMobiles: phone,
	}
	g := ding.SendGrouptype{
		MsgType: "text",
		Text:    t,
		At:      a,
		IsAtAll: true,
	}

	err := sendGroup(token, g)
	if err != nil {
		return fmt.Errorf("send ding Group err:[%s]", err.Error())
	}
	return nil
}

// SendGroup 发送消息到钉钉群机器人
func sendGroup(token string, sendGroup ding.SendGrouptype) error {

	client := &http.Client{
		Timeout: time.Duration(15) * time.Second,
	}

	reqData, err := json.Marshal(sendGroup)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(reqData)
	resp, err := client.Post("https://oapi.dingtalk.com/robot/send?access_token="+token, "application/json;charset=UTF-8", reader)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respData ding.SendGroupResp
	err = json.Unmarshal(data, &respData)
	if err != nil {
		return err
	}

	if respData.ErrCode != 0 {
		return errors.New(respData.ErrMsg)
	}
	return nil
}
