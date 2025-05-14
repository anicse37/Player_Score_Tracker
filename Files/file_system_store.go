package files

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type PlayerReadWriteSeeker struct {
	Database *json.Encoder
	league   League
}

/*---------------------------------------------------------------*/

func (f *PlayerReadWriteSeeker) GetLeague() []Player {
	return f.league
}

func (f *PlayerReadWriteSeeker) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

/*-------------RecordWin---------------*/
func (f *PlayerReadWriteSeeker) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}
	f.Database.Encode(f.league)
}

/*------------------------------------------------------------*/
func NewPlayerReadWriteSeeker(database *os.File) (*PlayerReadWriteSeeker, error) {
	database.Seek(0, io.SeekStart)
	league, err := NewLeague(database)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", database.Name(), err)
	}

	return &PlayerReadWriteSeeker{
			Database: json.NewEncoder(&Tape{File: database}),
			league:   league,
		},
		nil
}
