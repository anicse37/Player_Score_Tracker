package cmd

import (
	"bufio"
	"io"
	"strings"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.In)
	reader.Scan()
	cli.PlayerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(name string) string {
	return strings.Replace(name, " wins", "", 1)
}
