package main

type AiManager struct {
	Composite map[AiState]*DijkstraMap
	Level     *Level
	Actors []*Actor
	Actions []*Action
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
	Update(g *Game)
}

func (ai *AiManager) Add(state AiState, calcFunc Calc, level Map, goal Goal) {
	ai.Composite[state] = BlankDMap(level, DiagonalNeighbours, calcFunc)
	ai.Composite[state].CalcFunc(ai.Actors, ai.Composite[state], goal)
}

func (ai *AiManager) Remove(state AiState) {
	ai.Composite[state] = nil
}

func (ai *AiManager) Update(g *Game) {
	ai.CalulateFovs() 			 // 1. Calculate Field of Views
	ai.CheckTransitions()        // 2. Update active composite states
	ai.SetMaps()                 // 3. Copy Composite maps to each NPC.
	ai.CheckUnderlying()         // 4. Check the non-composite behaviour
	ai.DecideActions()           // 5. Check DMap costs and push appropriate action
	ai.ExecuteActions(g)          // 6. Execute each action
}

func NewAiManager() *AiManager {
	return &AiManager{
		Composite: make(map[AiState]*DijkstraMap, 0),
		Actors: make([]*Actor, 0),
		Actions: make([]*Action, 0),
	}
}

func (ai *AiManager) CheckTransitions() {
	for i := range ai.Actors {
		switch ai.Actors[i].State {
		case MoveAi:
			ai.Actors[i].MoveTransition(ai.Level); break
		case RangeAi:
			ai.Actors[i].RangeTransition(ai.Level); break
		case FleeAi:
			ai.Actors[i].FleeTransition(ai.Level); break
		case FlankAi:
			ai.Actors[i].FlankTransition(ai.Level); break
		}
	}
}

func (ai *AiManager) CheckUnderlying() {
	for i := range ai.Actors {
		switch ai.Actors[i].State {
		case MoveAi:
			ai.Actors[i].MoveUnderlying(ai.Level); break
		case RangeAi:
			ai.Actors[i].RangeUnderlying(ai.Level); break
		case FleeAi:
			ai.Actors[i].FleeUnderlying(ai.Level); break
		case FlankAi:
			ai.Actors[i].FlankUnderlying(ai.Level); break
		}
	}
}

func (ai *AiManager) DecideActions() {
	for i := range ai.Actors {
		ln := ai.Actors[i].DMap.LowestNeighbour(ai.Actors[i].Pos.X, ai.Actors[i].Pos.Y)
		switch ai.Actors[i].State {
		case MoveAi:
			ai.Actors[i].MoveBehaviour(ai.Level, ln, ai); break
		case RangeAi:
			ai.Actors[i].RangeBehaviour(ai.Level, ln, ai); break
		case FleeAi:
			ai.Actors[i].FleeBehaviour(ai.Level, ln, ai); break
		case FlankAi:
			ai.Actors[i].FlankBehaviour(ai.Level, ln, ai); break
		}
	}
}

func (ai *AiManager) ExecuteActions(g *Game) {
	for i := range ai.Actions {
		ai.Actions[i].Action(ai.Actions[i].Owner, ai.Actions[i].Pos, g)
	}
	ai.Actions = make([]*Action, 0)
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




