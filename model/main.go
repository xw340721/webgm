package model

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/config"
)

//NewConn 创建sql通道
func NewConn() *sql.DB {
	var err error
	conn, err := sql.Open(config.Get("env", "driver"), config.Get("env", "address"))
	if err != nil {
		logrus.Error("[conn] 创建mysql连接", err.Error())
		//os.Exit(0)
	}
	return conn
}

//ReturnToJson 将查询的数据返回为json格式
func ReturnToJson(rows *sql.Rows) ([]byte, int, error) {

	//分析columns
	columns, err := rows.Columns()

	if err != nil {
		mLogrus.Error("[test] 解析columns失败")
	}

	//todo num 不应该在这里面取值
	var num = 0
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
