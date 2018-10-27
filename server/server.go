package server

import (
	"github.com/sirupsen/logrus"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	peerHttp "github.com/davyxu/cellnet/peer/http"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/http"
	_ "github.com/davyxu/cellnet/codec/json"
)

type StatusMsg struct {
	Name string `json: name`
}

func RunControlEndpoint(peerAddress string) {
	queue := cellnet.NewEventQueue()
	httpAcceptor := peer.NewGenericPeer("http.Acceptor", "control-endpoint", peerAddress, nil).(cellnet.HTTPAcceptor)
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
