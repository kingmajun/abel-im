package configs

import "sync"

//系统监听端口
var Port string

//本机IP
var LocalIP string

//是否集群
var IsCluster bool

//etcd服务
var EtcdEndpoints   []string
// 注册到etcd客户端
var EtcdClientList     map[string]string
var EtcdClientListLock sync.RWMutex