package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jaronnie/music-player/server/global"
	"github.com/jaronnie/music-player/server/initalize"
	"github.com/jaronnie/music-player/server/middlewares"
	"github.com/jaronnie/music-player/server/routers"
	"github.com/jaronnie/music-player/server/ws"
)

func main() {
	global.GMD_DB = initalize.GormMysql()
	initalize.MysqlTables(global.GMD_DB)
	global.GMD_RD = initalize.ConnectRedis()
	ws.HUB = ws.NewHub()
	r := gin.Default()
	r.Use(middlewares.Cors())
	go ws.HUB.Run()
	routers.LoadRouter(r)
	err := r.Run(":8083")
	if err != nil {
		log.Fatal("run err")
	}
}
