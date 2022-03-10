package globalMiddleWare

// 所有请求都会经过此中间件
// 中间件推荐：https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E4%B8%AD%E9%97%B4%E4%BB%B6/%E4%B8%AD%E9%97%B4%E4%BB%B6%E6%8E%A8%E8%8D%90.html
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Routers(e *gin.Engine) {
	e.GET("/*", filter)
}

// GlobalMiddleWare 定义一个中间件
func GlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 执行目标函数之前
		begin := time.Now()
		fmt.Println("中间件开始执行")
		c.Set("aVariable", "设置的一个变量值")

		// 执行目标函数
		c.Next()

		// 执行目标函数之后
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		duration := time.Since(begin)
		fmt.Println("程序用时: ", duration)
	}
}

func filter(c *gin.Context) {
	req, _ := c.Get("aVariable")
	fmt.Println("aVariable: ", req)
	c.JSON(http.StatusOK, gin.H{"aVariable": req})
}
