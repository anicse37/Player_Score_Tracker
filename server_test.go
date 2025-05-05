package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/anicse37/Player_Score_Tracker"
)

func TestGETPlayers(t *testing.T) {
	Store := StubPlayerStore{
		map[string]int{
			"Player-1": 20,
			"Player-2": 10,
		},
	}
	server1 := &server.PlayerServer{&Store}
	t.Run("Return Player-1 Score", func(t *testing.T) {
		request := GetScoreRequest("Player-1")
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		asstetStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("Return Player-2 score", func(t *testing.T) {
		request := GetScoreRequest("Player-2")
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		asstetStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("Error 404, Not Found", func(t *testing.T) {
		request := GetScoreRequest("Player-3")
		response := httptest.NewRecorder()

		server1.ServeHTTP(response, request)

		asstetStatus(t, response.Code, http.StatusNotFound)

	})
}

/*--------------------------------------------------------*/

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func GetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func asstetStatus(t *testing.T, got, want int) {
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
