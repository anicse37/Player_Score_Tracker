package server_test

import (
	"io"
	"os"
	"strings"
	"testing"

	files "github.com/anicse37/Player_Score_Tracker/Files"
)

func TestStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		database := strings.NewReader(`
		[
			{"Name": "Player-1","Wins": 10},
			{"Name": "Player-2","Wins": 20},
			{"Name": "Player-3","Wins": 30}
		]`)

		store := files.PlayerDatabase{Database: database}

		got := store.GetLeague()
		want := []files.Player{
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

		store := files.PlayerReadSeeker{Database: database}

		got := store.GetLeague()
		want := []files.Player{
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
		store := files.PlayerReadSeeker{Database: database}

		got := store.GetPlayerScore("Player-2")
		want := 20

		AssertScoreEquals(t, got, want)
	})
}

/*------------------------------------------------------------------*/
func TestReadUsingFiles(t *testing.T) {
	t.Run("Read Write Seeker", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
		{"Name":"Player-1","Wins":10},
		{"Name":"Player-2","Wins":20},
		{"Name":"Player-3","Wins":30}
		]`)
		defer cleanDatabase()

		store := files.PlayerReadWriteSeeker{Database: database}

		got := store.GetLeague()
		want := []files.Player{
			{Name: "Player-1", Wins: 10},
			{Name: "Player-2", Wins: 20},
			{Name: "Player-3", Wins: 30},
		}
		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
	t.Run("Read Write Seeker", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
		{"Name":"Player-1","Wins":10},
		{"Name":"Player-2","Wins":20},
		{"Name":"Player-3","Wins":30}
		]`)
		defer cleanDatabase()

		store := files.PlayerReadWriteSeeker{Database: database}

		got := store.GetPlayerScore("Player-1")
		want := 10
		AssertScoreEquals(t, got, want)
	})

}

/*------------------------------------------------------------------*/
func CreateTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tempfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("Could not create tmp file %v", err)
	}

	tempfile.Write([]byte(initialData))

	removeFile := func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}

	return tempfile, removeFile
}

/*----------------------------------------------------------------*/
func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}
