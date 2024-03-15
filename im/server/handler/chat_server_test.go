package handler

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/gorilla/websocket"
)

func TestServerResourceConsume(t *testing.T) {
	for i := 0; i < 50000; i++ {
		go func() {
			name := "client" + strconv.Itoa(i)
			dialer := &websocket.Dialer{}
			var header = http.Header{}
			header.Add("ID", name)
			conn, _, _ := dialer.Dial("ws://localhost:8001/online", header)
			defer conn.Close()

			var forever chan struct{}
			<-forever
		}()
	}

	var forever chan struct{}
	<-forever
}
