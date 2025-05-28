package main

import (
	"log"
	"net/http"

	files "github.com/anicse37/Player-Score-Tracker/Files"
	servers "github.com/anicse37/Player-Score-Tracker/Servers"
)

func main() {
	File, closeFile, err := files.PlayerDataFromFiles()
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile()

	server := servers.NewPlayerServer(File)
	if err := http.ListenAndServe(":5050", server); err != nil {
		log.Fatalf("Could not listn on port :5050 %v", err)
	}
}
