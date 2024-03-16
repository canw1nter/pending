package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"

	"pending/im/server/handler"
)

var (
	mutex       sync.Mutex
	connections = map[string]*websocket.Conn{}
)

func main() {
	for i := 0; i < 10000; i++ {
		go func() {
			name := "client" + strconv.Itoa(i)
			dialer := &websocket.Dialer{}
			var header = http.Header{}
			header.Add("ID", name)
			conn, _, err := dialer.Dial("ws://localhost:8001/online", header)
			if err != nil {
				log.Printf("%s connect to server failed! err: %s\n", name, err.Error())
			}
			defer conn.Close()
			mutex.Lock()
			connections[name] = conn
			mutex.Unlock()

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

				log.Printf("%s Recieve message from %s: %s\n", message.ToUserID, message.FromUserID, message.Text)
			}
		}()
	}

	for {
		stdReader := bufio.NewReader(os.Stdin)
		line, _, err := stdReader.ReadLine()
		if err != nil {
			log.Printf("Read user input failed! err: %s\n", err.Error())
			return
		}
		info := strings.Split(string(line), " ")

		message := &handler.Message{
			ToUserID:   info[0],
			Text:       info[2],
			FromUserID: info[1],
		}
		messageData, err := json.Marshal(message)
		if err != nil {
			log.Printf("Encode message failed! err: %s\n", err.Error())
			return
		}
		mutex.Lock()
		conn := connections[info[1]]
		mutex.Unlock()
		err = conn.WriteMessage(websocket.BinaryMessage, messageData)
		if err != nil {
			log.Printf("Send message failed! err: %s\n", err.Error())
			return
		}
	}
}
