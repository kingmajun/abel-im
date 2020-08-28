package readconfig

import (
	"abel-im/configs"
	"github.com/astaxie/beego/config"
	"os"
	"strings"
)

var ConfigData config.Configer

func InitConfig() (err error) {
	lasTwoPath := map[string]bool{
		"readconfig":    true,
		"send2client":   true,
		"bind2group":    true,
		"send2group":    true,
		"getonlinelist": true,
		"register":      true,
		"closeclient":   true,
	}

	path, _ := os.Getwd()
	println(path)
	if strings.Contains(path, "servers") {
		path += "/.."
	} else {
		for key := range lasTwoPath {
			if strings.Contains(path, key) {
				path += "/../.."
				break
			}
		}
	}
	println(path)
	ConfigData, err = config.NewConfig("ini", path+"/configs/config.ini")
	if err != nil {
		return err
	}

	etcdHost := ConfigData.String("etcd::host")
	if len(etcdHost) > 0 {
		configs.EtcdEndpoints = make([]string, 0)
		configs.EtcdEndpoints = append(configs.EtcdEndpoints, etcdHost)
	}

	port := ConfigData.String("common::port")
	configs.Port = port

	cluster, err := ConfigData.Bool("common::cluster")
	if err != nil {
		return err
	}
	configs.IsCluster = cluster
	return
}
