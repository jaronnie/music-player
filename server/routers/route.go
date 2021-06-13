package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/music-player/server/controllers"
	"github.com/jaronnie/music-player/server/middlewares"
	"github.com/jaronnie/music-player/server/ws"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/api/v1/list", middlewares.JWTAuthMiddleware(), controllers.QueryAllData)
	e.GET("/api/v1/search", controllers.QueryDataByName)
	e.GET("/api/v1/song/url", controllers.PlayMusic)
	e.GET("/api/v1/click", controllers.ClickCount)
	e.POST("/api/v1/trueLogin", controllers.TrueLogin)
	e.POST("/api/v1/adduser", controllers.PostLogin)
	e.GET("/ws", func(c *gin.Context) {
		ws.ServeWs(ws.HUB, c)
	})
	e.GET("/api/v1/cache", controllers.GetCacheMessage)
}
