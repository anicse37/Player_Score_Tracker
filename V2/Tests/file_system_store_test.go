package server_test

import (
	"os"
	"testing"

	files "github.com/anicse37/Player_Score_Tracker/Files"
)

/*------------------------------------------------------------------*/
func TestReadUsingFiles(t *testing.T) {
	t.Run("Read Write Seeker", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
		{"Name":"Player-1","Wins":10},
		{"Name":"Player-2","Wins":20},
		{"Name":"Player-3","Wins":30}
		]`)
		defer cleanDatabase()

		store, err := files.NewPlayerReadWriteSeeker(database)
		AssertNoError(t, err)
		got := store.GetLeague()
		want := files.League{
			{Name: "Player-3", Wins: 30},
			{Name: "Player-2", Wins: 20},
			{Name: "Player-1", Wins: 10},
		}
		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
	t.Run("Read Write Seeker", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
		{"Name":"Player-3","Wins":30},
		{"Name":"Player-2","Wins":20},
		{"Name":"Player-1","Wins":10}
		]`)
		defer cleanDatabase()

		store, err := files.NewPlayerReadWriteSeeker(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Player-1")
		want := 10
		AssertScoreEquals(t, got, want)
	})

}

func TestRecordWin(t *testing.T) {
	t.Run("Old Player", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name":"Player-3","Wins":30},
			{"Name":"Player-2","Wins":20},
			{"Name":"Player-1","Wins":10}
		]`)
		defer cleanDatabase()

		store, err := files.NewPlayerReadWriteSeeker(database)
		AssertNoError(t, err)

		store.RecordWin("Player-2")

		got := store.GetPlayerScore("Player-2")
		want := 21
		AssertScoreEquals(t, got, want)
	})
	t.Run("New Player", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name":"Player-3","Wins":30},
			{"Name":"Player-2","Wins":20},
			{"Name":"Player-1","Wins":10}
		]`)
		defer cleanDatabase()

		store, err := files.NewPlayerReadWriteSeeker(database)

		AssertNoError(t, err)

		store.RecordWin("Player-4")

		got := store.GetPlayerScore("Player-4")
		want := 1
		AssertScoreEquals(t, got, want)
	})
}
func TestFunctions(t *testing.T) {
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := files.NewPlayerReadWriteSeeker(database)

		AssertNoError(t, err)

	})
	t.Run("league, sorted", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
		{"Name": "Aniket","Wins" : 33},
		{"Name": "Chris","Wins" : 10}]`)
		defer cleanDatabase()

		store, err := files.NewPlayerReadWriteSeeker(database)
		AssertNoError(t, err)

		got := store.GetLeague()
		want := files.League{
			{Name: "Aniket", Wins: 33},
			{Name: "Chris", Wins: 10},
		}
		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}

/*------------------------------------------------------------------*/
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
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
func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
