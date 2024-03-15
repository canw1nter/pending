package test

import (
	"net/http"
	"testing"

	"github.com/gorilla/websocket"

	"pending/websocket/demo"
)

func TestSimpleDemoServer(t *testing.T) {
	http.HandleFunc("/server", demo.SimpleServerHandler)
	_ = http.ListenAndServe(":8000", nil)
}

func TestSimpleDemoClient(t *testing.T) {
	dialer := &websocket.Dialer{}
	conn, _, _ := dialer.Dial("ws://localhost:8000/server", nil)
	defer conn.Close()

	//line := []byte("又是一条消息")
	//_ = conn.WriteMessage(websocket.TextMessage, line)
	//_, p, _ := conn.ReadMessage()
	//fmt.Println("Receive message from server: " + string(p))
	//for {
	//	time.Sleep(5 * time.Second)
	//
	//}
	var forever chan struct{}
	<-forever
}
