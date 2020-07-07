package service

import (
	"abel-im/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

type Client struct {
	ClientId    string          // 用户Id
	Socket      *websocket.Conn // 用户连接
	ConnectTime uint64          // 首次连接时间
	IsDeleted   bool            // 是否删除或下线
	UserInfo    map[string]interface{}
	Extend      string          // 扩展字段，用户可以自定义
	GroupList   []string        //存储用户所有群id
}


func NewClient(clientId string,UserInfo map[string]interface{},socket *websocket.Conn) *Client {
	return &Client{
		ClientId:    clientId,
		UserInfo:    UserInfo,
		Socket:      socket,
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
	}
}


func (c *Client) Read() {
	go func() {
		for {
			messageType, message, err := c.Socket.ReadMessage()
			if err != nil {
				if messageType == -1 && websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
				 	Manager.DisConnect <- c
				 	fmt.Println(" ->退出 ")
					return
				} else if messageType != websocket.PingMessage {
					return
				}
			}
			log.Println(c.UserInfo["user_id"]," 接收到客户端消息 -> ",string(message))
			// 解析json文本
			m := make(map[string]interface{})
			if json.Unmarshal(message, &m) != nil {
				return
			}
			msg,_ := m["msg"].(string)
			msgtype,_ := strconv.Atoi(m["msgType"].(string))
			protocolPort,_ := strconv.Atoi(m["protocolPort"].(string))

			switch protocolPort {
			case util.SingleMsgProtocol:
				toUserId,_ := m["toUserId"].(string)
				var extendData = make(map[string]interface{})
				extendData["msg"] = msg
				extendData["msgtype"] = msgtype
				extendData["sendUserId"] = c.ClientId
				extendData["toUserId"] = toUserId
				extendData["name"] = c.UserInfo["name"]
				ToClientChan  <- clientInfo{ClientId:c.ClientId,Msg:"单聊消息",ProtocolPort:protocolPort,ExtendData:extendData}
			case util.GroupMsgProtocol:
				groupId,_ := m["groupId"].(string)
				var extendData = make(map[string]interface{})
				extendData["msg"] = msg
				extendData["msgtype"] = msgtype
				extendData["sendUserId"] = c.ClientId
				extendData["name"] = c.UserInfo["name"]
				extendData["groupId"] = groupId
				ToClientChan  <- clientInfo{ClientId:c.ClientId,Msg:"群聊消息",ProtocolPort:protocolPort,ExtendData:extendData}
			}
		}
	}()
}



