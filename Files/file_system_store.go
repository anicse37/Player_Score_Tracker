package files

import (
	"io"
)

type FileSystemPlayerDatabase struct {
	Database io.Reader
}
type FileSystemPlayerReadSeeker struct {
	Database io.ReadSeeker
}

/*---------------------------------------------------------------*/
func (f *FileSystemPlayerDatabase) GetLeague() []Player {
	league, _ := NewLeague(f.Database)
	return league
}
func (f *FileSystemPlayerReadSeeker) GetLeague() []Player {
	f.Database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.Database)
	return league
}

/*---------------------------------------------------------------*/
func (f *FileSystemPlayerReadSeeker) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins

		}
	}
	return wins
}

/*---------------------------------------------------------------*/
