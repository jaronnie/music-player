package service

import (
	"github.com/nj-jay/music-player/server/global"
	"github.com/nj-jay/music-player/server/model"
)


func GetFromDB() []model.Song {
	var allSong []model.Song
	global.GMD_DB.Find(&allSong)
	return allSong
}

func QueryDataByName(name string) []model.Song {
	var songs []model.Song
	global.GMD_DB.Where("name REGEXP ?", name).Find(&songs)
	return songs
}

func PlayMusic(id string) string {
	var song model.Song
	global.GMD_DB.Where("id=?", id).Find(&song)
	return song.Url
}