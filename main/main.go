package main

import (
	"log"
	"net/http"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

func main() {
	server := server.NewPlayerServer(files.NewInMemoryStore())
	log.Fatal(http.ListenAndServe(":8080", server))
}
