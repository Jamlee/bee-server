package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go.etcd.io/etcd/embed"

	// init the cellnet lib global variable
	_ "github.com/davyxu/cellnet/codec/binary"
	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/http"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/jamlee/bee-server/pkg/control"
	"github.com/jamlee/bee-server/pkg/server"
	"github.com/jamlee/bee-server/pkg/worker"
)

var VERSION = "v0.0.0-dev"

func main() {
	waitEtcd := make(chan struct{})
	stop := make(chan struct{})

	app := cli.NewApp()
	app.Name = "bee-server"
	app.Version = VERSION
	app.Usage = "bee-server is a distrbuted server framework for real time game"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "d", Usage: "open debug log"},
		cli.StringFlag{Name: "initial-cluster", Usage: "etcd initial cluster:" +
			"etcd-1=http://127.0.0.1:2380,etcd-2=http://127.0.0.1:2380,etcd-3=http://127.0.0.1:2380"},
		cli.StringFlag{Name: "initial-cluster-token", Usage: "etcd-token"},
		cli.StringFlag{Name: "peer-url", Usage: "for example: http://localhost:2380", Value: "http://localhost:2380"},
		cli.StringFlag{Name: "advertise-client-url", Usage: "for example: http://localhost:2379", Value: "http://localhost:2379"},
		cli.StringFlag{Name: "name", Usage: "for example: default", Value: "default"},
	}
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "jamlee",
			Email: "jamlee@jamlee.cn",
		},
	}
	app.Copyright = "(c) 2018 Jam Lee"
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "run as a server role",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "address", Value: "127.0.0.1", Usage: "listen address"},
				cli.IntFlag{Name: "web-port", Value: 10001, Usage: "listen port"},
				cli.IntFlag{Name: "master-port", Value: 10002, Usage: "listen port"},
			},
			Action: func(c *cli.Context) error {
				go startEtcd(c.GlobalString("name"), c.GlobalString("peer-url"), c.GlobalString("advertise-client-url"),
					c.String("initial-cluster"), c.String("initial-cluster-token"), waitEtcd, stop)
				<-waitEtcd
				go control.StartServer(fmt.Sprintf("%s:%s", c.String("address"), strconv.Itoa(c.Int("web-port"))))
				server.RegisterMaster(c.String("address"), c.Int("master-port"))
				server.RunMasterEndpoint(fmt.Sprintf("%s:%s", c.String("address"), strconv.Itoa(c.Int("master-port"))))
				return nil
			},
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "run as a worker role",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "server-address", Value: "127.0.0.1", Usage: "listen address"},
				cli.IntFlag{Name: "server-port", Value: 10002, Usage: "listen port"},
			},
			Action: func(c *cli.Context) error {
				go startEtcd(c.GlobalString("name"), c.GlobalString("peer-url"), c.GlobalString("advertise-client-url"),
					c.String("initial-cluster"), c.String("initial-cluster-token"), waitEtcd, stop)
				<-waitEtcd
				worker.RegisterWorker(c.String("server-address"), c.Int("server-port"))
				worker.RunWorkerClient(fmt.Sprintf("%s:%s", c.String("server-address"), strconv.Itoa(c.Int("server-port"))))
				return nil
			},
		},
	}
	app.Before = func(c *cli.Context) error {
		logrus.SetLevel(logrus.InfoLevel)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func startEtcd(name string, advertiseClientURL string, peerURL string,
	cluster string, clusterToken string, isReady chan struct{},
	stop chan struct{}) {

	lpUrl, err := url.Parse(peerURL)
	acUrl, err := url.Parse(advertiseClientURL)
	if err != err {
		logrus.Fatal(err)
	}

	cfg := embed.NewConfig()
	cfg.Name = name
	cfg.InitialCluster = cluster
	cfg.InitialClusterToken = clusterToken
	cfg.LPUrls = []url.URL{*acUrl}
	cfg.LCUrls = []url.URL{*lpUrl}
	cfg.ACUrls = []url.URL{*lpUrl}
	cfg.Dir = "/tmp/etcd/" + name
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer e.Close()
	select {
	case <-e.Server.ReadyNotify():
		logrus.Info("Etcd Server is ready!")
		isReady <- struct{}{}
	case <-stop:
		e.Server.Stop()
	}
	logrus.Fatal(<-e.Err())
}
