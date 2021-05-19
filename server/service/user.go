package service

import (
	"errors"

	"github.com/nj-jay/music-player/server/global"
	"github.com/nj-jay/music-player/server/middlewares"
	"github.com/nj-jay/music-player/server/model"
	"github.com/nj-jay/music-player/server/util"
)

type Login model.Login

//返回所有用户的登录信息
func QueryLogin() []model.Login {
	var logins []model.Login
	result := global.GMD_DB.Find(&logins)
	if result.Error != nil {
		return nil
	}
	return logins
}

func TrueLogin(username, password string) (string, error) {
	var login model.Login
	err := global.GMD_DB.Where("username = ?", username).Find(&login).Error
	if err != nil {
		return "", errors.New("用户不存在")
	}

	if username == login.Username && util.HashDecrypt(login.Password, password) {
		tokenString, _ := middlewares.GenToken(login.Username, login.Password)
		return tokenString, nil
	} else {
		return "", errors.New("密码不正确")
	}
}

//添加用户
func PostLogin(username, password string) (int, string) {

	//password加密
	hashPwd := util.HashEncryption(password)
	login := &model.Login{
		Username: username,
		Password: hashPwd,
	}

	result := global.GMD_DB.Create(&login)
	tokenString, _ := middlewares.GenToken(login.Username, login.Password)
	if result.Error != nil {
		return 401, ""
	} else {
		return 200, tokenString
	}
}
