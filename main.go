package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"main.go/pkg/db"
	"main.go/pkg/server"
)

func main() {
	godotenv.Load()

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	if err := db.Init(dbFile); err != nil {
		log.Fatal(err)
	}

	defer db.DB.Close()

	webDir := "web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	server.Run()

	port := os.Getenv("POST")

	if port == "" {
		port = "7540"
	}

	http.ListenAndServe(":"+port, nil)
}
