package project03

import (
	"io"
	"net/http"
)

func HelloPage(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	io.WriteString(res, "hello, "+name)
}

func PostAction(res http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	io.WriteString(res, "注册成功! 用户名: "+username)
}
