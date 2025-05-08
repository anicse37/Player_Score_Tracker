package functions

import models "github.com/anicse37/Player_Score_Tracker/Models"

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{map[string]int{}}
}

type InMemoryStore struct {
	score map[string]int
}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return i.score[name]
}
func (i *InMemoryStore) RecordWin(name string) {
	i.score[name]++
}

func (i *InMemoryStore) GetLeague() []models.Player {
	var league []models.Player
	for Name, Wins := range i.score {
		league = append(league, models.Player{Name: Name, Wins: Wins})
	}
	return league
}
