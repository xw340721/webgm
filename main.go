package main

import (
	"os"

	"github.com/go-martini/martini"
	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/config"
	"github.com/xw340721/webgm/handle"
	"github.com/xw340721/webgm/middleware"
	"github.com/xw340721/webgm/model"
)

func init() {
	err := config.Load("./env.json")
	if err != nil {
		logrus.Error("[启动] 读取配置文件", err.Error())
		os.Exit(0)
	}

	//连接数据库
	model.NewConn()
}

func main() {
	m := martini.Classic()
	m.Post("/", handle.Main)
	m.Get("/", handle.Main)
	m.Handlers(middleware.API)

	m.RunOnAddr(":" + config.Get("env", "port"))
}
