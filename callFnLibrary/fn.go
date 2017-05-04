package callFnLibrary

import (
	"net/http"

	"encoding/json"

	"github.com/xw340721/webgm/iutil"
	"github.com/xw340721/webgm/model"
	"github.com/xw340721/webgm/reqtype"
	"github.com/xw340721/webgm/restype"
)

const (
	ContentType = "Content-Type"
	ContentJSON = "application/json"
)

//GetUser 为测试案例
func GetUser(res http.ResponseWriter, r *http.Request) error {
	var status bool = true

	r.ParseForm()
	users := iutil.DecodeBase(r.FormValue("data"))

	getUser := reqtype.GetUser{}
	json.Unmarshal([]byte(users), &getUser)

	test := model.Test{}
	returnData, err := test.Test(getUser.ServerID)

	res.Header().Set(ContentType, ContentJSON)
	if err != nil {
		status = false
		res.WriteHeader(400)
	} else {
		res.WriteHeader(200)
	}

	response := restype.Return{
		Status: status,
		Data:   returnData,
	}

	bytes, _ := json.Marshal(&response)
	res.Write(bytes)

	return nil
}
