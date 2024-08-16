package router

import (
	"net/http"
	"server/internal/user"
	"server/internal/ws"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouterAndStart(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	
	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/get-rooms", wsHandler.GetRooms)
	r.GET("/ws/get-clients/:roomId", wsHandler.GetClients)
	http.ListenAndServe(":80", r)
}
