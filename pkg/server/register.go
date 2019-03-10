package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

type Service struct {
	ID   string
	Type string
	Info Serverinfo
}

type Serverinfo struct {
	Id   string `json:"id"`
	Bind string `json:"Bind"`
	Port int    `json:"port"`
}

func RegisterMaster(address string, port int) {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"127.0.0.1:2379"}})
	if err != nil {
		logrus.Fatal("Error: cannot connec to etcd:", err)
	}
	prefix := "/services/server"
	key := fmt.Sprintf("%s/%s:%d", prefix, address, port)
	s := &Service{
		ID:   key,
		Type: "server",
		Info: Serverinfo{Id: key, Bind: address, Port: port},
	}

	// register to etcd
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		logrus.Fatal(err)
	}

	// write health key
	healthKey := fmt.Sprintf("%s/health", key)
	_, err = cli.Put(context.TODO(), healthKey, "")
	if err != nil {
		logrus.Fatal(err)
	}

	// write info key
	infoKey := fmt.Sprintf("%s/info", key)
	info, err := json.Marshal(s.Info)
	_, err = cli.Put(context.TODO(), infoKey, string(info), clientv3.WithLease(resp.ID))
	if err != nil {
		logrus.Fatal(err)
	}
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		logrus.Fatal(kaerr)
	}
	ka := <-ch
	logrus.Infof("[%s] register to etcd, heartbeat interval is %d second ", key, ka.TTL)
}
