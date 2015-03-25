package logsend

import (
/*
	"flag"
	"strings"
	"time"
*/
)

var (
	restapiCh = make(chan *[]string, 0)
)

func init() {
	RegisterNewSender("restapi", InitRestapi, NewRestapiSender)
}

func InitRestapi(conf interface{}) {
	go func() {
		for data := range restapiCh {
			debug("write by api", data)
		}
	}()
	return
}

func (self *RestapiSender) Name() string {
	return "restapi"
}

type RestapiSender struct {
	timing    [][]string
	gauge     [][]string
	increment []string
	sendCh    chan *[]string
}

func NewRestapiSender() Sender {
	sender := &RestapiSender{
		sendCh: restapiCh,
	}
	return Sender(sender)
}

func (self *RestapiSender) Send(data interface{}) {
	restapiCh <- data.(*[]string)
}

func (self *RestapiSender) SetConfig(rawConfig interface{}) error {
	return nil
}
