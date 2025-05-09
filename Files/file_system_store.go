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
	f.Database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.Database)
	return league
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

/*---------------------------------------------------------------*/
func (f *PlayerReadWriteSeeker) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins

		}
	}
	return wins
}

/*-------------RecordWin---------------*/
func (f *PlayerReadWriteSeeker) RecordWin(name string) {
	league := f.GetLeague()

	for i, players := range league {
		if players.Name == name {
			league[i].Wins++
		}
	}
	f.Database.Seek(0, io.SeekStart)
	json.NewEncoder(f.Database).Encode(league)
}
