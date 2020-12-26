package ding

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// NewApp 初始化项
func NewApp(config AppConfig) (app *App) {
	app = newApp(config)
	app.accessToken.getAccessTokenHandler = app.GetAccessToken
	return
}

// GetAccessToken 获取token
// 从缓存获取 access_token
// 如果没有 access_token 或者 已过期，那么刷新
func (app *App) GetAccessToken() (accessToken string, err error) {

	cacheKey := "access_token:" + app.Config.AppKey
	accessToken, err = app.accessToken.cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	var expiresIn int

	accessToken, expiresIn, err = app.refreshAccessToken()
	if err != nil {
		return
	}

	// 提前 5min 过期 提供冗余时间
	// Token 有效期为 2 小时，在此期间调用该接口 token 不会改变。当 token 有效期小于 10 分的时候，再次请求获取 token 的时候，会生成一个新的 token，与此同时老的 token 依然有效。
	d := time.Duration(expiresIn-300) * time.Second
	_ = app.accessToken.cache.Save(cacheKey, accessToken, d)

	if app.logger != nil {
		app.logger.Debugf("%s %s %d\n", "refreshAccessTokenFromServer", accessToken, expiresIn)
	}

	return
}

// refreshAccessToken 从服务器获取新的
// See: https://ding-doc.dingtalk.com/document#/org-dev-guide/obtain-access_token
// GET https://oapi.dingtalk.com/gettoken?appkey=appkey&appsecret=appsecret
func (app *App) refreshAccessToken() (accessToken string, expiresIn int, err error) {

	params := url.Values{}
	params.Add("appkey", app.Config.AppKey)
	params.Add("appsecret", app.Config.AppSecret)

	apiGetToken := fmt.Sprintf("%s/gettoken?%s", DingdingServerURL, params.Encode())
	response, err := app.httpClient.Get(apiGetToken)
	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("GET %s RETURN %s", apiGetToken, response.Status)
		return
	}

	var result = struct {
		Errcode     int    `json:"errcode"`
		AccessToken string `json:"access_token"`
		Errmsg      string `json:"errmsg"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return
	}

	if result.AccessToken == "" {
		err = fmt.Errorf("%s", result.Errmsg)
		return
	}

	return result.AccessToken, result.ExpiresIn, nil
}
