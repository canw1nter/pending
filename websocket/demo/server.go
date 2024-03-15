package demo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SimpleServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start upgrade http connection")
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println("Get a message from connection")
		if err != nil {
			break
		}
		fmt.Printf("Get message from client: %s\n", string(p))
		_ = conn.WriteMessage(messageType, p)
	}
}
