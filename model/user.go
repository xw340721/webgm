package model

import (
	"encoding/json"

	"github.com/gogap/logrus"
)

type Test struct {
	ID     int  `json:"id"`
	GameId uint `json:"game_id"`
}
type TestEntry []Test

//GetUser 返回查询的玩家数量
func (t *Test) Test(serverID int) (interface{}, error) {
	conn := NewConn()

	defer conn.Close()
	stmt, err := conn.Prepare(`SELECT id,game_id FROM server WHERE server_id = ?`)
	if err != nil {
		logrus.Error("[mysql] 准备解析失败", err.Error())
	}
	rows, err := stmt.Query(serverID)
	defer rows.Close()

	if err != nil {
		logrus.Error("[mysql] 查询失败", err.Error())
	}

	columns, err := rows.Columns()

	if err != nil {
		logrus.Error("[mysql] 解析columns失败")
	}

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
	}

	data, _ := json.Marshal(tableData)

	testEntries := TestEntry{}

	json.Unmarshal(data, &testEntries)

	err = rows.Err()
	if err != nil {
		logrus.Error("[mysql] 查询结果失败", err.Error())
	}

	logrus.Info(testEntries)
	return &testEntries, nil
}
