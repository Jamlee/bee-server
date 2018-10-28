package server

import (
	"github.com/jamlee/bee-server/pkg"
	"github.com/sirupsen/logrus"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	"github.com/davyxu/cellnet/proc"

	peerHttp "github.com/davyxu/cellnet/peer/http"
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
		switch msg := ev.Message().(type) {
			case *cellnet.SessionAccepted: 
				logrus.Debug("server accepted")
			case *pkg.Ping:
				logrus.Debugf("client recv %+v\n", msg)
			case *cellnet.SessionClosed:
				logrus.Debugf("session closed: ", ev.Session().ID())
			}
		})
	tcpAcceptor.Start()
	queue.StartLoop()
}
