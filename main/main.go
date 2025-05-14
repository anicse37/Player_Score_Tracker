package main

import (
	"log"
	"net/http"
	"os"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Problem opening %s, %v", dbFileName, err)
	}
	store := &files.PlayerReadWriteSeeker{Database: db}
	server := server.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not listen on port 8080 %v", err)
	}
}
