package callFnLibrary

import (
	"net/http"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/iutil"
)

func GetUser(res http.ResponseWriter, r *http.Request) string {
	r.ParseForm()
	data := iutil.DecodeBase( r.FormValue("data"))





	logrus.Info(data)
	return "fdas"
}
