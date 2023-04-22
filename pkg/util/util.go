package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"strings"
)

// CalcMD5 计算字符串的 MD5 值
func CalcMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return fmt.Sprintf("%x", cipherStr)
}

// CalSha1 计算字符串的 Sha1 值
func CalSha1(str string) string {
	data := []byte(str)
	sha1 := sha1.Sum(data)
	return fmt.Sprintf("%x", sha1)
}

// Join 将字符串切片组装成逗号分隔符组成的字符串
func Join(items []string) {
	strings.Join(items, ",")
}

// Split 将用逗号分隔符组成的字符串切分成字符串切片
func Split(str string) []string {
	return strings.Split(str, ",")
}
