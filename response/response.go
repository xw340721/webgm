package response

//返回数据+状态
type Return struct {
	Status bool        `json:"status,int"`
	Data   interface{} `json:"data"`
}
