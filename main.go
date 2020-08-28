package main

import (
	"abel-im/configs"
	readconfig "abel-im/configs/init"
	"abel-im/models"
	_ "abel-im/routers"
	"abel-im/service/etcd"
	"abel-im/service/grpc"
	"abel-im/util"
	"fmt"
	orm "github.com/kingmajun/king-orm"
	"net/http"
	"os"
)

func init()  {
	ROOT,_ := os.Getwd()
	path := ROOT+"/mapper"
	models.SqlMapper,_ = orm.ReaderConfigBuilder(path)
}

func main() {
	err := readconfig.InitConfig()
	if err != nil {
		println("ini配置文件读取异常")
		return
	}
	configs.LocalIP = util.GetIntranetIp()
	fmt.Printf("服务端IP：%v \n", configs.LocalIP)
	//把本机IP：端口注册到ETCD上
	registEtcdService(configs.Port)
	initGRPCService()

	fmt.Printf("service start port: %v", configs.Port)
	http.ListenAndServe(":"+configs.Port, nil)
}

//注册etcd服务发现
func registEtcdService(port string) {
	if configs.IsCluster {
		ser, error := etcd.NewServiceReg(configs.EtcdEndpoints, 5)
		if error != nil {
			println("etcd connect failed, err:", error)
			return
		}
		ser.PutService("/abel-im-service/"+configs.LocalIP+":"+port, configs.LocalIP+":"+port)

		cli,_ := etcd.NewClientDis(configs.EtcdEndpoints)
		cli.GetService("/abel-im-service")
	}
}

//初始化grpc服务
func initGRPCService()  {
	if configs.IsCluster {
		go grpc.CreateGRPCService(configs.LocalIP+":"+configs.Port)
	}
}
