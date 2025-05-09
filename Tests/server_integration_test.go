package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	functions "github.com/anicse37/Player_Score_Tracker"
	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := functions.NewInMemoryStore()
	server1 := server.NewPlayerServer(store)
	player := "Pepper"

	server1.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))
	server1.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))
	server1.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))

	response := httptest.NewRecorder()
	server1.ServeHTTP(response, server.GetScoreRequest(player))
	AssertStatus(t, response.Code, http.StatusOK)

	AssertResponseBody(t, response.Body.String(), "3")
}
