package server

import (
	"github.com/sirupsen/logrus"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"

	peerHttp "github.com/davyxu/cellnet/peer/http"
	_ "github.com/davyxu/cellnet/proc/http"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
	_ "github.com/davyxu/cellnet/codec/json"
)

type StatusMsg struct {
	Name string `json: name`
}

func RunControlEndpoint(peerAddress string) {
	queue := cellnet.NewEventQueue()
	httpAcceptor := peer.NewGenericPeer("http.Acceptor", "web-endpoint", peerAddress, nil).(cellnet.HTTPAcceptor)
	proc.BindProcessorHandler(httpAcceptor, "http", func(ev cellnet.Event) {
		httpSession := ev.(*cellnet.RecvMsgEvent).Session()
		result := &peerHttp.MessageRespond{
			StatusCode: 200,
			Msg: &StatusMsg{Name: "jamlee"},
			CodecName: "json",
		}
		logrus.Info(result.String())
		httpSession.Send(result)
	})
	httpAcceptor.Start()
	queue.StartLoop()
	queue.Wait()
}

func RunMasterEndpoint(peerAddress string) {
	queue := cellnet.NewEventQueue()
	tcpAcceptor := peer.NewGenericPeer("tcp.Acceptor", "master-endpoint", peerAddress, queue)
	proc.BindProcessorHandler(tcpAcceptor, "tcp.ltv", func(ev cellnet.Event) {
	})
	tcpAcceptor.Start()
	queue.StartLoop()
}
