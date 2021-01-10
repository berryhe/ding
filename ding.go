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
	"net/http"

	"github.com/Berry961103/ding/cache"
)

// GetAccessTokenFunc 获取 access_token 方法接口
type GetAccessTokenFunc func(appKey, appSecret string) (accessToken string, err error)

// DingCtx 实例
type DCtx struct {
	Config      Config
	accessToken AccessToken
	httpClient  *http.Client
	logger      Logger
}

//AccessToken 管理器 处理缓存和刷新逻辑
type AccessToken struct {
	cache                 cache.Cache
	getAccessTokenHandler GetAccessTokenFunc
}

// Config 主配置
type Config struct {
	CorpID         string
	AgentID        int
	AppKey         string
	AppSecret      string
	Token          string
	RobotToken     string
	EncodingAESKey string
}

// NewDCtx 初始化项
func NewDCtx(config Config) (app *DCtx) {
	app = newDCtx(config)
	app.accessToken.getAccessTokenHandler = app.getAccessToken
	return
}

// 创建实例
func newDCtx(config Config) (app *DCtx) {
	return &DCtx{
		Config:     config,
		httpClient: &http.Client{},
		accessToken: AccessToken{
			cache: cache.NewDefaultCache(),
		},
	}
}

// SetHTTPClient 自定义设置钉钉Api请求HTTP客户端
func (dCtx *DCtx) SetHTTPClient(client *http.Client) {
	dCtx.httpClient = client
}

// SetAccessTokenCacheDriver 设置 AccessToken 缓存器 默认为内存缓存
// 驱动接口类型 为 cache.Cache
func (dCtx *DCtx) SetAccessTokenCacheDriver(driver cache.Cache) {
	dCtx.accessToken.cache = driver
}

// SetGetAccessTokenHandler 设置 AccessToken 获取方法。默认从sync.Map获取（过期从钉钉接口刷新）
// 如果有多实例服务，可以设置为 Redis 或 rpc 等中间件中获取就可以避免 AccessToken 刷新冲突
func (dCtx *DCtx) SetGetAccessTokenHandler(f GetAccessTokenFunc) {
	dCtx.accessToken.getAccessTokenHandler = f
}

// SetLogger 日志记录
// 默认不开启日志，开了也只是输出debug级别的日志
func (dCtx *DCtx) SetLogger(logger Logger) {
	dCtx.logger = logger
}
