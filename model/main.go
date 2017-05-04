package model

import (
	"database/sql"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/config"

	"encoding/json"

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

//EntryToJson 将查询的数据返回为json格式
func ReturnToJson(columns []string, rows *sql.Rows) ([]byte, int, error) {
	//todo num 不应该在这里面取值
	var num int = 0
	count := len(columns)

	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuesPtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuesPtrs[i] = &values[i]
		}
		rows.Scan(valuesPtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v

		}
		tableData = append(tableData, entry)
		num++
	}
	data, _ := json.Marshal(tableData)

	return data, num, nil
}
