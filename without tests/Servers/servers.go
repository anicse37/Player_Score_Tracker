package servers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	files "github.com/anicse37/Player-Score-Tracker/Files"
)

const (
	contentType = "application/json"
)

type PlayerFunctions interface {
	RecordWin(name string)
	GetPlayerWins(name string) int
	GetLeague() files.League
}
type PlayerServer struct {
	Store PlayerFunctions
	http.Handler
}

func NewPlayerServer(file PlayerFunctions) *PlayerServer {
	p := new(PlayerServer)
	p.Store = file
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.LeagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.PlayerHandler))

	p.Handler = router
	return p
}

func (P *PlayerServer) LeagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", contentType)
	json.NewEncoder(w).Encode(P.Store.GetLeague())
}

func (P *PlayerServer) PlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		P.MethodGetFunction(w, player)
	case http.MethodPost:
		P.MethodPostFunction(w, player)
	}
}
func (P *PlayerServer) MethodGetFunction(w http.ResponseWriter, player string) {
	store := P.Store.GetPlayerWins(player)
	if store == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, store)
}

func (P *PlayerServer) MethodPostFunction(w http.ResponseWriter, player string) {
	P.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
