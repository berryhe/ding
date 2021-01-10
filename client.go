// MIT License

// Copyright (c) 2019 Berryhe

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package ding http 请求客户端封装
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

// HTTPGet GET 请求
func (dctx *DCtx) HTTPGet(uri string) (resp []byte, err error) {
	uri, err = dctx.applyAccessToken(uri)
	if err != nil {
		return
	}
	return dctx.httpGet(uri)
}

func (dctx *DCtx) httpGet(uri string) (resp []byte, err error) {

	uri = DingdingServerURL + uri
	if dctx.logger != nil {
		dctx.logger.Debugf("GET %s", uri)
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", UserAgent)
	response, err := dctx.httpClient.Do(req)

	if err != nil {
		return
	}

	defer response.Body.Close()
	return responseFilter(response)
}

//HTTPPost POST 请求
func (dctx *DCtx) HTTPPost(uri string, payload []byte, contentType string) (resp []byte, err error) {
	uri, err = dctx.applyAccessToken(uri)
	if err != nil {
		return
	}

	return dctx.httpPost(uri, bytes.NewReader(payload), contentType)
}

// RobotHTTPPost 为钉钉群机器人专门封装一个
func (dctx *DCtx) RobotHTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	url := fmt.Sprintf("%s?access_token=%s", uri, dctx.Config.RobotToken)
	return dctx.httpPost(url, payload, contentType)
}

func (dctx *DCtx) httpPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {

	uri = DingdingServerURL + uri
	if dctx.logger != nil {
		dctx.logger.Debugf("POST %s", uri)
	}

	req, err := http.NewRequest(http.MethodPost, uri, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", contentType)
	response, err := dctx.httpClient.Do(req)

	if err != nil {
		return
	}

	defer response.Body.Close()
	return responseFilter(response)
}

// 在请求地址上附加上 access_token
func (dctx *DCtx) applyAccessToken(oldURL string) (newURL string, err error) {
	accessToken, err := dctx.accessToken.getAccessTokenHandler(dctx.Config.AppKey, dctx.Config.AppSecret)
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
