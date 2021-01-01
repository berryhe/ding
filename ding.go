package ding

import (
	"net/http"

	"github.com/Berry961103/ding/cache"
)

// GetAccessTokenFunc 获取 access_token 方法接口
type GetAccessTokenFunc func(appKey, appSecret string) (accessToken string, err error)

// App 实例
type App struct {
	Config      AppConfig
	accessToken AccessToken
	httpClient  *http.Client
	logger      Logger
}

//AccessToken 管理器 处理缓存和刷新逻辑
type AccessToken struct {
	cache                 cache.Cache
	getAccessTokenHandler GetAccessTokenFunc
}

// AppConfig 主配置
type AppConfig struct {
	CorpID         string
	AgentID        int
	AppKey         string
	AppSecret      string
	Token          string
	RobotToken     string
	EncodingAESKey string
}

// NewApp 初始化项
func NewApp(config AppConfig) (app *App) {
	app = newApp(config)
	app.accessToken.getAccessTokenHandler = app.getAccessToken
	return
}

// 创建实例
func newApp(config AppConfig) (app *App) {
	return &App{
		Config:     config,
		httpClient: &http.Client{},
		accessToken: AccessToken{
			cache: cache.NewDefaultCache(),
		},
	}
}

// SetHTTPClient 自定义设置钉钉Api请求HTTP客户端
func (app *App) SetHTTPClient(client *http.Client) {
	app.httpClient = client
}

// SetAccessTokenCacheDriver 设置 AccessToken 缓存器 默认为内存缓存
// 驱动接口类型 为 cache.Cache
func (app *App) SetAccessTokenCacheDriver(driver cache.Cache) {
	app.accessToken.cache = driver
}

// SetGetAccessTokenHandler 设置 AccessToken 获取方法。默认从sync.Map获取（过期从钉钉接口刷新）
// 如果有多实例服务，可以设置为 Redis 或 rpc 等中间件中获取就可以避免 AccessToken 刷新冲突
func (app *App) SetGetAccessTokenHandler(f GetAccessTokenFunc) {
	app.accessToken.getAccessTokenHandler = f
}

// SetLogger 日志记录
// 默认不开启日志，开了也只是输出debug级别的日志
func (app *App) SetLogger(logger Logger) {
	app.logger = logger
}
