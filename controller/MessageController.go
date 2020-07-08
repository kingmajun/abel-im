package controller

import (
	"abel-im/util"
	"net/http"
)

func GetUserMsgByUserId(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	userId := request.PostForm.Get("userId")
	rs, _ := dbConn.GetAll("SELECT i.id,i.post_messages,i.from_user_id,i.to_user_id,i.status,i.create_time,"+
		" (select name from im_user where user_id=i.from_user_id)name"+
		"	FROM im_messages i where ( i.from_user_id=? or i.to_user_id=?) ORDER BY  i.create_time", userId, userId)
	util.OK(writer, rs, "")
}
