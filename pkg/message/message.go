package message

import (
	"fmt"
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/binary"
	"github.com/davyxu/cellnet/util"
)

type Ping struct {
	Message string
}

type Pong struct {
	Message string
}

func NewPing() *Ping {
	return &Ping{"ping"}
}

func NewPong() *Pong {
	return &Pong{"pong"}
}

func (self *Ping) String() string { return fmt.Sprintf("%+v", *self) }
func (self *Pong) String() string { return fmt.Sprintf("%+v", *self) }

func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*Ping)(nil)).Elem(),
		ID:    int(util.StringHash("main.Ping")),
	})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*Pong)(nil)).Elem(),
		ID:    int(util.StringHash("main.Pong")),
	})
}
