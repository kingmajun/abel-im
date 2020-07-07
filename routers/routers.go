package routers

import (
	"abel-im/controller"
	"net/http"
)

func init()  {
	http.HandleFunc("/",controller.Login)
	http.HandleFunc("/userLogin",controller.UserLogin)
	http.HandleFunc("/main",controller.Main)
	http.HandleFunc("/register",controller.Register)

	//查询好友
	http.HandleFunc("/api/getMyFriends",controller.GetMyFriends)
	http.HandleFunc("/api/getUserMsgByUserId",controller.GetUserMsgByUserId)
	http.HandleFunc("/api/saveRegister",controller.SaveRegister)
	http.HandleFunc("/api/getMyGroupList",controller.GetMyGroupList)
	http.HandleFunc("/api/GetGroupMsgList",controller.GetGroupMsgList)

	http.HandleFunc("/SendMsg",controller.SendMsg)
	http.HandleFunc("/ws",controller.Run)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}