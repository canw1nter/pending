package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	mutex    sync.Mutex
	userConn = map[string]*websocket.Conn{}
	upgrader = websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
	}
)

type CreateConnectionRequest struct {
	UserID string `json:"user_id"`
}

type Message struct {
	ToUserID   string `json:"to_user_id"`
	Text       string `json:"text"`
	FromUserID string `json:"from_user_id"`
}

func CreateConnectionHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("ID")
	if userID == "" {
		log.Printf("Identify failed")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Can't create connection to %s! err: %s\n", userID, err.Error())
		return
	}
	defer conn.Close()
	mutex.Lock()
	userConn[userID] = conn
	mutex.Unlock()
	log.Printf("Connect to %s successfully!", userID)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Some error when read message from %s's connection! err: %s\n", userID, err.Error())
			if websocket.IsCloseError(err) {
				mutex.Lock()
				delete(userConn, userID)
			}
			break
		}
		var message Message
		if err = json.Unmarshal(p, &message); err != nil {
			log.Printf("Some error when decode message from %s! err: %s\n", userID, err.Error())
			continue
		}

		if sendConn, ok := userConn[message.ToUserID]; ok {
			if err = sendConn.WriteMessage(websocket.BinaryMessage, p); err != nil {
				log.Printf("User %s send message to %s failed! err: %s\n",
					message.FromUserID, message.ToUserID, err.Error())
				return
			}
		}
	}
}
