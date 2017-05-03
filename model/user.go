package model

import (
	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/restype"
)

//GetUser 返回查询的玩家数量
func Test(serverID int) restype.Test {
	var (
		id      int
		game_id uint
	)

	conn := NewConn()

	defer conn.Close()
	stmt, err := conn.Prepare("SELECT id,game_id FROM server WHERE server_id = ? ")
	if err != nil {
		logrus.Error("[mysql] 准备解析失败", err.Error())
	}
	rows, err := stmt.Query(serverID)
	if err != nil {
		logrus.Error("[mysql] 查询失败", err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&id, &game_id)
		if err != nil {
			logrus.Error("[mysql] 查询结果解析失败", err.Error())
		}
	}
	err = rows.Err()
	if err != nil {
		logrus.Error("[mysql] 查询结果失败", err.Error())
	}

	return restype.Test{ID: id, GameId: game_id}
}
