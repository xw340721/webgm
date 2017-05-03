package model

import (
	"database/sql"
	"os"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/config"

	_ "github.com/go-sql-driver/mysql"
)

var conn *sql.DB

//NewConn 创建sql通道
func NewConn() *sql.DB {
	var err error
	conn, err = sql.Open(config.Get("env", "driver"), config.Get("env", "address"))
	defer conn.Close()
	if err != nil {
		logrus.Error("[启动] 创建mysql连接", err.Error())
		os.Exit(0)
	}
	return conn
}

//GetConn 获取已知通道
func GetConn() *sql.DB {
	return conn
}
