/*
* 发送远程消息客户端
*/
package grpc

import (
	"abel-im/configs"
	"abel-im/service/grpc/pb"
	"context"
	"google.golang.org/grpc"
	"strconv"
)

//发送远程消息
func SendMessage2RemoteClient(sendUserId string, toUserId string,name string, msgtype string, msg string) {

	for _,etcdClient := range configs.EtcdClientList {
		conn,err := grpc.Dial(etcdClient,grpc.WithInsecure())
		if err != nil {
			return
		}
		defer conn.Close() //关闭
		// 连接GRPC
		c := pb.NewGoSpiderClient(conn)
		num, _ := strconv.Atoi(msgtype)
		req := pb.Send2Client{SendUserId:sendUserId,Name:name,ToUserId:toUserId,Msgtype: int32(num),Msg:msg}
		rs, err1 := c.GetSend2ClientResponse(context.Background(),&req)
		if err1 != nil {
			return
		}
		println(rs)
	}
}