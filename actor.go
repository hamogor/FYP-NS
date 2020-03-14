package main

type Actor struct {
	Name        string
	HP          int
	Pos         Position
	Anims       map[string]*Animation
	Fov         *FovMap
	CAnim       *CAnim
	ActionTaken bool
	Direction   float64
	State       AiState
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
	}
}

func (l *Level) newActor(state AiState, x, y int, a *Assets) {

}
