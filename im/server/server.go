package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"pending/im/server/handler"
)

func main() {
	server := gin.Default()

	handler.RegisterRoute(server)

	err := server.Run(":8001")
	if err != nil {
		log.Fatalf("Start server failed! err: %s\n", err.Error())
		return
	}
}
