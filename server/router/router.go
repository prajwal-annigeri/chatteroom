package router

import (
	"net/http"
	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouterAndStart(userHandler *user.Handler, wsHandler *ws.Handler) {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	
	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/get-rooms", wsHandler.GetRooms)
	r.GET("/ws/get-clients/:roomId", wsHandler.GetClients)
	http.ListenAndServe(":80", r)
}
