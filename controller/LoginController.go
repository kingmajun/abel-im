package controller

import (
	"abel-im/util"
	"html/template"
	"net/http"
)

//打开登录页面
func Login(writer http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/login.html")
	tmp.Execute(writer, nil)
}

//登录
func UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.PostForm.Get("username")
	password := request.PostForm.Get("password")
	params := map[string]interface{}{"username":username}
	rs, _ := dbConn.ExecOneSqlMapper("mapper.user.getUserByUsername",params)
	if len(rs) > 0 {
		if util.MD5(password) != rs["password"] {
			util.Fail(writer, "密码错误")
		} else {
			util.OK(writer, rs, "")
		}
	} else {
		util.Fail(writer, "用户名不存在")
	}
}

