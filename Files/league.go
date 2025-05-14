package files

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("error parsing league, %v ", err)
	}
	return league, err
}

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}
