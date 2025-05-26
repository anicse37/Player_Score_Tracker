package main

import (
	"bufio"
	"log"
	"os"

	cli "github.com/anicse37/Player-Score-Tracker/CLI"
	files "github.com/anicse37/Player-Score-Tracker/Files"
)

const (
	JSONfileName = "gameDB.json"
)

func main() {
	store, closeFunc, err := files.GetPlaterDataFromFile(JSONfileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	reader := bufio.NewReader(os.Stdin)
	cli.NewCLI(store, reader, os.Stdin, cli.BlindAlerterFunc(cli.StdOutAlerter)).PlayPoker()
}
