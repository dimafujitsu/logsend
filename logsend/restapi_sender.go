package logsend

import (
	"github.com/jmcvetta/napping"
	/*
		"flag"
		"strings"
		"time"
	*/)

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

			type Foo struct {
				Bar string
			}
			type Spam struct {
				Eggs int
			}
			payload := Foo{
				Bar: "baz",
			}
			result := Spam{}
			url := "http://foo.com/bar"
			resp, err := napping.Post(url, &payload, &result, nil)
			if err != nil {
				panic(err)
			}
			if resp.Status() == 200 {
				println(result.Eggs)
			}
		}
	}()
	return
}

func (sender *RestapiSender) Name() string {
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

func (sender *RestapiSender) Send(data interface{}) {
	restapiCh <- data.(*[]string)
}

func (sender *RestapiSender) SetConfig(rawConfig interface{}) error {
	return nil
}
