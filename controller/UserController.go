package controller

import (
	"abel-im/models"
	"abel-im/util"
	"html/template"
	"net/http"
)

//获取我的好友
func GetMyFriends(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	userId := request.PostForm.Get("userId")
	params := map[string]interface{}{"userId":userId}
	sql,sqlParams,_ :=models.ReadSqlParams("mapper.friends.getMyFriends",params)
	rs, _ := dbConn.GetAll(sql, sqlParams...)
	util.OK(writer, rs, "")
}

//打开注册页面
func Register(writer http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/register.html")
	tmp.Execute(writer, nil)
}

//保存用户注册
func SaveRegister(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.PostForm.Get("username")
	password := request.PostForm.Get("password")
	name := request.PostForm.Get("name")
	confirmPwd := request.PostForm.Get("confirmPwd")
	params := map[string]interface{}{"username":username}
 	rs, _ := dbConn.ExecAllSqlMapper("mapper.user.getUserByUsername", params)
	if len(rs) > 0 {
		util.Fail(writer, "手机号码已存在")
	} else {
		if password != confirmPwd {
			util.Fail(writer, "两次密码输入不一致")
		} else {
			id := models.GetStringId()
			dbConn.GetAll("INSERT INTO im_user(user_id,username,password,name) VALUES(?,?,?,?)", id, username, util.MD5(password), name)
			dbConn.GetAll("INSERT INTO im_friends(id,firend_user_id,user_id,name) VALUES(?,?,?,?)", models.GetStringId(), id, "1", name)
			dbConn.GetAll("INSERT INTO im_friends(id,firend_user_id,user_id,name) VALUES(?,?,?,?)", models.GetStringId(), "1", id, "系统管理员")
			util.OK(writer, "", "注册成功")
		}
	}

}
