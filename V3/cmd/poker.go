package cmd

import (
	"time"

	server "github.com/anicse37/Player_Score_Tracker/Servers"
)

type Poker struct {
	alerter BlindAlerter
	store   server.PlayerStore
}

type GameSpy struct {
	StartWith    int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartWith = numberOfPlayers
}
func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func (p *Poker) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		p.alerter.ScheduledAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}

}
func (p *Poker) Finish(winner string) {
	p.store.RecordWin(winner)
}
func NewGame(alerter BlindAlerter, store server.PlayerStore) *Poker {
	return &Poker{
		alerter: alerter,
		store:   store,
	}
}
