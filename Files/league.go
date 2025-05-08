package files

import (
	"encoding/json"
	"fmt"
	"io"

	models "github.com/anicse37/Player_Score_Tracker/Models"
)

func NewLeague(rdr io.Reader) ([]models.Player, error) {
	var league []models.Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("error parsing league, %v ", err)
	}
	return league, err
}
