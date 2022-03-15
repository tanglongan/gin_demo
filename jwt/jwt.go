package jwt

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("疫情悄悄的过去吧")

type UserInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Routers(e *gin.Engine) {
	e.POST("/auth", authHandler)
	e.GET("/home", AuthMiddleware(), homeHandler)
}

func authHandler(c *gin.Context) {
	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 2001, "msg": "无效的参数"})
		return
	}

	if user.UserName == "tanglongan" && user.Password == "123456" {
		tokenString, err := GenToken(user.UserName)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

// AuthMiddleware 基于JWT的认证中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按照空格分开
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数解析它
		token, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的token",
			})
			c.Abort()
			return
		}

		// 上下文中设置当前用户信息
		c.Set("username", token.Username)
		//后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Next()
	}
}

func homeHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "tanglongan",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的Secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
