package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/gogap/logrus"
	"github.com/xw340721/webgm/config"
	"github.com/xw340721/webgm/iutil"
)

var mLogrus = logrus.WithField("package", "middleware")

//API 为整个控制节点
func API(r *http.Request) error {
	r.ParseForm()

	data := iutil.DecodeBase(r.FormValue("data"))

	auth := r.FormValue("auth")

	var apiKey = config.Get("env", "auth_key")
	h := md5.New()

	io.WriteString(h, data+apiKey)

	//转换成16位的
	hash := hex.EncodeToString(h.Sum(nil))

	if auth != hash[5:25] {
		mLogrus.Error("[API] 校验错误")
	}

	return nil
}

func LOG(r *http.Request, c martini.Context) {
	start := time.Now()
	c.Next()
	mLogrus.Info("[LOG] 请求方式:", r.Method, " 来源:", r.RemoteAddr, " action:", r.PostFormValue("action"), " during:",
		time.Since(start))
}
