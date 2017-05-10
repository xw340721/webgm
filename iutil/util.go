package iutil

import (
	"encoding/base64"
	"errors"
	"reflect"
	"strings"

	"github.com/gogap/logrus"
)

var StringUpperIndexError = errors.New("传入字段小于1个字符")

var mLogrus = logrus.WithFields(logrus.Fields{"package": "iutls"})

//DecodeBase64 将string进行解码
func DecodeBase(s string) string {
	str, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		mLogrus.Error("[util] 解析base64错误", err.Error())
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

func DeepField(iface interface{}) []reflect.Value {
	fields := make([]reflect.Value, 0)
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			fields = append(fields, DeepField(v.Interface())...)
		default:
			fields = append(fields, v)
		}
	}
	return fields
}
