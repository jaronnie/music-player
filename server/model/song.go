package model

type Song struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Url string `json:"url"`
}

func NewSong(name, url string) *Song {
	return &Song{
		Name: name,
		Url: url,
	}
}