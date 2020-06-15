package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"strings"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

//Md5Check 验证一个文本MD5的值是否和指定值相同
func Md5Check(content, encrypted string) bool {
	return strings.EqualFold(EncodeMD5(content), encrypted)
}

//Md5SumFile 得到文件的md5值
func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return
}
