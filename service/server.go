package service

import (
	"abel-im/models"
	"abel-im/util"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

//channel通道
var ToClientChan = make(chan clientInfo)

var SendMsg = make(chan RetData)

var dbConn = &models.DBConn{}

//channel通道结构体
type clientInfo struct {
	ClientId     string
	SendUserId   string //消息发送人
	ToUserId     string //消息接收人
	MessageId    string
	ProtocolPort int //消息协议端口
	Msg          string
	Data         *string
	ExtendData   map[string]interface{}
}

type RetData struct {
	ProtocolPort int                    `json:"protocolPort"`
	Desc         string                 `json:"desc"`
	Status       int                    `json:"status"`
	Date         string                 `json:"date"`
	ExtendData   map[string]interface{} `json:"extendData"`
}

var Manager = NewClientManager() // 管理者

//发送信息到指定分组
func SendMessage2Group(sendUserId, groupName string, code int, msg string, data *string) (messageId string) {
	log.Println("发送消息到制定组")
	messageId = fmt.Sprintf("%d", models.GetSnowflakeId())
	//如果是单机服务，则只发送到本机
	Manager.SendMessage2LocalGroup(messageId, sendUserId, groupName, code, msg, data)
	return
}

//消息发送到客户端
func (manager *ClientManager) Render(conn *websocket.Conn, protocolPort int, desc string, extendData map[string]interface{}) error {
	extendData["messageId"] = models.GetStringId()
	retData := RetData{ProtocolPort: protocolPort, Desc: desc, Status: 200, Date: time.Now().Format("2006-01-02 15:04:05"), ExtendData: extendData}
	switch protocolPort {
	case util.AuthProtocol: //用户登录认证
		for _, v := range manager.ClientIdMap {
			if v.Socket != conn {
				v.Socket.WriteJSON(retData)
			}
		}
		conn.WriteJSON(retData)
	case util.QuitProtocol: //用户退出
		for _, v := range manager.ClientIdMap {
			v.Socket.WriteJSON(retData)
		}
	case util.SingleMsgProtocol: //单聊天记录
		toUserId := extendData["toUserId"]
		if cn, err := manager.GetByClientId(toUserId.(string)); err == nil && conn != nil {
			cn.Socket.WriteJSON(retData)
		}
		SendMsg <- retData
		conn.WriteJSON(retData)
	case util.GroupMsgProtocol:
		groupId := extendData["groupId"]
		for _, clientId := range manager.Groups[groupId.(string)] {
			if cn, err := manager.GetByClientId(clientId); err == nil && conn != nil {
				cn.Socket.WriteJSON(retData)
			}
		}
	}
	return nil
}

//监听并发送给客户端信息
func (manager *ClientManager) WriteMessage() {
	fmt.Print("================")
	for {
		clientInfo := <-ToClientChan
		switch clientInfo.ProtocolPort {
		case util.QuitProtocol:
			clientInfo.ExtendData["userCount"] = len(manager.ClientIdMap)
			clientInfo.ExtendData["userId"] = clientInfo.ClientId
			if err := manager.Render(nil, clientInfo.ProtocolPort, clientInfo.Msg, clientInfo.ExtendData); err != nil {
				log.Println("客户端异常", err.Error())
			}
		case util.SingleMsgProtocol, util.GroupMsgProtocol, util.AuthProtocol:
			if conn, err := manager.GetByClientId(clientInfo.ClientId); err == nil && conn != nil {
				if err := manager.Render(conn.Socket, clientInfo.ProtocolPort, clientInfo.Msg, clientInfo.ExtendData); err != nil {
					Manager.DisConnect <- conn
					log.Println("客户端异常离线", err.Error())
				}
			} else {
				log.Println(" -> ", "当前在线用户：", len(manager.ClientIdMap), "。没有找到该客户端：", clientInfo.ClientId)
				//离线发送
				clientInfo.ExtendData["messageId"] = models.GetStringId()
				retData := RetData{ProtocolPort: clientInfo.ProtocolPort, Desc: clientInfo.Msg, Status: 200, Date: time.Now().Format("2006-01-02 15:04:05"), ExtendData: clientInfo.ExtendData}
				SendMsg <- retData
			}
		default:
			log.Println("消息类型无法处理  -> ", clientInfo)

		}
	}
}

//报错聊天记录
func SaveMsg() {
	for {
		select {
		case msg := <-SendMsg:
			if msg.ProtocolPort == util.SingleMsgProtocol {
				dbConn.GetAll("INSERT INTO im_messages(id,post_messages,from_user_id,to_user_id,status,create_time) values(?,?,?,?,?,?)",
					msg.ExtendData["messageId"], msg.ExtendData["msg"], msg.ExtendData["sendUserId"], msg.ExtendData["toUserId"], 0, msg.Date)
			}
		}
		/*msg := <-SendMsg
		if msg.ProtocolPort == util.SingleMsgProtocol {
			dbConn.GetAll("INSERT INTO im_messages(id,post_messages,from_user_id,to_user_id,status,create_time) values(?,?,?,?,?,?)",
				msg.ExtendData["messageId"], msg.ExtendData["msg"], msg.ExtendData["sendUserId"], msg.ExtendData["toUserId"], 0, msg.Date)
		}*/
	}
}
