package callFnLibrary

import (
	"net/http"

	"encoding/json"

	"github.com/xw340721/webgm/iutil"
	"github.com/xw340721/webgm/model"
	"github.com/xw340721/webgm/reqtype"
)

//GetUser 为测试案例
func GetUser(res http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	data := iutil.DecodeBase(r.FormValue("data"))

	getUser := reqtype.GetUser{}
	json.Unmarshal([]byte(data), &getUser)

	row := model.Test(getUser.ServerID)

	response, _ := json.Marshal(row)
	res.Write(response)
	return nil
}
