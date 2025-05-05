package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
	}
	switch r.Method {
	case http.MethodPost:
		p.MethodPostFunction(w, r)
	case http.MethodGet:
		p.MethodGetFunction(w, r)
	}
}
func (p *PlayerServer) MethodGetFunction(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	status := p.Store.GetPlayerScore(player)

	if status == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, status)
}

func (p *PlayerServer) MethodPostFunction(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// server.go
func (p *PlayerServer) ProcessWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
