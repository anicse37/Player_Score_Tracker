package server_test

import (
	"io"
	"testing"

	files "github.com/anicse37/Player_Score_Tracker/Files"
)

func TestTape_Write(t *testing.T) {
	file, clean := CreateTempFile(t, "12345")
	defer clean()

	tape := &files.Tape{File: file}
	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("Got %v || Want %v ", got, want)
	}

}
