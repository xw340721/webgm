package model

import (
	"database/sql"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/config"

	_ "github.com/go-sql-driver/mysql"
)

//NewConn 创建sql通道
func NewConn() *sql.DB {
	var err error
	conn, err := sql.Open(config.Get("env", "driver"), config.Get("env", "address"))
	if err != nil {
		logrus.Error("[启动] 创建mysql连接", err.Error())
		//os.Exit(0)
	}
	return conn
}
