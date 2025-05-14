package main

import (
	"fmt"
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

	store, err1 := files.NewPlayerReadWriteSeeker(db)
	if err1 != nil {
		fmt.Printf("didn't expect an error but got one, %v", err1)
	}

	server := server.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not listen on port 8080 %v", err)
	}
}
