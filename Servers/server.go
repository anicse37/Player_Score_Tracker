package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	models "github.com/anicse37/Player_Score_Tracker/Models"
)

const (
	JSONContentType = "application/json"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []models.Player
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

/*------------------------------------------------------------------*/
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.Store = store
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.LeagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.PlayerHandler))
	p.Handler = router
	return p
}

/*------------------------------------------------------------------*/
func (p *PlayerServer) LeagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JSONContentType)
	json.NewEncoder(w).Encode(p.Store.GetLeague())
}

func (p *PlayerServer) PlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.MethodPostFunction(w, player)
	case http.MethodGet:
		p.MethodGetFunction(w, player)
	}
}

/*------------------------------------------------------------------*/
func (p *PlayerServer) MethodGetFunction(w http.ResponseWriter, player string) {
	status := p.Store.GetPlayerScore(player)

	if status == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, status)
}

func (p *PlayerServer) MethodPostFunction(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

/*------------------------------------------------------------------*/

func (p *PlayerServer) ProcessWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

/*------------------------------------------------------------------*/
func GetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}
func PostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, (fmt.Sprintf("/players/%s", name)), nil)
	return request
}

/*------------------------------------------------------------------*/
