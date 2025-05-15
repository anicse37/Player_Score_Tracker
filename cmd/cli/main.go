package main

import (
	"fmt"
	"log"
	"os"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	"github.com/anicse37/Player_Score_Tracker/cmd"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	store, closeFunc, err := files.PlayerReadWriteSeekerFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	fmt.Println("Let's Play a Game")
	fmt.Println("Type {Name} wins to record a win")

	cmd.NewCLI(store, os.Stdin).PlayPoker()
}
