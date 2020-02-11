package main

type Actor struct {
	Name        string
	Pos         Position
	Anims       map[string]*Animation
	CAnim       *CAnim
	ActionTaken bool
	Direction   float64
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
	}
}
