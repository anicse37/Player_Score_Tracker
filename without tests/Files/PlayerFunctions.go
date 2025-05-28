package files

func (p *PlayerDatabase) GetLeague() League {
	return p.league
}
func (p *PlayerDatabase) GetPlayerWins(name string) int {
	player := p.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}
func (p *PlayerDatabase) RecordWin(name string) {
	player := p.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		p.league = append(p.league, *NewPlayer(name))
	}
	p.jsonFile.Encode(p.league)
}
func (L League) Find(name string) *Player {
	for i, player := range L {
		if player.Name == name {
			return &L[i]
		}
	}
	return nil
}
