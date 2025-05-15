package cmd

import (
	"bufio"
	"io"
	"strings"
	"time"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

type CLI struct {
	PlayerStore server.PlayerStore
	In          *bufio.Scanner
	Alerter     BlindAlerter
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type SpyBlindAlerter struct {
	Alerts []struct {
		ScheduledAt time.Duration
		Amount      int
	}
}

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

func (cli *CLI) PlayPoker() {
	reader := cli.readline()
	cli.PlayerStore.RecordWin(extractWinner(reader))
}
func (cli *CLI) PlayPokerWithBlindAlerter() {
	cli.Alerter.ScheduleAlertAt(5*time.Second, 100)
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

/*--------------------------------------------------------------*/
func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, struct {
		ScheduledAt time.Duration
		Amount      int
	}{duration, amount})
}
