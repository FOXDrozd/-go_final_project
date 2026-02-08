package server

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"main.go/pkg/api"
)

func Run() {
	godotenv.Load()
	api.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "7540"
	}

	log.Println("server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
