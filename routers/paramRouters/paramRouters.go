package paramRouters

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Routers(e *gin.Engine) {
	e.GET("/hello", hello)
	e.GET("/paramRouters/:name/*action", queryUrl)
	e.GET("/paramRouters", queryParam)
	e.POST("/form", form)
	e.POST("/upload", upload)
	e.POST("/multiUpload", multiUpload)

	// routes group是为了管理一些相同的URL
	v1 := e.Group("/v1")
	{
		v1.GET("/paramParse", login)
		v1.GET("submit", submit)
	}
	v2 := e.Group("/v2")
	{
		v2.GET("/paramParse", login)
		v2.GET("submit", submit)
	}
}

// hello 没有参数
func hello(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

// queryUrl 路径参数
func queryUrl(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
}

// queryParam 查询参数
func queryParam(c *gin.Context) {
	name := c.DefaultQuery("name", "我是参数默认值") // 使用DefaultQuery()获取URL参数值，如果没传就返回默认值
	age := c.Query("age")                     // 使用Query()获取URL参数值，如果没传就返回空串
	c.String(http.StatusOK, fmt.Sprintf("name: %s\nage: %s", name, age))
}

// form FORM参数
func form(c *gin.Context) {
	//PostForm()方法默认解析的是x-www-form-urlencoded或from-data格式的参数
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}

// upload 单文件上传
func upload(c *gin.Context) {
	// FormFile()获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传失败")
	}
	_ = c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
}

// multiUpload 多文件上传
func multiUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
	}
	//文件域字段
	files := form.File["files"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("upload OK! %d files", len(files)))
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
