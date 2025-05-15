package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
type CLI struct {
	PlayerStore server.PlayerStore
	In          *bufio.Scanner
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
	reader := cli.readline()
	cli.PlayerStore.RecordWin(extractWinner(reader))
}
func (cli *CLI) PlayPokerWithBlindAlerter() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.Alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += 10 * time.Minute
	}
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
func NewCLIWithBlindAlterter(store server.PlayerStore, in io.Reader, alteter BlindAlerter) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
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
func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{At: duration, Amount: amount})
}

/*---------------------------------------------------------------*/
