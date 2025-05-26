package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	files "github.com/anicse37/Player-Score-Tracker/Files"
)

const (
	PromptText = "Please enter the number of players: "
)

type CLI struct {
	PlayerStore files.PlayerStore
	In          *bufio.Scanner
	Out         io.Writer
	Alerter     BlindAlerter
}
type SpyBlindAlerter struct {
	Alerts []files.ScheduledAlert
}

type BlindAlerter interface {
	ScheduledAlertAt(duration time.Duration, amount int)
}
type BlindAlerterFunc func(duraton time.Duration, amount int)

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d.\n", amount)
	})
}
func NewCLI(store files.PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		PlayerStore: store,
		In:          bufio.NewScanner(in),
		Out:         out,
		Alerter:     alerter,
	}
}
func (a BlindAlerterFunc) ScheduledAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}
func (s *SpyBlindAlerter) ScheduledAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, files.ScheduledAlert{At: duration, Anount: amount})
}
func (cli *CLI) ScheduleBlindAlerts(numberOfPlayer int) {
	blinfIncrement := time.Duration(5+numberOfPlayer) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.Alerter.ScheduledAlertAt(blindTime, blind)
		blindTime += blinfIncrement
	}
}

/*----------------------------------------------------------*/
func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}
func extractWinner(name string) string {
	return strings.Replace(name, " wins", "", 1)
}
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.Out, PromptText)

	numberOfPlayer, _ := strconv.Atoi(cli.readLine())

	cli.ScheduleBlindAlerts(numberOfPlayer)
	reader := cli.readLine()
	cli.PlayerStore.RecordWin(extractWinner(reader))
}
