package functions

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{map[string]int{}}
}

type InMemoryStore struct {
	score map[string]int
}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return i.score[name]
}
func (i *InMemoryStore) RecordWin(name string) {
	i.score[name]++
}
