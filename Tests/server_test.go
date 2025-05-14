package server_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   files.League
}

func (s *StubPlayerStore) GetLeague() files.League {
	return s.league
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

/*--------------------------------------------------------*/
func TestGETPlayers(t *testing.T) {
	Store := StubPlayerStore{
		map[string]int{
			"Player-1": 20,
			"Player-2": 10,
		},
		nil,
		nil,
	}
	server1 := server.NewPlayerServer(&Store)
	t.Run("Return Player-1 Score", func(t *testing.T) {
		request := server.GetScoreRequest("Player-1")
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("Return Player-2 score", func(t *testing.T) {
		request := server.GetScoreRequest("Player-2")
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("Error 404, Not Found", func(t *testing.T) {
		request := server.GetScoreRequest("Player-3")
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusNotFound)
	})
	t.Run("It returs accepted on Post", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Player-1", nil)
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)
	})
}

/*--------------------------------------------------------*/
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server1 := server.NewPlayerServer(&store)
	t.Run("It Records wins when POST", func(t *testing.T) {
		player := "Player-1"
		request := server.PostWinRequest(player)
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("Got %d || Want %d\n", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("Got %s || Want %s\n", store.winCalls[0], player)

		}
	})
}
func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server1 := server.NewPlayerServer(&store)

	t.Run("It runs 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		var got []files.Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		AssertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("Returns league table as JSON", func(t *testing.T) {
		wantedLeague := []files.Player{
			{Name: "Player-1", Wins: 10},
			{Name: "Player-2", Wins: 20},
			{Name: "Player-3", Wins: 30},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server1 := server.NewPlayerServer(&store)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)

		if response.Result().Header.Get("content-type") != "application/json" {
			t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
		}
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
	})
}

/*----------------------------------------------------------------------------------*/
func GetLeagueFromResponse(t testing.TB, body io.Reader) (league []files.Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return
}
func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

/*--------------------------------------------------------*/
func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got Status %v || Want Status %v \n", got, want)
	}
}
func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v || Want %v \n", got, want)
	}
}
func AssertLeague(t testing.TB, got, want []files.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}
