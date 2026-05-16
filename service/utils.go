package service

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"regexp"
)

// 验证密码格式
func IsValidPassword(password string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9!@#\$%\^&*_\-]+$`)
	return regex.MatchString(password)
}

// 验证邮箱格式
func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

// MD5函数
func MakeMd5(str string) (md5Str string) {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
