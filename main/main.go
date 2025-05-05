package main

import (
	"log"
	"net/http"

	server "github.com/anicse37/Player_Score_Tracker"
)

type InMemoryStore struct {
	score map[string]int
}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return i.score[name]
}
func (i *InMemoryStore) RecordWin(name string) {}

func main() {
	Score := InMemoryStore{
		map[string]int{
			"Player-1": 20,
			"Player-2": 10,
		},
	}
	server := &server.PlayerServer{Store: &Score}
	log.Fatal(http.ListenAndServe(":8080", server))
}
