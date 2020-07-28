package util

import (
	readconfig "abel-im/configs/init"
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"net"
	"strconv"
	"strings"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//判断数组种是否包含元素
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//GenUUID 生成uuid
func GenUUID() string {
	uuidFunc := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

//是否集群
func IsCluster() bool {
	cluster, _ := readconfig.ConfigData.Bool("common::cluster")
	return cluster
}

//生成RPC通信端口号，目前是ws端口号+1000
func GenRpcPort(port string) string {
	iPort, _ := strconv.Atoi(port)
	return strconv.Itoa(iPort + 1000)
}

//获取本机内网IP
func GetIntranetIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GenGroupKey(systemId, groupName string) string {
	return systemId + ":" + groupName
}
