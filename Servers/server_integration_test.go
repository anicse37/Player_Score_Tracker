package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	functions "github.com/anicse37/Player_Score_Tracker"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := functions.NewInMemoryStore()
	server1 := NewPlayerServer(store)
	player := "Pepper"

	server1.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))
	server1.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))
	server1.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))

	response := httptest.NewRecorder()
	server1.ServeHTTP(response, GetScoreRequest(player))
	AssertStatus(t, response.Code, http.StatusOK)

	AssertResponseBody(t, response.Body.String(), "3")
}
func AssertStatus(t *testing.T, got, want int) {
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
