package paramValidate

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func Routers(e *gin.Engine) {
	e.GET("/validatePerson", validatePersonHandle)
}

type Person struct {
	//年龄大于10，必填
	Age int `form:"age" json:"age" xml:"age" binding:"required,gt>10"`
	//通过自定义函数校验
	Name string `form:"name" json:"name" xml:"name" binding:"NotNullAndAdmin"`
	//日期格式
	Birthday time.Time `form:"birthday" json:"birthday" xml:"birthday" binding:"BirthdayEarlyThanToday" time_format:"2006-01-02"`
}

func validatePersonHandle(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusOK, "person bind err:%v", err.Error())
	}
}

// ValidateNameNotNullAndAdmin 校验名称不为空并且不等于tanglongan
func ValidateNameNotNullAndAdmin(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		// 字段不能为空，并且不等于  tanglongan
		return value != "" && !("tanglongan" == value)
	}

	return true
}

// ValidateBirthdayEarlyThanToday 校验日期必须在当前日期之前
func ValidateBirthdayEarlyThanToday(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() > date.Unix() {
			return false
		}
	}
	return true
}
