package main

type Action struct {
	Owner *Actor
	Action ActionFunc
	Pos    Position
}

type ActionFunc func(a *Actor, pos Position, g *Game)

func move(a *Actor, pos Position, g *Game) {
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

func melee(a *Actor, pos Position, g *Game) {
	if pos.X < a.Pos.X {
		a.Direction = -1
	} else if pos.X > a.Pos.X {
		a.Direction = 1
	}
}


func pickupItem(a *Actor, pos Position, g *Game) {
	if isItem(g.Level, pos.X, pos.Y) {
		g.Level.Items[pos.X][pos.Y].Active = false
		if g.Level.Items[pos.X][pos.Y].Type == Health {
			a.HP += 50
		} else {
			a.Ammo += 10
		}
	}
}

func pushAction(ai *AiManager, a *Actor, action ActionFunc, pos Position) {
	ai.Actions = append(ai.Actions, &Action{
		Owner:  a,
		Action: action,
		Pos:    pos,
	})
}