package server_test

import (
	"strings"
	"testing"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	models "github.com/anicse37/Player_Score_Tracker/Models"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		database := strings.NewReader(`
		[
		{"Name": "Player-1","Wins": 10},
		{"Name": "Player-2","Wins": 20},
		{"Name": "Player-3","Wins": 30}
		]`)

		store := files.FileSystemPlayerDatabase{Database: database}

		got := store.GetLeague()
		want := []models.Player{
			{Name: "Player-1", Wins: 10},
			{Name: "Player-2", Wins: 20},
			{Name: "Player-3", Wins: 30},
		}

		AssertLeague(t, got, want)
	})
	t.Run("League Twice from a reader", func(t *testing.T) {
		database := strings.NewReader(`
		[
		{"Name": "Player-1","Wins": 10},
		{"Name": "Player-2","Wins": 20},
		{"Name": "Player-3","Wins": 30}
		]`)

		store := files.FileSystemPlayerReadSeeker{Database: database}

		got := store.GetLeague()
		want := []models.Player{
			{Name: "Player-1", Wins: 10},
			{Name: "Player-2", Wins: 20},
			{Name: "Player-3", Wins: 30},
		}
		AssertLeague(t, got, want)
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
	t.Run("Get Player Score", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name":"Player-1","Wins":10},
		{"Name":"Player-2","Wins":20},
		{"Name":"Player-3","Wins":30}
		]`)
		store := files.FileSystemPlayerReadSeeker{Database: database}

		got := store.GetPlayerScore("Player-2")
		want := 20

		AssertScoreEquals(t, got, want)
	})
}
