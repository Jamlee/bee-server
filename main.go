package main

import (
	"strconv"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	// init the cellnet lib global variable
	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/codec/binary"
	_ "github.com/davyxu/cellnet/proc/http"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/jamlee/bee-server/server"
	"github.com/jamlee/bee-server/worker"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "bee-server"
	app.Version = VERSION
	app.Usage = "bee-server is a distrbuted server framework for real time game"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "d", Usage: "open debug log",},
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
				cli.StringFlag{Name: "address", Value: "127.0.0.1", Usage: "listen address",},
				cli.IntFlag{Name: "web-port", Value: 10001, Usage: "listen port",},
				cli.IntFlag{Name: "master-port", Value: 10002, Usage: "listen port",},
      },
      Action:  func(c *cli.Context) error {
				server.RegisterMaster(c.String("address"), c.Int("master-port"))
				server.RunMasterEndpoint(fmt.Sprintf("%s:%s", c.String("address"), strconv.Itoa(c.Int("master-port"))))
				server.RunControlEndpoint(fmt.Sprintf("%s:%s", c.String("address"), strconv.Itoa(c.Int("web-port"))))
        return nil
      },
		},
		{
      Name:    "worker",
      Aliases: []string{"w"},
			Usage:   "run as a worker role",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "server-address", Value: "127.0.0.1", Usage: "listen address",},
				cli.IntFlag{Name: "server-port", Value: 10002, Usage: "listen port",},
      },
      Action:  func(c *cli.Context) error {
				worker.RegisterWorker(c.String("server-address"), c.Int("server-port"))
				worker.RunWorkerClient(fmt.Sprintf("%s:%s", c.String("server-address"), strconv.Itoa(c.Int("server-port"))))
				return nil
			},
		},
	}
	app.Before = func(c *cli.Context) error {
		logrus.SetLevel(logrus.WarnLevel)
    return nil
  }

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
