package middleware

import (
	"github.com/xw340721/webgm/config"

	"net/http"

	"crypto/md5"

	"io"

	"encoding/hex"

	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/iutil"
)

//API 为整个控制节点
func API(res http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	data := iutil.DecodeBase(r.FormValue("data"))

	auth := r.FormValue("auth")

	var apiKey = config.Get("env", "auth_key")

	h := md5.New()

	io.WriteString(h, data+apiKey)

	//转换成16位的
	hash := hex.EncodeToString(h.Sum(nil))

	if auth != hash[5:25] {
		logrus.Error("[auth] 校验错误")
	}

	return nil
}
