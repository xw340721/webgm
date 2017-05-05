package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gogap/logrus"
)

var configData map[string]map[string]string

var mLogrus = logrus.WithField("package", "config")

//Load 加载config文件
func Load(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &configData)
	if err != nil {
		return err
	}
	return nil
}

//Get 获取config文件
func Get(table, key string) string {
	t, ok := configData[table]
	if !ok {
		mLogrus.Warnf("[config] 无法读取%v.%v", table, key)
		return ""
	}
	val, ok := t[key]
	if !ok {
		mLogrus.Warnf("[config] 无法读取%v.%v", table, key)
		return ""
	}
	return val
}

//Set 设置config文件
func Set(key1, key2, val string) {
	if kmap, ok := configData[key1]; ok {
		kmap[key2] = val
	}
}
