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

//User信息
type Login struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string	`json:"password"`
}