package routers

import (

	"github.com/gin-gonic/gin"

	"github.com/nj-jay/music-player/server/controllers"


)

func LoadRouter(e *gin.Engine) {

	e.GET("/api/v1/list", controllers.QueryAllData)

	e.GET("/api/v1/search", controllers.QueryDataByName)

	e.GET("/api/v1/song/url", controllers.PlayMusic)

	e.GET("/api/v1/click", controllers.ClickCount)

}