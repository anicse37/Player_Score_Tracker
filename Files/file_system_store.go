package files

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type PlayerReadWriteSeeker struct {
	Database *json.Encoder
	league   League
}

/*---------------------------------------------------------------*/

func (f *PlayerReadWriteSeeker) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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
func NewPlayerReadWriteSeeker(file *os.File) (*PlayerReadWriteSeeker, error) {
	err := InitialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initilising player db file %v", err)

	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &PlayerReadWriteSeeker{
			Database: json.NewEncoder(&Tape{File: file}),
			league:   league,
		},
		nil
}
func InitialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}
