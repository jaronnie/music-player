package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/music-player/server/model"
)

type MyClaims struct {
	model.Login
	jwt.StandardClaims
}

const (
	TokenExpireDuration = time.Hour * 10000
)

var (
	MySecret = []byte("gocloudcoder.com")
)

func GenToken(username, password string) (string, error) {
	mc := MyClaims{
		Login: model.Login{Username: username, Password: password},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "niejian",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求头中带token
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 2003,
				"meg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		fmt.Println(parts)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Set("password", mc.Password)
		c.Next()
	}
}

func JWTAuthPathMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求路径中带token
		token := c.Query("token")

		if token == "" {
			c.IndentedJSON(http.StatusOK, gin.H{
				"code": 2003,
				"meg":  "token为空",
			})
			c.Abort()
			return
		}

		mc, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		fmt.Println(mc.Username)
		c.Set("username", mc.Username)
		c.Set("password", mc.Password)
		c.Next()
	}
}
