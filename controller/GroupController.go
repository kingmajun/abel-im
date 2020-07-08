package controller

import (
	"abel-im/util"
	"net/http"
)

//查询我的群
func GetMyGroupList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userId := r.PostForm.Get("userId")
	rows, _ := dbConn.GetAll("select  a.groups_id,b.group_name from  im_groups_to_user a,im_user_groups b where a.groups_id=b.id and a.user_id=?", userId)
	util.OK(w, rows, "")
}

//查询群消息
func GetGroupMsgList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userId := r.PostForm.Get("userId")
	rows, _ := dbConn.GetAll("select a.* from im_groups_messages a,im_groups_to_user b where a.groups_id=b.groups_id and b.user_id=? ORDER BY a.create_time", userId)
	util.OK(w, rows, "")
}
