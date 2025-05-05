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

/*--------------------------------------------------------*/
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
	}
	switch r.Method {
	case http.MethodPost:
		p.MethodPostFunction(w, player)
	case http.MethodGet:
		p.MethodGetFunction(w, player)
	}
}
func (p *PlayerServer) MethodGetFunction(w http.ResponseWriter, player string) {

	status := p.Store.GetPlayerScore(player)

	if status == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, status)
}

/*--------------------------------------------------------*/

func (p *PlayerServer) MethodPostFunction(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) ProcessWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

/*--------------------------------------------------------*/
func GetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}
func PostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, (fmt.Sprintf("/players/%s", name)), nil)
	return request
}

/*--------------------------------------------------------*/
