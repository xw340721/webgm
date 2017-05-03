package reqtype

//API 接口统一地址
type API struct {
	Action string `json:"action"`
	Data   string `json:"data"`
	Auth   string `json:"auth"`
}

//GetUser 查询服务器玩家列表
type GetUser struct {
	Page         int      `json:"page"`
	Size         int      `json:"size"`
	ServerID     string   `json:"server_id"`
	PlatfromID   int      `json:"platform_id"`
	OpID         int      `json:"op_id"`
	PlayerID     string   `json:"player_id"`
	UserAccount  string   `json:"user_account"`
	Name         string   `json:"name"`
	Gm           int      `json:"gm"`
	Cash         []int    `json:"cash"`
	Vip          []int    `json:"vip"`
	Level        []int    `json:"level"`
	CreatedAt    []string `json:"created_at"`
	RechargeTime []string `json:"recharge_time"`
}
