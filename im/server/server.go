package main

import (
	"net/http"

	"pending/im/server/handler"
)

func main() {
	http.HandleFunc("/online", handler.CreateConnectionHandler)
	_ = http.ListenAndServe(":8001", nil)
}
