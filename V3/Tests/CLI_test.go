package server_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/anicse37/Player_Score_Tracker/cmd"
)

var (
	DummyBlindAlerter = &cmd.SpyBlindAlerter{}
	// dummySpyAlerter   = &cmd.SpyBlindAlerter{}
	dummyPlayerStore = &StubPlayerStore{}
	// dummyStdIn       = &bytes.Buffer{}
	// dummyStdOut       = &bytes.Buffer{}
)

func TestCLI(t *testing.T) {

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &cmd.SpyBlindAlerter{}

		game := cmd.NewGame(blindAlerter, dummyPlayerStore)
		cli := cmd.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := cmd.PromptText

		if got != want {
			t.Errorf("Got %v || Want %v ", got, want)
		}

		cases := []cmd.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.Alerts) <= 1 {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				AssertScheduledAlert(t, got.String(), want.String())
			})
		}
	})
}
