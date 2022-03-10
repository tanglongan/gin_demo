package asyncRequest

// goroutine机制可以方便地实现异步处理
// 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Routers(e *gin.Engine) {
	// 异步请求
	e.GET("async", func(c *gin.Context) {
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	// 同步请求
	e.GET("sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

}
