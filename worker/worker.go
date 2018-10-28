package worker

import (
	"time"
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
				go HeartBeat(ev)
			case *pkg.Pong:
				logrus.Debugf("client recv %+v\n", msg)
			case *cellnet.SessionClosed:
				logrus.Debug("client closed")
			}
	})
	tcpConnector.Start()
	queue.StartLoop()
	queue.Wait()
}

// send heartbeat to master
func HeartBeat(ev cellnet.Event) {
	for {
		ev.Session().Send(pkg.NewPing())
		time.Sleep(5 * time.Second)
	}
}