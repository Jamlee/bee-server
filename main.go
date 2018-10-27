package main

import (
	"os"
	"strconv"
	"fmt"

	"github.com/jamlee/bee-server/server"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/hashicorp/consul/api"
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
			Usage:   "run as a server mode",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "address", Value: "127.0.0.1", Usage: "listen address",},
				cli.StringFlag{Name: "id", Value: "0001", Usage: "must special different id",},
				cli.IntFlag{Name: "port", Value: 10001, Usage: "listen port",},
      },
      Action:  func(c *cli.Context) error {
				registerMaster(c.String("address"), c.Int("port"), c.String("id"))
				server.RunControlEndpoint(fmt.Sprintf("%s:%s", c.String("address"), strconv.Itoa(c.Int("port"))))
        return nil
      },
		},
		{
      Name:    "room",
      Aliases: []string{"r"},
      Usage:   "run as a room mode",
      Action:  func(c *cli.Context) error {
        return nil
      },
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func registerMaster(address string, port int, id string) {
	serviceDef := &api.AgentServiceRegistration{
		Kind: "bee-server",
		ID: id,
		Name: "control-server",
		Tags: []string{"server"},
		Port: port,
		Address: address,
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	agent := client.Agent()
	err = agent.ServiceRegister(serviceDef)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("register to etcd successful")
}