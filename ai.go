package main

type Ai struct {
	Composite map[AiState]*DijkstraMap
	Actions []*Action
}

type AiState int

const (
	MoveAi   AiState = iota
	FleeAi           = 1
	FlankAi          = 2
	RangeAi          = 3
	PlayerAi         = 4
)


