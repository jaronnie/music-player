package service

import (
	"context"
	"github.com/nj-jay/music-player/server/global"
	"io"
	"log"
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

func get(key string) string{
	result, err := global.GMD_RD.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return result
}