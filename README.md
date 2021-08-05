# DingAPI Server Golang SDK
    封装操作钉钉api，核心功能已实现。后面会在闲余时间慢慢的加上各种api的实现，意在开箱即用。请给个star
# 安装

    go get -u github.com/berryhe/ding
# 快速开始

#### 钉钉群机器人，目前只实现了文本格式的发送       

    import (
	
        "github.com/berryhe/ding"
        "github.com/berryhe/ding/app/robot"
    )

    func main() {
        config := ding.Config{
            RobotToken: "1111",
        }
        dingCtx := ding.DCtx(config)

        err := robot.DingRobotText(dingCtx, "berry", []string{"12345678"})
        if err != nil {
           panic(err)
        }
    }
## API的实现
#### 目前实现的功能比较少，后面有空会慢慢添加，目前实现功能在[app](https://github.com/berryhe/ding/tree/master/app)中。请持续关注
    import (
        "fmt"

        "github.com/berryhe/ding"
	    "github.com/berryhe/ding/entity"
        "github.com/berryhe/ding/app/calendar"
    )

    func main(){
        config := ding.Config{
            AgentID:   12345678,
            AppKey:    "AppKey",
            AppSecret: "AppSecre",
        }
        dingCtx := ding.DCtx(config)

        calendarCreate := entity.CalendarCreateRequest{}
        resp, err := calendar.Create(dingCtx, calendarCreate)
        if err != nil {
            fmt.Println(err)
            return
        }

       fmt.Printf("%+v\n",resp)    
    }

# 自定义
### 默认为您实现了内存缓存，默认从ding api server获取access_token，默认的http客户端

## 缓存
#### 按照[cache](https://github.com/berryhe/ding/blob/master/cache/cache.go)接口实现功能就可以自定义缓存token

    cache:=&RedisCacheImpl{}
    dingCtx.SetAccessTokenCacheDriver(cache)

## 获取acces_token
### 每次api的调用都会先去调一次获取access_token的方法，请在获取access_token的方法逻辑里面存入缓存和获取缓存的逻辑，如想自实现请参考源码中默认[getAccessToken](https://github.com/berryhe/ding/blob/master/apps.go)方法的实现

    // appKey和appSecretKey 会自动帮你从appCoonfig中设置的传入
    func GetAccessToken(appKey, appSecretKey string)(string,error){
        // 你的代码
        return accessToken,nil
    }

    dingCtx.SetGetAccessTokenHandler(GetAccessToken)

### 日志

#### 日志目前只输出debug级别的日志，只需要实现 [logger](https://github.com/berryhe/ding/blob/master/logger.go) 接口即可

    支持zap

        logger, _ := zap.NewProduction()
        sugar := logger.Sugar()
        dingCtx.SetLogger(sugar)

    支持logrus

        loggger = logrus.New() 
        dingCtx.SetLogger(sugar)   

### 钉钉请求API的HTTP客户端配置

    httpClient:=&http.Client{
        Timeout: 10*time.Second,
        // ... 其他自定义配置
    }

    dingCtx.SetHTTPClient(httpClient)


