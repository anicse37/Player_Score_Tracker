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
	store, closeFunc, err := files.PlayerSeekerFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	fmt.Println("Let's Play a Game")
	fmt.Println("Type {Name} wins to record a win")

	game := cmd.NewGame(cmd.BlindAlerterFunc(cmd.StdOutAlerter), store)
	cmd.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
