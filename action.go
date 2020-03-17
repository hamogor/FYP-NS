package main

type Action struct {
	Owner *Actor
	Action ActionFunc
	Pos    Position
}

type ActionFunc func(a *Actor, pos Position, l *Level, assets *Assets)

func (a *Actor) move(pos Position, g *Game) {
	if pos.X < a.Pos.X {
		a.Direction = -1
	} else if pos.X > a.Pos.X {
		a.Direction = 1
	}
	if !g.Level.Tiles[pos.X][pos.Y].Blocks {
		a.Pos = pos
		PPos = pos
	}
	if g.Level.Tiles[pos.X][pos.Y].Terrain == Door {
		pos.resolveDoorTypeAndOpen(g.Level, g.Assets)
		g.Player.Actor.Fov.Block(pos.X, pos.Y, false)
		g.Player.Actor.calculateFov()
	}
}