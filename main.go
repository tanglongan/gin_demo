package main

import (
	"fmt"
	"gin_demo/cookies"
	"gin_demo/jwt"
	"gin_demo/middleWare/globalMiddleWare"
	"gin_demo/routers/asyncRequest"
	"gin_demo/routers/paramParse"
	"gin_demo/routers/paramRouters"
	"gin_demo/routers/paramValidate"
	"gin_demo/routers/redirect"
	"gin_demo/routers/responseType"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

const (
	port = ":8080"
)

type Option func(c *gin.Engine)

var options []Option

// Init 初始化
func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}

func Include(opts ...Option) {
	options = append(options, opts...)
}

func main() {
	// 收集路由列表
	Include(paramRouters.Routers,
		paramParse.Routers,
		responseType.Routers,
		redirect.Routers,
		asyncRequest.Routers,
		cookies.Routers,
		paramValidate.Routers,
		jwt.Routers)

	// 注册路由
	r := Init()

	// 注册中间件
	r.Use(globalMiddleWare.GlobalMiddleWare())

	// 注册中间件：参数校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("NameNotNullAndAdmin", paramValidate.ValidateNameNotNullAndAdmin)
		_ = v.RegisterValidation("BirthdayEarlyThanToday", paramValidate.ValidateBirthdayEarlyThanToday)
	}

	// 注册中间件：注册日志的自定义格式化器
	r.Use(gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			p.ClientIP,
			p.TimeStamp.Format(time.RFC1123),
			p.Method,
			p.Path,
			p.Request.Proto,
			p.StatusCode,
			p.Latency,
			p.Request.UserAgent(),
			p.ErrorMessage,
		)
	}))

	// 注册中间件：主进程的异常恢复
	r.Use(gin.Recovery())
	// 启动并监听服务
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Listen At %s, err:%v", port, err)
	}

}
