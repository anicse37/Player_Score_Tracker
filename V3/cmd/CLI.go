package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

const (
	PromptText = "Please enter the number of players: "
)

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

type CLI struct {
	PlayerStore server.PlayerStore
	In          *bufio.Scanner
	out         io.Writer
	Alerter     BlindAlerter
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
	fmt.Fprint(cli.out, PromptText)

	number_of_players, _ := strconv.Atoi(cli.readline())

	cli.ScheduleBlindAlerts(number_of_players)

	reader := cli.readline()
	cli.PlayerStore.RecordWin(extractWinner(reader))
}

/*---------------------------------------------------------------*/
func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

/*---------------------------------------------------------------*/
func NewCLI(store server.PlayerStore, in io.Reader, out io.Writer, alteter BlindAlerter) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
		out:         out,
		Alerter:     alteter,
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

func (cli *CLI) ScheduleBlindAlerts(number int) {
	blindIncrement := time.Duration(5+number) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.Alerter.ScheduledAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}
