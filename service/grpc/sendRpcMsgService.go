/*
* 发送远程消息服务端
 */
package grpc

import (
	"abel-im/service"
	"abel-im/service/grpc/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type rpcService struct{}

//提供一个远程调用的grpc调用接口
func (s *rpcService) GetSend2ClientResponse(ctx context.Context,c *pb.Send2Client) (*pb.Send2ClientReply,error) {
	sendUserId  := c.SendUserId
	name 		:= c.Name
	toUserId 	:= c.ToUserId
	msgtype  	:= c.Msgtype
	msg 		:= c.Msg
	service.SendMessage2LocalClient(sendUserId,toUserId,name,string(msgtype),msg)
	rsp := pb.Send2ClientReply{HttpCode:200,Response:"消息发送成功"}
	return &rsp,nil
}

//创建GRPC服务  address = ip+":"+端口
func CreateGRPCService(address string)  {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	s := grpc.NewServer()// 创建GRPC
	pb.RegisterGoSpiderServer(s,&rpcService{})
	reflection.Register(s)
	err = s.Serve(listener)
	if err != nil {
		return
	}
}