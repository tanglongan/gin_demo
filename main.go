package main

import (
	"gin_demo/cookies"
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
	// 创建路由
	r := gin.Default()

	//注册路由
	Include(paramRouters.Routers,
		paramParse.Routers,
		responseType.Routers,
		redirect.Routers,
		asyncRequest.Routers,
		cookies.Routers,
		paramValidate.Routers)

	r.Use(globalMiddleWare.GlobalMiddleWare())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("NameNotNullAndAdmin", paramValidate.ValidateNameNotNullAndAdmin)
		_ = v.RegisterValidation("BirthdayEarlyThanToday", paramValidate.ValidateBirthdayEarlyThanToday)
	}

	// 启动并监听服务
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Listen At %s, err:%v", port, err)
	}

}
