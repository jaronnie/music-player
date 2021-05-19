package service

import (
	"context"
	"io"
	"log"

	"github.com/nj-jay/music-player/server/global"
)

//count.go
var ctx = context.Background()

func AddClickCount() string {
	_, err := global.GMD_RD.Incr(ctx, "music_player_count").Result()
	if err == io.EOF {
		AddClickCount()
	}
	return get("music_player_count")
}

func get(key string) string {
	result, err := global.GMD_RD.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func CacheMessage(msg string) {
	global.GMD_RD.RPush(ctx, "music_player_cacheMessage", msg)
}

func GetCacheMessage() []string {
	result := global.GMD_RD.LRange(ctx, "music_player_cacheMessage", -5, -1)
	msg, _ := result.Result()
	return msg
}
