package restype

//UserField 为玩家字段
type UserField struct {
	PlayerID    int    `json:"player_id"`
	Name        string `json:"name"`
	UserAccount int    `json:"user_account"`
}

//UserInfo 为返回玩家数据
type UserInfo struct {
	AllTotal int         `json:"all_total"`
	Fields   UserField   `json:"fields"`
	List     []UserField `json:"list"`
}

//GetUser 为返回数据+状态
type GetUser struct {
	Status bool     `json:"status"`
	Data   UserInfo `json:"data"`
}

type Test struct {
	ID     int  `json:"id,string"`
	GameId uint `json:"game_id,string"`
}
