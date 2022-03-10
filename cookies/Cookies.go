package cookies

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("cookie", hanleCookie)
}

func hanleCookie(c *gin.Context) {
	cookie, err := c.Cookie("key_cookie")
	if err != nil {
		// 给客户端设置cookie
		cookie = "NotSet"
		// maxAge int, 单位为秒
		// path,cookie所在目录
		// domain string,域名
		// secure 是否智能通过https访问
		// httpOnly bool  是否允许别人通过js获取自己的cookie
		c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
	}
	fmt.Printf("cookie的值：%s", cookie)
}
