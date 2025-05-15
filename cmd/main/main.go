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
	fmt.Println("Let's Play a Game")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Problem opening %s, %v", dbFileName, err)
	}

	store, err1 := files.NewPlayerReadWriteSeeker(db)
	if err1 != nil {
		fmt.Printf("didn't expect an error but got one, %v", err1)
	}

	game := cmd.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
