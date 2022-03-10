package redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e *gin.Engine) {
	e.GET("/goBaidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
}
