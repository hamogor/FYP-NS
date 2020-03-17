package main

type AiManager struct {
	Composite map[AiState]*DijkstraMap
	Actors []*Actor
	Manager
}

type AiState int

type Calc func(actors []*Actor, d *DijkstraMap, points ...Point)
type Goal Position

func (g Goal) GetXY() (int, int) {
	if g.X < LevelW && g.Y < LevelH && g.X > 0 && g.Y > 0 {
		return g.X, g.Y
	}
	return 0, 0
}

type Manager interface {
	Add(AiState, Calc, Map, Goal)
	Remove(AiState)
	TakeTurn()
}

func (ai *AiManager) Add(state AiState, calcFunc Calc, level Map, goal Goal) {
	ai.Composite[state] = BlankDMap(level, DiagonalNeighbours, calcFunc)
	ai.Composite[state].CalcFunc(ai.Actors, ai.Composite[state], goal)
}

func (ai *AiManager) Remove(state AiState) {
	ai.Composite[state] = nil
}

func (ai *AiManager) TakeTurn() {

}

func NewAiManager() *AiManager {
	return &AiManager{
		Composite: make(map[AiState]*DijkstraMap, 0),
	}
}


