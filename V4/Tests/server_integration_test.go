package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	files "github.com/anicse37/Player_Score_Tracker/Files"
	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	database, cleanDatabase := CreateTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := files.NewPlayerSeeker(database)

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
	server1, _ := server.NewPlayerServer(store)
	player := "Player-2"

	server1.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))
	server1.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))
	server1.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))

	response := httptest.NewRecorder()
	server1.ServeHTTP(response, server.GetScoreRequest(player))
	AssertStatus(t, response.Code, http.StatusOK)

	AssertResponseBody(t, response.Body.String(), "3")
}
