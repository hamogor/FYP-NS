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
	Update()
}

func (ai *AiManager) Add(state AiState, calcFunc Calc, level Map, goal Goal) {
	ai.Composite[state] = BlankDMap(level, DiagonalNeighbours, calcFunc)
	ai.Composite[state].CalcFunc(ai.Actors, ai.Composite[state], goal)
}

func (ai *AiManager) Remove(state AiState) {
	ai.Composite[state] = nil
}

func (ai *AiManager) Update() {
	ai.CalulateFovs() 			 // 1. Calculate Field of Views
	ai.CheckTransitions()        // 2. Update active composite states
	ai.SetMaps()                 // 3. Copy Composite maps to each NPC.
}

func NewAiManager() *AiManager {
	return &AiManager{
		Composite: make(map[AiState]*DijkstraMap, 0),
		Actors: make([]*Actor, 0),
	}
}

func (ai *AiManager) CheckTransitions() {
	for i := range ai.Actors {
		switch ai.Actors[i].State {
		case MoveAi:
			ai.Actors[i].MoveTransition(); break
		case RangeAi:
			ai.Actors[i].RangeTransition(); break
		case FleeAi:
			ai.Actors[i].FleeTransition(); break
		case FlankAi:
			ai.Actors[i].FlankTransition(); break
		}
	}
}

func (ai *AiManager) SetMaps() {
	needed := make(map[AiState]bool, 0)
	calculated := make(map[AiState]bool, 0)
	for i := range ai.Actors {
		switch ai.Actors[i].State {
		case MoveAi:
			if !calculated[MoveAi] {
				ai.Composite[MoveAi].CalcFunc(ai.Actors, ai.Composite[MoveAi], PPos)
				needed[MoveAi] = true
				calculated[MoveAi] = true
				ai.Actors[i].DMap = ai.Composite[MoveAi].copy()
				break
			}
		case RangeAi:
			if !calculated[RangeAi] {
				ai.Composite[RangeAi].CalcFunc(ai.Actors, ai.Composite[RangeAi], PPos)
				needed[RangeAi] = true
				calculated[RangeAi] = true
				ai.Actors[i].DMap = ai.Composite[RangeAi].copy()
				break
			}
		case FlankAi:
			if !calculated[FlankAi] {
				ai.Composite[FlankAi].CalcFunc(ai.Actors, ai.Composite[FlankAi], PPos)
				needed[FlankAi] = true
				calculated[FlankAi] = true
				ai.Actors[i].DMap = ai.Composite[FlankAi].copy()
				break
			}
		case FleeAi:
			if !calculated[FleeAi] {
				ai.Composite[FleeAi].CalcFunc(ai.Actors, ai.Composite[FleeAi], PPos)
				needed[FleeAi] = true
				calculated[FleeAi] = true
				ai.Actors[i].DMap = ai.Composite[FleeAi].copy()
				break
			}
		}
	}
}


