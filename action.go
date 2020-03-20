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
	if !g.Level.Tiles[pos.X][pos.Y].Blocks && !pos.actor(g.Ai) {
		a.Pos = pos
		PPos = pos
	}
	if g.Level.Tiles[pos.X][pos.Y].Terrain == Door {
		pos.resolveDoorTypeAndOpen(g.Level, g.Assets)
		g.Player.Actor.Fov.Block(pos.X, pos.Y, false)
		g.Player.Actor.calculateFov()
	}
}

func (a *Actor) pickupItem(pos Position, l *Level) {
	if isItem(l, pos.X, pos.Y) {
		l.Items[pos.X][pos.Y].Active = false
		if l.Items[pos.X][pos.Y].Type == Health {
			a.HP += 50
		} else {
			a.Ammo += 10
		}
	}
}