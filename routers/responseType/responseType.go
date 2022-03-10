package responseType

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func Routers(e *gin.Engine) {
	e.GET("/aJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})

	e.POST("/aStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})

	e.GET("/aXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})

	e.GET("/aYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan"})
	})

	e.GET("aProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
}
