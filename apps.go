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

package ding

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// 从缓存获取 access_token
// 如果没有 access_token 或者已过期，那就刷新
func (dCtx *DCtx) getAccessToken(appKey, appSecretKey string) (accessToken string, err error) {

	cacheKey := fmt.Sprintf("ding_access_token:%s", appKey)

	accessToken, err = dCtx.accessToken.cache.Fetch(cacheKey)
	if len(accessToken) != 0 {
		return
	}

	var expiresIn int

	accessToken, expiresIn, err = dCtx.refreshAccessToken(appKey, appSecretKey)
	if err != nil {
		return
	}

	// 提前 5min 过期 提供冗余时间
	// Token 有效期为 2 小时，在此期间调用该接口 token 不会改变。当 token 有效期小于 10 分的时候，再次请求获取 token 的时候，会生成一个新的 token，与此同时老的 token 依然有效。
	d := time.Duration(expiresIn-300) * time.Second
	_ = dCtx.accessToken.cache.Save(cacheKey, accessToken, d)

	if dCtx.logger != nil {
		dCtx.logger.Debugf("%s %s %d\n", "refreshAccessTokenFromServer", accessToken, expiresIn)
	}

	return
}

// refreshAccessToken 从服务器获取新的
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-access_token
// GET https://oapi.dingtalk.com/gettoken?dctxkey=dctxkey&dctxsecret=dctxsecret
func (dCtx *DCtx) refreshAccessToken(appKey, appSecretKey string) (accessToken string, expiresIn int, err error) {

	params := url.Values{}
	params.Add("appkey", appKey)
	params.Add("appsecret", appSecretKey)

	apiGetToken := fmt.Sprintf("%s/gettoken?%s", DingdingServerURL, params.Encode())
	response, err := dCtx.httpClient.Get(apiGetToken)
	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("ding apis：%shttp response code：%s", apiGetToken, response.Status)
		return
	}

	var result = struct {
		ErrCode     int    `json:"errcode"`
		AccessToken string `json:"access_token"`
		ErrMsg      string `json:"errmsg"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return
	}

	if len(result.AccessToken) == 0 || result.ErrCode != 0 {
		err = fmt.Errorf("%s", result.ErrMsg)
		return
	}

	return result.AccessToken, result.ExpiresIn, nil
}
