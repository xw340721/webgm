package main

import (
	"flag"
	"os"

	"github.com/go-martini/martini"
	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/config"
	"github.com/xw340721/webgm/handle"
	"github.com/xw340721/webgm/middleware"
)

var configPath string

var mLogrus = logrus.WithField("package", "main")

func init() {
	flag.StringVar(&configPath, "c", "./env.json", "--set application config path")
}

func main() {
	flag.Parse()
	//config文件读取
	err := config.Load(configPath)
	if err != nil {
		mLogrus.Error("[main] 读取配置文件", err.Error())
		os.Exit(0)
	}

	m := martini.Classic()
	m.Post("/", handle.Main)
	m.Get("/", handle.Main)
	m.Handlers(middleware.API, middleware.LOG)

	m.RunOnAddr(":" + config.Get("env", "port"))
}
