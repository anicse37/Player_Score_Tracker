package files

import (
	"encoding/json"
	"io"
)

type PlayerDatabase struct {
	Database io.Reader
}
type PlayerReadSeeker struct {
	Database io.ReadSeeker
}
type PlayerReadWriteSeeker struct {
	Database io.ReadWriteSeeker
	league   League
}

/*---------------------------------------------------------------*/
func (f *PlayerDatabase) GetLeague() []Player {
	league, _ := NewLeague(f.Database)
	return league
}
func (f *PlayerReadSeeker) GetLeague() []Player {
	f.Database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.Database)
	return league
}
func (f *PlayerReadWriteSeeker) GetLeague() []Player {
	return f.league
}

/*---------------------------------------------------------------*/
func (f *PlayerReadSeeker) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins

		}
	}
	return wins
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

	f.Database.Seek(0, io.SeekStart)
	json.NewEncoder(f.Database).Encode(f.league)
}

/*------------------------------------------------------------*/
func NewPlayerReadWriteSeeker(database io.ReadWriteSeeker) *PlayerReadWriteSeeker {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &PlayerReadWriteSeeker{
		Database: database,
		league:   league,
	}
}
