package request

//API 接口统一地址
type API struct {
	Action string `json:"action"`
	Data   string `json:"data"`
	Auth   string `json:"auth"`
}

//GetUser 查询服务器玩家列表
type GetUser struct {
	Page         int      `json:"page,string"`
	Size         int      `json:"size,string"`
	ServerID     int      `json:"server_id,string"`
	PlatfromID   int      `json:"platform_id,string"`
	OpID         int      `json:"op_id,string"`
	PlayerID     string   `json:"player_id,string"`
	UserAccount  string   `json:"user_account,string"`
	Name         string   `json:"name,string"`
	Gm           int      `json:"gm,string"`
	Cash         []int    `json:"cash,[]string"`
	Vip          []int    `json:"vip,[]string"`
	Level        []int    `json:"level,[]string"`
	CreatedAt    []string `json:"created_at"`
	RechargeTime []string `json:"recharge_time"`
}
