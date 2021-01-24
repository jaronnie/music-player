package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nj-jay/music-player/server/controllers"
	"github.com/nj-jay/music-player/server/global"
	"github.com/nj-jay/music-player/server/ws"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/api/v1/list", controllers.QueryAllData)
	e.GET("/api/v1/search", controllers.QueryDataByName)
	e.GET("/api/v1/song/url", controllers.PlayMusic)
	e.GET("/api/v1/click", controllers.ClickCount)
	e.GET("/ws", func(c *gin.Context) {
		ws.ServeWs(global.HUB, c)
	})
}