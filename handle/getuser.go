package handle

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/actionManager"
	"github.com/xw340721/webgm/iutil"
	"github.com/xw340721/webgm/mutations"
	"github.com/xw340721/webgm/response"
)

var Action *actionManager.Action
var mLogrus = logrus.WithField("package", "handle")

func init() {
	Action = actionManager.NewAction()
	Action.AddAction("getUser", mutations.TestDemo)
}

//Main 主handle
func Main(res http.ResponseWriter, r *http.Request) {
	var errorCollection []error
	r.ParseForm()
	actions, err := iutil.StringUpperIndex(r.FormValue("action"))
	if err != nil {
		logrus.Error(err.Error())
	}

	err, fn := Action.GetAction(actions)
	if err != nil {
		mLogrus.Errorf("[Main] 处理函数错误 函数名为: %s ", actions)
		errorCollection = append(errorCollection, err)
	} else {
		callBack := reflect.ValueOf(fn)
		callBackType := callBack.Type()
		if callBackType.Kind() != reflect.Func || callBackType.NumIn() != 2 {
			mLogrus.Error("[Main] 被调用函数不符合要求")
			errorCollection = append(errorCollection, err)
		}

		returns := callBack.Call([]reflect.Value{reflect.ValueOf(res), reflect.ValueOf(r)})
		if returns[0].Interface() != nil {
			mLogrus.Error("[Main] 出现错误")
			errorCollection = append(errorCollection, err)
		}
	}

	if len(errorCollection) != 0 {
		ret := response.Return{
			Status: false,
		}
		data, _ := json.Marshal(ret)

		res.WriteHeader(404)
		res.Write(data)
	}

}
