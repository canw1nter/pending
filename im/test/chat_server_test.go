package test

import (
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/gorilla/websocket"
)

func TestServerResourceConsume(t *testing.T) {
	for i := 0; i < 10000; i++ {
		i := i
		go func() {
			name := "client" + strconv.Itoa(i)
			dialer := &websocket.Dialer{}
			var header = http.Header{}
			header.Add("ID", name)
			conn, _, err := dialer.Dial("ws://localhost:8001/online", header)
			if err != nil {
				log.Printf("%d: %s connect to server failed! err: %s\n", i, name, err.Error())
			}
			defer conn.Close()

			var forever chan struct{}
			<-forever
		}()
	}

	var forever chan struct{}
	<-forever
}
