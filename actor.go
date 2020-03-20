package main

type Actor struct {
	Name        string
	HP          int
	Pos         Position
	Anims       map[string]*Animation
	Fov         *FovMap
	CAnim       *CAnim
	Action      *Action
	ActionTaken bool
	Direction   float64
	State       AiState
	Type        AiState
	Ammo        int
	DMap        *DijkstraMap
	Points      []Point
}

func newPlayerActor(name string, pos Position, a *Assets) *Actor {
	anims := make(map[string]*Animation, 0)

	anims["player_idle"] = a.Anims["player_idle"]

	return &Actor{
		Name:      name,
		Pos:       pos,
		Anims:     anims,
		CAnim:     buildCAnim(),
		Direction: -1,
		HP:        100,
		State:     PlayerAi,
	}
}

func (g *Game) newActor(state AiState, pos Position) {
	anims := make(map[string]*Animation, 0)
	anims["enemy_idle"] = g.Assets.Anims["enemy_idle"]
	a := &Actor{
		Name:        string(state),
		HP:          100,
		Pos:         pos,
		Anims:       anims,
		Fov:         nil,
		CAnim:       buildCAnim(),
		Action:      nil,
		ActionTaken: false,
		Direction:   -1,
		State:       state,
		Type:        state,
		Ammo:        10,
		DMap:        nil,
	}
	g.Ai.Actors = append(g.Ai.Actors, a)
}

func (l *Level) newActor(state AiState, x, y int, a *Assets) {

}
