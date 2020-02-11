package main

type Ai struct {
}

func (a *Actor) move(pos Position, l *Level) {
	if pos.X < a.Pos.X {
		a.Direction = -1
	} else if pos.X > a.Pos.X {
		a.Direction = 1
	}
	if !l.Tiles[pos.X][pos.Y].Blocks {
		a.Pos = pos
	}

}
