package main

import (
	"abel-im/configs"
	readconfig "abel-im/configs/init"
	_ "abel-im/routers"
	"abel-im/service"
	"abel-im/util"
	"fmt"
	"net/http"
)

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
	fmt.Printf("service start port: %v", configs.Port)
	http.ListenAndServe(":"+configs.Port, nil)
}

//注册etcd服务发现
func registEtcdService(port string) {
	ser, error := service.NewServiceReg(configs.EtcdEndpoints, 5)
	if error != nil {
		println("etcd connect failed, err:", error)
		return
	}
	ser.PutService("/abel-im-service/"+configs.LocalIP, configs.LocalIP+":"+port)
}
