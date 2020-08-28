package controller

import (
	"abel-im/models"
	"abel-im/util"
	"net/http"
)

//查询我的群
func GetMyGroupList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userId := r.PostForm.Get("userId")
	params := map[string]interface{}{"userId":userId}
	sql,sqlParams,_ :=models.ReadSqlParams("mapper.group.getMyGroupsByUserId",params)
	rows, _ := dbConn.GetAll(sql, sqlParams...)
	util.OK(w, rows, "")
}

//查询群消息
func GetGroupMsgList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userId := r.PostForm.Get("userId")
	params := map[string]interface{}{"userId":userId}
	sql,sqlParams,_ :=models.ReadSqlParams("mapper.message.getMyGroupMsgsByUserId",params)
	rows, _ := dbConn.GetAll(sql, sqlParams...)
	util.OK(w, rows, "")
}
