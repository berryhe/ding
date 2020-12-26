package ding

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// DefaultPostDecodeStr 默认的post传输 编码格式
	DefaultPostDecodeStr = "application/json;charset=utf-8"
)

var (

	// UserAgent 默认 UserAgent
	UserAgent = "berry-go-client"

	// DingdingServerURL 钉钉 api 服务器地址
	DingdingServerURL = "https://oapi.dingtalk.com"
)

// Client 用于向接口发送请求
type Client struct {
	Ctx *App
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(uri string) (resp []byte, err error) {
	uri, err = client.applyAccessToken(uri)
	if err != nil {
		return
	}
	return client.httpGet(uri)
}

func (client *Client) httpGet(uri string) (resp []byte, err error) {

	uri = DingdingServerURL + uri
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Debugf("GET %s", uri)
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", UserAgent)
	response, err := client.Ctx.httpClient.Do(req)

	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload []byte, contentType string) (resp []byte, err error) {
	uri, err = client.applyAccessToken(uri)
	if err != nil {
		return
	}

	return client.httpPost(uri, bytes.NewReader(payload), contentType)
}

// RobotHTTPPost 为钉钉群机器人专门封装一个
func (client *Client) RobotHTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	return client.httpPost(uri, payload, contentType)
}

func (client *Client) httpPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {

	uri = DingdingServerURL + uri
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Debugf("POST %s", uri)
	}

	req, err := http.NewRequest(http.MethodPost, uri, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", contentType)
	response, err := client.Ctx.httpClient.Do(req)

	if err != nil {
		return
	}

	defer response.Body.Close()
	return responseFilter(response)
}

// 在请求地址上附加上 access_token
func (client *Client) applyAccessToken(oldURL string) (newURL string, err error) {
	accessToken, err := client.Ctx.accessToken.GetAccessTokenHandler()
	if err != nil {
		return
	}

	parse, err := url.Parse(oldURL)
	if err != nil {
		return
	}

	newURL = parse.Host + parse.Path + "?access_token=" + accessToken
	if len(parse.RawQuery) > 0 {
		newURL += "&" + parse.RawQuery
	}

	if len(parse.Fragment) > 0 {
		newURL += "#" + parse.Fragment
	}
	return
}

/*
筛查钉钉 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("http Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if bytes.HasPrefix(resp, []byte(`{`)) { // 只针对 json
		errorResponse := struct {
			Errcode int64  `json:"errcode"`
			Errmsg  string `json:"errmsg"`
		}{}
		err = json.Unmarshal(resp, &errorResponse)
		if err != nil {
			return
		}

		if errorResponse.Errcode != 0 {
			err = errors.New(errorResponse.Errmsg)
			return
		}
	}
	return
}
