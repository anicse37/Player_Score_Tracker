package main

import (
	"log"
	"net/http"

	server "github.com/anicse37/Player_Score_Tracker"
)

func main() {
	server := &server.PlayerServer{}
	log.Fatal(http.ListenAndServe(":8080", server))
}
