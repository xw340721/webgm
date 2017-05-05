package mutations

import (
	"encoding/json"
	"net/http"

	"github.com/xw340721/webgm/iutil"
	"github.com/xw340721/webgm/model"
	"github.com/xw340721/webgm/request"
	"github.com/xw340721/webgm/response"
)

const (
	ContentType = "Content-Type"
	ContentJSON = "application/json"
	Charset     = "charset"
	UTF_8       = "utf-8"
)

//GetUser 为测试案例
func TestDemo(res http.ResponseWriter, r *http.Request) error {
	var status bool = true

	r.ParseForm()
	users := iutil.DecodeBase(r.FormValue("data"))

	getUser := request.GetUser{}
	json.Unmarshal([]byte(users), &getUser)

	test := model.Test{}
	returnData, err := test.Test(getUser.ServerID)

	res.Header().Set(ContentType, ContentJSON)
	res.Header().Set(Charset, UTF_8)
	if err != nil {
		status = false
		res.WriteHeader(404)
	} else {
		res.WriteHeader(200)
	}

	ret := response.Return{
		Status: status,
		Data:   returnData,
	}

	bytes, _ := json.Marshal(&ret)
	res.Write(bytes)

	return nil
}
