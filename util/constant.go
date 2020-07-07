package util


//消息类型 MsgType
const (
	//发送文本聊天消息
	TextMsgType = 2001
	//语音消息
	VoiceMsgType = 2002
	//图片
	PicMsgType   = 2003
)
//消息协议端口 ProtocolPort
const (
	//系统消息
	SysProtocol = 10000
	//用户认证消息
	AuthProtocol = 10001
	//用户退出
	QuitProtocol = 10004
	//群消息
	GroupMsgProtocol = 10002
	//单聊
	SingleMsgProtocol = 10003
)
