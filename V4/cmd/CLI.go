package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	PromptText = "Please enter the number of players: "
)

type CLI struct {
	In   *bufio.Scanner
	out  io.Writer
	Game Game
}
type ScheduledAlert struct {
	At     time.Duration
	Amount int
}
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

type BlindAlerter interface {
	ScheduledAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduledAlertAt(duration time.Duration, amount int) {
	a(duration, amount)

}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}

/*---------------------------------------------------------------*/

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PromptText)

	numberOfPlayerInput := cli.readline()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayerInput, "\n"))
	if err != nil {
		return
	}

	cli.Game.Start(numberOfPlayers)

	winnerInput := cli.readline()
	winner := extractWinner(winnerInput)

	cli.Game.Finish(winner)
}

/*---------------------------------------------------------------*/
func extractWinner(name string) string {
	return strings.Replace(name, " wins", "", 1)
}
func (cli *CLI) readline() string {
	cli.In.Scan()
	return cli.In.Text()
}
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		In:   bufio.NewScanner(in),
		out:  out,
		Game: game,
	}
}
func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

/*---------------------------------------------------------------*/
func (s *SpyBlindAlerter) ScheduledAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{At: duration, Amount: amount})
}

/*---------------------------------------------------------------*/
