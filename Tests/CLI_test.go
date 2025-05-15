package server_test

import (
	"strings"
	"testing"

	"github.com/anicse37/Player_Score_Tracker/cmd"
)

func TestCLI(t *testing.T) {
	t.Run("Record Ani win from user input", func(t *testing.T) {

		in := strings.NewReader("Ani wins\n")
		playerStore := &StubPlayerStore{}

		cli := &cmd.CLI{PlayerStore: playerStore, In: in}
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Ani")
	})
	t.Run("Record Aniket win from user input", func(t *testing.T) {
		in := strings.NewReader("Aniket wins\n")
		playerStore := &StubPlayerStore{}

		cli := &cmd.CLI{PlayerStore: playerStore, In: in}
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Aniket")
	})
}
