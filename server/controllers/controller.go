package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nj-jay/music-player/server/model"
	"github.com/nj-jay/music-player/server/service"
)

func QueryAllData(c *gin.Context) {
	c.IndentedJSON(200, service.GetFromDB())
}

func QueryDataByName(c *gin.Context) {
	name := c.Query("name")
	songs := service.QueryDataByName(name)

	c.IndentedJSON(200, gin.H{
		"code": 200,
		"data": songs,
	})
}

func PlayMusic(c *gin.Context) {
	id := c.Query("id")
	url := service.PlayMusic(id)

	c.IndentedJSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"url": url,
		},
	})
}

func ClickCount(c *gin.Context) {
	count := service.AddClickCount()

	c.IndentedJSON(200, gin.H{
		"code": 200,
		"data": count,
	})
}

func TrueLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	tokenString, err := service.TrueLogin(username, password)

	if err != nil {
		c.IndentedJSON(200, model.NewLoginErr())
	} else {
		c.IndentedJSON(http.StatusOK, model.NewLoginSuccess(tokenString))
	}
}

func PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	status, tokenString := service.PostLogin(username, password)

	if status == 200 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "success add user",
			"data":     tokenString,
			"username": username,
		})
	} else {
		c.IndentedJSON(http.StatusOK, model.NewUserHasExistedErr(username))
	}
}

func GetCacheMessage(c *gin.Context) {
	msg := service.GetCacheMessage()
	c.IndentedJSON(200, gin.H{
		"data": msg,
	})
}
