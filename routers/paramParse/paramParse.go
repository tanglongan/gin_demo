package paramParse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 定义接收数据的结构体
type Login struct {
	UserName string `form:"username" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"pass" xml:"pass" binding:"required"`
}

func Routers(e *gin.Engine) {
	e.POST("/loginJSON", loginJSON)
	e.POST("/loginForm", loginForm)
	e.POST("/loginURL", loginURL)
}

func loginJSON(c *gin.Context) {
	var json Login
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate(json, c)
}

func loginForm(c *gin.Context) {
	var form Login
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate(form, c)
}

func loginURL(c *gin.Context) {
	var login Login
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate(login, c)
}

func validate(data Login, c *gin.Context) {
	if data.UserName != "root" || data.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
