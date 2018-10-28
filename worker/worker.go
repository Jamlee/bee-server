package worker

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"

	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
)


func RunWorkerClient(peerAddress string) {
	queue := cellnet.NewEventQueue()
	tcpConnector := peer.NewGenericPeer("tcp.Connector", "work-client", peerAddress, queue)
	proc.BindProcessorHandler(tcpConnector, "tcp.ltv", func(ev cellnet.Event) {
	})
	tcpConnector.Start()
	queue.StartLoop()
}
