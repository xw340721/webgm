package model

import (
	"encoding/json"

	"fmt"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/sql"
)

type Test struct {
	ID     int  `json:"id"`
	GameId uint `json:"game_id"`
}
type TestData struct {
	AllTotal int               `json:"all_total"`
	Field    map[string]string `json:"fields"`
	List     []Test            `json:"list"`
}

var mLogrus = logrus.WithField("package", "model")

//GetUser 返回查询的玩家数量
func (t *Test) Test(serverID int) (interface{}, error) {
	conn := NewConn()
	defer conn.Close()

	query := fmt.Sprintf(sql.NormalSelect, "id,game_id", "server", "server_id = ?")

	//query
	stmt, err := conn.Prepare(query)
	if err != nil {
		mLogrus.Error("[test] 准备解析失败", err.Error())
	}
	rows, err := stmt.Query(serverID)
	defer rows.Close()

	if err != nil {
		mLogrus.Error("[test] 查询失败", err.Error())
	}

	//分析columns
	columns, err := rows.Columns()

	if err != nil {
		mLogrus.Error("[test] 解析columns失败")
	}

	data, num, _ := ReturnToJson(columns, rows)

	//创建field
	field := make(map[string]string)
	field["game_id"] = "游戏ID"
	field["id"] = "ID"

	//构造返回值
	testEntries := TestData{
		AllTotal: num,
		Field:    field,
	}

	json.Unmarshal(data, &testEntries.List)

	err = rows.Err()
	if err != nil {
		mLogrus.Error("[mysql] 查询结果失败", err.Error())
	}

	return &testEntries, nil
}
