package main

import (
	"strconv"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/jamlee/bee-server/server"
	"github.com/jamlee/bee-server/worker"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "bee-server"
	app.Version = VERSION
	app.Usage = "bee-server is a distrbuted server framework for real time game"
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
				cli.StringFlag{Name: "address", Value: "127.0.0.1", Usage: "listen address",},
				cli.IntFlag{Name: "worker-port", Value: 10003, Usage: "listen port",},
      },
      Action:  func(c *cli.Context) error {
				worker.RegisterWorker(c.String("address"), c.Int("worker-port"))
				worker.RunWorkerClient(fmt.Sprintf("%s:%s", c.String("address"), strconv.Itoa(c.Int("worker-port"))))
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
