package files

import (
	"encoding/json"
	"io"
	"os"
)

type Player struct {
	Name string
	Wins int
}
type League []Player

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Wins: 1,
	}
}

type PlayerDatabase struct {
	jsonFile *json.Encoder
	league   League
}

type Tape struct {
	File *os.File
}

func (t *Tape) Write(data []byte) (n int, err error) {
	t.File.Truncate(0)
	t.File.Seek(0, io.SeekStart)
	return t.File.Write(data)
}
