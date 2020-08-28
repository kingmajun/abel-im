package service

import (
	"abel-im/util"
	"errors"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
)

// 连接管理
type ClientManager struct {
	ClientIdMap     map[string]*Client // 全部的连接
	ClientIdMapLock sync.RWMutex       // 读写锁

	Connect    chan *Client // 连接处理
	DisConnect chan *Client // 断开连接处理

	GroupLock sync.RWMutex
	Groups    map[string][]string
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		ClientIdMap: make(map[string]*Client),
		Connect:     make(chan *Client, 10000),
		DisConnect:  make(chan *Client, 10000),
		Groups:      make(map[string][]string, 100),
	}
	return
}

// 管道处理程序
func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.Connect:
			// 建立连接事件
			manager.EventConnect(client)
		case conn := <-Manager.DisConnect:
			// 断开连接事件
			manager.EventDisconnect(conn)
		}
	}
}

// 建立连接事件
func (manager *ClientManager) EventConnect(client *Client) {
	manager.AddClient(client)
	manager.AddUserGroupList(client)
	logrus.WithFields(logrus.Fields{
		"host":     "localhost",
		"port":     "8080",
		"clientId": client.ClientId,
		"counts":   manager.Count(),
	}).Info("客户端已连接")
}

//增加用户群组信息
func (manager *ClientManager) AddUserGroupList(client *Client) {
	userId := client.ClientId
	rows, _ := dbConn.GetAll("select  groups_id from  im_groups_to_user where user_id=?", userId)
	for _, row := range rows {
		groupId := row["groups_id"].(string)
		client.GroupList = append(client.GroupList, groupId)
		//将个人信息加入到群组中
		if !util.IsContain(manager.Groups[groupId], client.ClientId) {
			manager.Groups[groupId] = append(manager.Groups[groupId], client.ClientId)
		}
	}
}

// 断开连接事件
func (manager *ClientManager) EventDisconnect(client *Client) {
	//关闭连接
	_ = client.Socket.Close()
	manager.DelClient(client)
	var extendData = make(map[string]interface{})
	//发送下线通知
	ToClientChan <- clientInfo{ClientId: client.ClientId, Msg: "下线通知", ProtocolPort: util.QuitProtocol, ExtendData: extendData}

	//标记销毁
	client.IsDeleted = true
	client = nil
}

// 添加客户端
func (manager *ClientManager) AddClient(client *Client) {
	manager.ClientIdMapLock.Lock()
	defer manager.ClientIdMapLock.Unlock()
	manager.ClientIdMap[client.ClientId] = client
	//定义map并初始化
	var extendData = make(map[string]interface{})
	extendData["userCount"] = len(manager.ClientIdMap)
	ToClientChan <- clientInfo{ClientId: client.ClientId, Msg: "上线通知", ProtocolPort: util.AuthProtocol, ExtendData: extendData}
}

// 客户端数量
func (manager *ClientManager) Count() int {
	manager.ClientIdMapLock.RLock()
	defer manager.ClientIdMapLock.RUnlock()
	return len(manager.ClientIdMap)
}

// 删除客户端
func (manager *ClientManager) DelClient(client *Client) {
	manager.delClientIdMap(client.ClientId)
	//删除所在的分组
	if len(client.GroupList) > 0 {
		for _, groupName := range client.GroupList {
			manager.delGroupClient(util.GenGroupKey("", groupName), client.ClientId)
		}
	}

}

// 删除分组里的客户端
func (manager *ClientManager) delGroupClient(groupKey string, clientId string) {
	manager.GroupLock.Lock()
	defer manager.GroupLock.Unlock()
	for index, groupClientId := range manager.Groups[groupKey] {
		if groupClientId == clientId {
			manager.Groups[groupKey] = append(manager.Groups[groupKey][:index], manager.Groups[groupKey][index+1:]...)
		}
	}
}

// 删除clientIdMap
func (manager *ClientManager) delClientIdMap(clientId string) {
	manager.ClientIdMapLock.Lock()
	defer manager.ClientIdMapLock.Unlock()

	delete(manager.ClientIdMap, clientId)
}

// 通过clientId获取
func (manager *ClientManager) GetByClientId(clientId string) (*Client, error) {
	manager.ClientIdMapLock.RLock()
	defer manager.ClientIdMapLock.RUnlock()

	if client, ok := manager.ClientIdMap[clientId]; !ok {
		return nil, errors.New("客户端不存在")
	} else {
		return client, nil
	}
}

// 发送到本机分组
func (manager *ClientManager) SendMessage2LocalGroup(messageId, sendUserId, groupName string, code int, msg string, data *string) {
	if len(groupName) > 0 {
		clientIds := manager.GetGroupClientList(util.GenGroupKey("", groupName))
		if len(clientIds) > 0 {
			for _, clientId := range clientIds {
				if _, err := Manager.GetByClientId(clientId); err == nil {
					//添加到本地
					//SendMessage2LocalClient(messageId, clientId, sendUserId, code, msg, data)
				} else {
					//删除分组
					manager.delGroupClient(util.GenGroupKey("", groupName), clientId)
				}
			}
		}
	}
}

// 获取本地分组的成员
func (manager *ClientManager) GetGroupClientList(groupKey string) []string {
	manager.GroupLock.RLock()
	defer manager.GroupLock.RUnlock()
	return manager.Groups[groupKey]
}

//通过本服务器发送信息
func SendMessage2LocalClient(sendUserId string, toUserId string,name string, msgtype string, msg string) {
	log.Println("发送到通道")
	var extendData = make(map[string]interface{})
	extendData["msg"] = msg
	extendData["msgtype"] = msgtype
	extendData["sendUserId"] = sendUserId
	extendData["toUserId"] = toUserId
	extendData["name"] = name

	//go Manager.WriteMessage() //todo 临时解决，当所有人都没登录情况下，发送消息导致管道消息没有被消费，从而导致线程阻塞问题
	//go SaveMsg() //todo 临时解决
	ToClientChan <- clientInfo{ClientId: sendUserId, Msg: "单聊消息", ProtocolPort: util.SingleMsgProtocol, ExtendData: extendData}

	return
}
