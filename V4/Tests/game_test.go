package server_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
	"github.com/anicse37/Player_Score_Tracker/cmd"
	"github.com/gorilla/websocket"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alearts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &cmd.SpyBlindAlerter{}
		game := cmd.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []cmd.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("schedules alerts on game start for7 players", func(t *testing.T) {
		blindAlerter := &cmd.SpyBlindAlerter{}
		game := cmd.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(7)
		cases := []cmd.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("itprompts the userto enter number of players and start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &cmd.GameSpy{}

		cli := cmd.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := cmd.PromptText

		if gotPrompt != wantPrompt {
			t.Errorf("Got %v || want %v", gotPrompt, wantPrompt)
		}
		if game.StartCalledWith != 7 {
			t.Errorf("wanted to start with 7 but got %v", game.StartCalledWith)
		}
	})
	t.Run("itprints an error when a non numeric value i entered", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &cmd.GameSpy{}

		cli := cmd.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("Game shot have started")
		}
	})
}
func TestGame_Finish(t *testing.T) {
	t.Run("GET /gamereturns 500 ", func(t *testing.T) {
		server, _ := server.NewPlayerServer(&StubPlayerStore{})

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusInternalServerError)
	})

	t.Run("get a mesage over a websocket it is a winener of a game", func(t *testing.T) {
		store := &StubPlayerStore{}
		winner := "Ruth"

		temp, _ := server.NewPlayerServer(store)
		server := httptest.NewServer(temp)
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
		}
		defer ws.Close()

		if err := ws.WriteMessage(websocket.TextMessage, []byte(winner)); err != nil {
			t.Fatalf("could not sed message over ws connection %v ", err)
		}
		time.Sleep(10 * time.Millisecond)
		AssertPlayerWin(t, store, winner)
	})
}

func checkSchedulingCases(cases []cmd.ScheduledAlert, t *testing.T, blindAlerter *cmd.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}
			got := blindAlerter.Alerts[i]
			AssertScheduledAlert(t, got.String(), want.String())
		})
	}
}
func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}
