package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	"github.com/gorilla/websocket"
)

const (
	JSONContentType  = "application/json"
	htmlTemplatePath = "game.html"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() files.League
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler
	template *template.Template
}

/*------------------------------------------------------------------*/
func NewPlayerServer(store PlayerStore) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
	}

	p.template = tmpl
	p.Store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.LeagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.PlayerHandler))
	router.Handle("/game", http.HandlerFunc(p.GameFunc))
	router.Handle("/ws", http.HandlerFunc(p.WebSocket))

	p.Handler = router
	return p, nil
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
func (p *PlayerServer) GameFunc(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}
func (p *PlayerServer) WebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _ := upgrader.Upgrade(w, r, nil)
	_, winnerMSG, _ := conn.ReadMessage()
	p.Store.RecordWin(string(winnerMSG))
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
func GetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}
func PostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, (fmt.Sprintf("/players/%s", name)), nil)
	return request
}

/*------------------------------------------------------------------*/
