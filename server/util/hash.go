package util

import "golang.org/x/crypto/bcrypt"

//hash加密
func HashEncryption(password string) string {
	pwdByte := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pwdByte, bcrypt.MinCost)
	return string(hash)
}

//hash解密
func HashDecrypt(hashPwd, password string) bool {
	hashPwdByte := []byte(hashPwd)
	pwdByte := []byte(password)
	b := bcrypt.CompareHashAndPassword(hashPwdByte, pwdByte)
	if b == nil {
		return true
	} else {
		return false
	}
}