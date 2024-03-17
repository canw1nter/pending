package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pending/im/server/middleware"
)

func RegisterRoute(engine *gin.Engine) {
	engine.Handle(http.MethodPost, "/login", loginHandler)

	engine.Use(middleware.VerifyToken())

	engine.Handle(http.MethodGet, "/online", createWebSocketConnectionHandler)
}
