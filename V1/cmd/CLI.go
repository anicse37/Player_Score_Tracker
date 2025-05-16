package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          *bufio.Scanner
}
type ScheduledAlert struct {
	At     time.Duration
	Amount int
}
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

/*---------------------------------------------------------------*/

func (cli *CLI) PlayPoker() {
	reader := cli.readline()
	cli.PlayerStore.RecordWin(extractWinner(reader))
}

/*---------------------------------------------------------------*/
func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

/*---------------------------------------------------------------*/
func NewCLI(store server.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
	}
}
func extractWinner(name string) string {
	return strings.Replace(name, " wins", "", 1)
}
func (cli *CLI) readline() string {
	cli.In.Scan()
	return cli.In.Text()
}

/*---------------------------------------------------------------*/
func (s *SpyBlindAlerter) ScheduledAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{At: duration, Amount: amount})
}

/*---------------------------------------------------------------*/
