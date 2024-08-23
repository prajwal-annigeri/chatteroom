package main

import (
	"flag"
	"log"
	"os"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)


var connString = flag.String("conn-string", "", "Postgres connection string")

func parseFlags() {
	flag.Parse()

	if *connString == "" {
		log.Fatalf("Need conn-string")
	}
}

func main() {
	parseFlags()

	if os.Getenv("secretkey") == "" {
		log.Fatalf("Need secretkey environment variable")
	}

	dbConn, err := db.NewDatabase(*connString)
	if err != nil {
		log.Fatalf("Couldn't connect to DB: %s", err)
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
