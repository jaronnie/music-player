package service

import (

	"fmt"

	"github.com/nj-jay/music-player/server/global"

	"github.com/nj-jay/music-player/server/model"

)

//func GetAllRemoteFile() {
//
//	accessKey := "0BXxOytiWWOySUiEbg-t8_08c0tvPpwTwDA6Ivfn"
//
//	secretKey := "ewA76OHGBEz43vQg-gCqOHd_5paEurmSEzFdx-dz"
//
//	mac := qbox.NewMac(accessKey, secretKey)
//
//	//var allSong []model.Song
//
//	cfg := storage.Config{
//
//		UseHTTPS: true,
//
//	}
//
//	bucketManager := storage.NewBucketManager(mac, &cfg)
//
//	bucket := "nj-jay"
//
//	limit := 1000
//
//	prefix := ""
//
//	delimiter := ""
//
//	//初始列举marker为空
//	marker := ""
//
//	for {
//
//		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
//
//		if err != nil {
//
//			fmt.Println("list error,", err)
//
//			break
//
//		}
//
//		for _, entry := range entries {
//
//			if strings.HasSuffix(entry.Key, "mp3") {
//
//				name := entry.Key[0:strings.LastIndexAny(entry.Key, ".")]
//
//				song := model.NewSong(name, "https://picture.nj-jay.com/" + entry.Key)
//
//				global.GMD_DB.Create(&song)
//
//			}
//
//		}
//
//		if hasNext {
//
//			marker = nextMarker
//
//		} else {
//
//			break
//
//		}
//	}
//
//}

func GetFromDB() []model.Song {

	var allSong []model.Song

	global.GMD_DB.Find(&allSong)

	return allSong

}

func QueryDataByName(name string) []model.Song {

	fmt.Println(name)

	var songs []model.Song

	global.GMD_DB.Where("name REGEXP ?", name).Find(&songs)

	return songs

}

func PlayMusic(id string) string {

	var song model.Song

	global.GMD_DB.Where("id=?", id).Find(&song)

	return song.Url

}