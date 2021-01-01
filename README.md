# DingAPI Server Golang SDK
    封装操作钉钉api，核心功能已实现。后面会在闲余时间慢慢的加上各种api的实现，意在开箱即用。请给个star
# 安装

    go get -u github.com/Berry961103/ding
# 快速开始

#### 钉钉群机器人，目前只实现了文本格式的发送       

    import (
	
        "github.com/Berry961103/ding"
        "github.com/Berry961103/ding/internal/robot"
    )

    func main() {
        config := ding.AppConfig{
            RobotToken: "1111",
        }
        appCtx := ding.NewApp(config)

        err := robot.DingRobotText(appCtx, "berry", []string{"12345678"})
        if err != nil {
           panic(err)
        }
    }
## API的实现
#### 目前实现的功能比较少，后面有空会慢慢添加，目前实现功能在[internal](https://github.com/Berry961103/ding/tree/master/internal)中。请持续关注
    import (
        "fmt"

        "github.com/Berry961103/ding"
	    "github.com/Berry961103/ding/entity"
        "github.com/Berry961103/ding/internal/calendar"
    )

    func main(){
        config := ding.AppConfig{
            AgentID:   12345678,
            AppKey:    "AppKey",
            AppSecret: "AppSecre",
        }
        appCtx := ding.NewApp(config)

        calendarCreate := entity.CalendarCreateRequest{}
        resp, err := calendar.Create(appCtx, calendarCreate)
        if err != nil {
            fmt.Println(err)
            return
        }

       fmt.Printf("%+v\n",resp)    
    }

# 自定义
### 默认为您实现了内存缓存，默认从ding api server获取access_token，默认的http客户端

## 缓存
#### 按照[cache](https://github.com/Berry961103/ding/blob/master/cache/cache.go)接口实现功能就可以自定义缓存token

    cache:=&RedisCacheImpl{}
    appCtx.SetAccessTokenCacheDriver(cache)

### 刷新缓存
    按照 type GetAccessTokenFunc func() (accessToken string, err error) 类型传gettoken方法即可

    func GetAccessToken()(string,error){
        // 你的代码
        return accessToken,nil
    }

    appCtx.SetGetAccessTokenHandler(GetAccessToken)

### 日志

#### 日志目前只输出debug级别的日志，只需要实现 [logger](https://github.com/Berry961103/ding/blob/master/logger.go) 接口即可

    支持zap

        logger, _ := zap.NewProduction()
        sugar := logger.Sugar()
        appCtx.SetLogger(sugar)

    支持logrus

        loggger = logrus.New() 
        appCtx.SetLogger(sugar)   

### 钉钉请求API的HTTP客户端配置

    httpClient:=&http.Client{
        Timeout: 10*time.Second,
        // ... 其他自定义配置
    }

    appCtx.SetHTTPClient(httpClient)


