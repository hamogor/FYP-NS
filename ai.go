package main

type Ai struct {
}

func (a *Actor) move(pos Position, g *Game) {
	if pos.X < a.Pos.X {
		a.Direction = -1
	} else if pos.X > a.Pos.X {
		a.Direction = 1
	}
	if !g.Level.Tiles[pos.X][pos.Y].Blocks {
		a.Pos = pos
	}
	if g.Level.Tiles[pos.X][pos.Y].Terrain == Door {
		pos.resolveDoorTypeAndOpen(g.Level, g.Assets)
		g.Player.Actor.Fov.Block(pos.X, pos.Y, false)
		g.Player.Actor.calculateFov()
	}
}
