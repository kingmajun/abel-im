package etcd


import (
	"abel-im/configs"
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
	"time"
)

type ClientDis struct {
	client        *clientv3.Client
}

func NewClientDis (addr []string)( *ClientDis, error){
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	if client, err := clientv3.New(conf); err == nil {
		configs.EtcdClientList = make(map[string]string)  //必须初始化后才可以赋值，否则报错
		return &ClientDis{
			client:client,
		}, nil
	} else {
		return nil ,err
	}
}


func (c * ClientDis) GetService(prefix string) ([]string ,error){
	resp, err := c.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	addrs := c.extractAddrs(resp)
	go c.watcher(prefix)
	return addrs ,nil
}


func (c *ClientDis) watcher(prefix string) {
	rch := c.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			println(ev.Type)
			switch ev.Type {
			case mvccpb.PUT:
				c.SetClientList(string(ev.Kv.Key),string(ev.Kv.Value))
			case mvccpb.DELETE:
				c.DelClientList(string(ev.Kv.Key))
			}
		}
	}
}

func (c *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string,0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			c.SetClientList(string(resp.Kvs[i].Key),string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

//增加
func (c *ClientDis) SetClientList(key,val string) {
	configs.EtcdClientListLock.Lock()
	defer configs.EtcdClientListLock.Unlock()
	configs.EtcdClientList[key] = val
	log.Println("set data key :",key,"val:",val)
}

//删除客户端
func (c *ClientDis) DelClientList(key string) {
	configs.EtcdClientListLock.Lock()
	defer configs.EtcdClientListLock.Unlock()
	delete(configs.EtcdClientList,key)
	log.Println("del data key:", key)
}

