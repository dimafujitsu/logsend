package logsend

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestRest(t *testing.T) {
	var result string
	sender := &RestapiSender{}
	result = sender.Name()
	if result != "restapi" {
		t.Error("Failure!!")
	}
}

func TestKube(t *testing.T) {
	orig := os.Stdout
	r, w, _ := os.Pipe()

	c := make(chan string, 1)
	go func() {
		c <- "test"
		go KubernetesReader()
	}()
	select {
	case <-time.After(time.Second * 2):
		os.Exit(0)
	}
	var buf bytes.Buffer
	io.Copy(&buf, r)
	w.Close()
	os.Stdout = orig
	fmt.Printf("Res: %+v\n", buf.String())
	if buf.String() != "Some message" {
		t.Error("Failure!")
	}
}
