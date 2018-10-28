package worker

import (
	"github.com/jamlee/bee-server/pkg"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"
	"github.com/sirupsen/logrus"
)


func RunWorkerClient(peerAddress string) {
	queue := cellnet.NewEventQueue()
	tcpConnector := peer.NewGenericPeer("tcp.Connector", "work-client", peerAddress, queue)
	proc.BindProcessorHandler(tcpConnector, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
			case *cellnet.SessionConnected: // 已经连接上
				logrus.Debug("client connected")
				ev.Session().Send(pkg.NewPing())
			case *pkg.Pong:
				logrus.Infof("client recv %+v\n", msg)
			case *cellnet.SessionClosed:
				logrus.Info("client closed")
			}
	})
	tcpConnector.Start()
	queue.StartLoop()
	queue.Wait()
}
