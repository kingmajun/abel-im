package controller

import (
	"abel-im/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	// 最大的消息大小
	maxMessageSize = 8192
)

var Manager = service.NewClientManager() // 管理者
//websocket 连接
func Run(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error: %v", err)
		http.NotFound(w, r)
		return
	}

	//设置读取消息大小上线
	conn.SetReadLimit(maxMessageSize)
	//
	userId := r.FormValue("userId")
	userInfo, _ := dbConn.GetOne("select * from im_user where user_id=?", userId)
	clientSocket := service.NewClient(userId, userInfo, conn)
	//读取客户端消息
	clientSocket.Read()
	// 用户连接事件
	Manager.Connect <- clientSocket
	go Manager.Start()
	go Manager.WriteMessage()

	go service.SaveMsg()

}
