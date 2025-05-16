package server_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/anicse37/Player_Score_Tracker/cmd"
)

var DummySpyAlerter = &cmd.SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("Record Ani win from user input", func(t *testing.T) {

		in := strings.NewReader("Ani wins\n")
		playerStore := &StubPlayerStore{}

		cli := cmd.NewCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Ani")
	})
	t.Run("Record Aniket win from user input", func(t *testing.T) {
		in := strings.NewReader("Aniket wins\n")
		playerStore := &StubPlayerStore{}

		cli := cmd.NewCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Aniket")
	})
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Ani wins\n")
		playerStore := &StubPlayerStore{}

		cli := cmd.NewCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Ani")
	})
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Ani wins\n")
		playerStore := &StubPlayerStore{}
		blindAlerter := &cmd.SpyBlindAlerter{}

		cli := cmd.NewCLIWithBlindAlterter(playerStore, in, blindAlerter)
		cli.PlayPokerWithBlindAlerter()

		cases := []cmd.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want.String()), func(t *testing.T) {

				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
				got := blindAlerter.Alerts[i]

				AssertScheduledAlert(t, got.String(), want.String())
			})
		}
	})
}
