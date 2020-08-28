package controller

import (
	"abel-im/models"
	"abel-im/util"
	"net/http"
)


//获取用户消息
func GetUserMsgByUserId(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	userId := request.PostForm.Get("userId")
	params := map[string]interface{}{"userId":userId}
	sql,sqlParams,_ :=models.ReadSqlParams("mapper.message.getUserMsgByUserId",params)
	rs, _ := dbConn.GetAll(sql, sqlParams...)
	util.OK(writer, rs, "")
}
