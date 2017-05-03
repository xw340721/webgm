package iutil

import (
	"encoding/base64"

	"strings"

	"errors"

	"github.com/gogap/logrus"
)

var StringUpperIndexError = errors.New("[函数库] 传入字段小于1个字符")

//DecodeBase64 将string进行解码
func DecodeBase(s string) string {
	str, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		logrus.Error("[util] 解析base64错误", err.Error())
	}
	return string(str)
}

func StringUpperIndex(s string) (string, error) {
	if len(s) > 0 {
		strings.ToUpper(s[:1])
		return s, nil
	}
	return "", StringUpperIndexError
}
