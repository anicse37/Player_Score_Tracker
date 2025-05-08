package main

import (
	"log"
	"net/http"

	functions "github.com/anicse37/Player_Score_Tracker"
	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

func main() {
	server := server.NewPlayerServer(functions.NewInMemoryStore())
	log.Fatal(http.ListenAndServe(":8080", server))
}
