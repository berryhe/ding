package ding

import (
	"log"
	"net/http"
	"os"

	"github.com/Berry961103/cache"
)

// GetAccessTokenFunc 获取 access_token 方法接口
type GetAccessTokenFunc func() (accessToken string, err error)

/*
App 实例
*/
type App struct {
	Config      AppConfig
	accessToken AccessToken
	Client      Client
	httpClient  *http.Client
	Logger      *log.Logger
}

/*
AccessToken 管理器 处理缓存 和 刷新 逻辑
*/
type AccessToken struct {
	Cache                 cache.Cache
	GetAccessTokenHandler GetAccessTokenFunc
}

// AppConfig 主配置
type AppConfig struct {
	CorpID         string
	AgentID        string
	AppKey         string
	AppSecret      string
	Token          string
	RebotToken     string
	EncodingAESKey string
}

// 创建实例
func newApp(config AppConfig) (app *App) {
	instance := App{
		Config: config,
		accessToken: AccessToken{
			Cache: cache.NewDefaultCache(),
		},
	}

	instance.Client = Client{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	return &instance
}

// SetAccessTokenCacheDriver 设置 AccessToken 缓存器 默认为内存缓存
// 驱动接口类型 为 cache.Cache
func (app *App) SetAccessTokenCacheDriver(driver cache.Cache) {
	app.accessToken.Cache = driver
}

// SetGetAccessTokenHandler 设置 AccessToken 获取方法。默认 从本地缓存获取（过期从钉钉接口刷新）
// 如果有多实例服务，可以设置为 Redis 或 RPC 等中控服务器 获取 就可以避免 AccessToken 刷新冲突
func (app *App) SetGetAccessTokenHandler(f GetAccessTokenFunc) {
	app.accessToken.GetAccessTokenHandler = f
}

/*
SetLogger 日志记录 默认输出到 os.Stdout

可以新建 logger 输出到指定文件

如果不想开启日志，可以 SetLogger(nil)
*/
func (app *App) SetLogger(logger *log.Logger) {
	app.Logger = logger
}
