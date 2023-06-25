package service

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

//校验手机号
func CheckPhone(phone string) bool {
	rule := "^[1][3,4,5,7,8][0-9]{9}"
	result, _ := regexp.MatchString(rule, phone)
	return result
}

//校验邮政编码
func CheckEmail(email string) bool {
	rule := "^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}"
	result, _ := regexp.MatchString(rule, email)
	return result
}

//密码加密
func EncryptPassword(password string) string {
	data := []byte(password)
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}
