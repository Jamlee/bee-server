package server

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/jamlee/bee-server/pkg/message"
	"github.com/sirupsen/logrus"
)

type StatusMsg struct {
	Name string `json:"name"`
}

func RunMasterEndpoint(peerAddress string) {
	queue := cellnet.NewEventQueue()
	tcpAcceptor := peer.NewGenericPeer("tcp.Acceptor", "master-endpoint", peerAddress, queue)
	proc.BindProcessorHandler(tcpAcceptor, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionAccepted:
			logrus.Debug("server accepted")
		case *message.Ping:
			logrus.Debugf("client recv %+v\n", msg)
			ev.Session().Send(message.NewPong())
		case *cellnet.SessionClosed:
			logrus.Debugf("session closed: %d", ev.Session().ID())
		}
	})
	tcpAcceptor.Start()
	queue.StartLoop()
	queue.Wait()
}
