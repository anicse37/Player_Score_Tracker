package files

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type PlayerSeeker struct {
	Database *json.Encoder
	league   League
}

func (f *PlayerSeeker) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
func (f *PlayerSeeker) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}
func (f *PlayerSeeker) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}
	f.Database.Encode(f.league)
}

/*---------------------------------------------------------------------*/
func PlayerSeekerFromFile(path string) (*PlayerSeeker, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}
	closeFunc := func() {
		db.Close()
	}
	store, err := NewPlayerSeeker(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store, closeFunc, nil
}
func NewPlayerSeeker(file *os.File) (*PlayerSeeker, error) {
	err := InitialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initilising player db file %v", err)

	}
	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &PlayerSeeker{
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
