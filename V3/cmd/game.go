package cmd

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
