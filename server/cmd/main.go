package main

import (
	"log"
	"os"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Couldn't connect to DB: %s", err)
	}

	if os.Getenv("secretkey") == "" {
		log.Fatalf("No secretkey environment variable")
	}

	userRepository := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRepository)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouterAndStart(userHandler, wsHandler)
	// router.Start("localhost:8080")
}
