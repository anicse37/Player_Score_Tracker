package server_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/anicse37/Player_Score_Tracker/cmd"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alearts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &cmd.SpyBlindAlerter{}
		game := cmd.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(5)

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

		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("schedules alerts on game start for7 players", func(t *testing.T) {
		blindAlerter := &cmd.SpyBlindAlerter{}
		game := cmd.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(7)
		cases := []cmd.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("itprompts the userto enter number of players and start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &cmd.GameSpy{}

		cli := cmd.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := cmd.PromptText

		if gotPrompt != wantPrompt {
			t.Errorf("Got %v || want %v", gotPrompt, wantPrompt)
		}
		if game.StartWith != 7 {
			t.Errorf("wanted to start with 7 but got %v", game.StartWith)
		}

	})
}
func TestGame_Finish(t *testing.T) {
	store := &StubPlayerStore{}
	game := cmd.NewGame(DummyBlindAlerter, store)
	winner := "Bob"

	game.Finish(winner)
	AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []cmd.ScheduledAlert, t *testing.T, blindAlerter *cmd.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			AssertScheduledAlert(t, got.String(), want.String())
		})
	}
}
