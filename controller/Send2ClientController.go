package controller

import (
	"abel-im/configs"
	"abel-im/service"
	"abel-im/service/grpc"
	"abel-im/util"
	"net/http"
)

//Post 发送消息到个人  一对一发送
func Send2ClientMsg(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	sendUserId := r.PostForm.Get("sendUserId")
	toUserId   := r.PostForm.Get("toUserId")
	name       := r.PostForm.Get("name")
	msgtype    := r.PostForm.Get("msgtype")
	msg        := r.PostForm.Get("msg")
	if configs.IsCluster {
		grpc.SendMessage2RemoteClient(sendUserId,toUserId,name,msgtype,msg)
	}else {
		service.SendMessage2LocalClient(sendUserId,toUserId,name,msgtype,msg)
	}
	util.OK(w, "", "发送成功")
}