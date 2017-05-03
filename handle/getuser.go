package handle

import (
	"net/http"

	"reflect"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/actionManager"
	"github.com/xw340721/webgm/callFnLibrary"
	"github.com/xw340721/webgm/iutil"
)

var Action *actionManager.Action

func init() {
	Action = actionManager.NewAction()
	Action.AddAction("getUser", callFnLibrary.GetUser)
}

//GetUser 获取玩家数据
func Main(res http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	actions, err := iutil.StringUpperIndex(r.FormValue("action"))
	if err != nil {
		logrus.Error(err.Error())
		//os.Exit(0)
	}

	err, fn := Action.GetAction(actions)
	if err != nil {
		logrus.Errorf("[处理函数] 处理函数错误 函数名为: %s ; 错误代码 %s", actions, err.Error())
		//os.Exit(0)
	}

	callBack := reflect.ValueOf(fn)
	callBackType := callBack.Type()
	if callBackType.Kind() != reflect.Func || callBackType.NumIn() != 2 {
		panic("Expected a unary function returning a single value")
	}

	returns := callBack.Call([]reflect.Value{reflect.ValueOf(res), reflect.ValueOf(r)})
	if returns[0].Interface() != nil {
		logrus.Error("[处理函数] 出现错误")
	}
}
