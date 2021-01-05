package controllers

import (

	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/music-player/server/service"

)

func QueryAllData(c *gin.Context) {

	c.IndentedJSON(200, service.GetFromDB())

}

func QueryDataByName(c *gin.Context) {

	name := c.Query("name")

	fmt.Println(name)

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