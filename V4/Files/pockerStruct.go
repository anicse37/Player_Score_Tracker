package files

import (
	"encoding/json"
	"net/http"
	"time"
)

type Player struct {
	Name string
	Wins int
}
type League []Player

type PlayerSeeker struct {
	Database *json.Encoder
	league   League
}

type ScheduledAlert struct {
	At     time.Duration
	Anount int
}
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}
type PlayerServer struct {
	Store PlayerStore
	http.Handler
}
