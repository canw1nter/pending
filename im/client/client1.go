package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"

	"pending/im/server/handler"
)

func main() {
	dialer := &websocket.Dialer{}
	var header = http.Header{}
	header.Add("ID", "client1")
	conn, _, _ := dialer.Dial("ws://localhost:8001/online", header)
	defer conn.Close()

	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read message from server failed! err: %s\n", err.Error())
				break
			}
			var message handler.Message
			err = json.Unmarshal(p, &message)
			if err != nil {
				log.Printf("Decode message from server failed! err: %s\n", err.Error())
				return
			}

			log.Printf("Recieve message from %s: %s", message.FromUserID, message.Text)
		}
	}()

	for {
		stdReader := bufio.NewReader(os.Stdin)
		line, _, err := stdReader.ReadLine()
		if err != nil {
			log.Printf("Read user input failed! err: %s\n", err.Error())
			return
		}

		message := &handler.Message{
			ToUserID:   "client2",
			Text:       string(line),
			FromUserID: "client1",
		}
		messageData, err := json.Marshal(message)
		if err != nil {
			log.Printf("Encode message failed! err: %s\n", err.Error())
			return
		}
		err = conn.WriteMessage(websocket.BinaryMessage, messageData)
		if err != nil {
			log.Printf("Send message failed! err: %s\n", err.Error())
			return
		}
	}
}
