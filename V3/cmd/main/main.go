package main

import (
	"log"
	"net/http"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	store, closeFunc, err := files.PlayerSeekerFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	server := server.NewPlayerServer(store)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
