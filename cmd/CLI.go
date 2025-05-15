package cmd

import (
	"bufio"
	"io"
	"strings"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          *bufio.Scanner
}

func NewCLI(store server.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	reader := cli.readline()
	cli.PlayerStore.RecordWin(extractWinner(reader))
}

func extractWinner(name string) string {
	return strings.Replace(name, " wins", "", 1)
}

func (cli *CLI) readline() string {
	cli.In.Scan()
	return cli.In.Text()
}
