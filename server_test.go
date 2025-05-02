package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/anicse37/Player_Score_Tracker"
)

func TestG(t *testing.T) {
	t.Run("Return Score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Aniket", nil)
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("Got %v || Want %v \n", got, want)
		}
	})
}
