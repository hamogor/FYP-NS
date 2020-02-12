package main

type Player struct {
	Actor *Actor
	Depth int
}

func (g *Game) initPlayer() {
	g.Player = &Player{
		Actor: newPlayerActor("Player", Position{X: 1, Y: 1}, g.Assets),
		Depth: 0,
	}
}
