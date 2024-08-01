package router

import (
	"net/http"
	"server/internal/user"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouterAndStart(userHandler *user.Handler) {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	http.ListenAndServe(":80", r)
}
