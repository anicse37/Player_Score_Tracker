package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.Store.GetPlayerScore(player))
}

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")
// 	fmt.Fprint(w, GetPlayerScore(player))
// }

func GetPlayerScore(player string) string {

	if player == "Player-1" {
		return "20"
	}

	if player == "Player-2" {
		return "10"
	}
	return ""
}
