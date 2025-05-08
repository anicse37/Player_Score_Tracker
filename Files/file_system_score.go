package files

import (
	"io"

	models "github.com/anicse37/Player_Score_Tracker/Models"
)

type FileSystemPlayerDatabase struct {
	Database io.Reader
}

func (f *FileSystemPlayerDatabase) GetLeague() []models.Player {
	league, _ := NewLeague(f.Database)
	return league
}
