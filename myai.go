package main

const (
	MoveAi   AiState = iota
	FleeAi           = 1
	FlankAi          = 2
	RangeAi          = 3
)

func (g *Game) initAi() {
	g.Ai = NewAiManager()
	g.Ai.Add(MoveAi, MoveCalculate, g.Level, Goal(g.Player.Actor.Pos))
	g.Ai.Add(FleeAi, FleeCalculate, g.Level, Goal(g.Player.Actor.Pos))
	g.Ai.Add(FlankAi, FlankCalculate, g.Level, Goal(g.Player.Actor.Pos))
	g.Ai.Add(RangeAi, RangeCalculate, g.Level, Goal(g.Player.Actor.Pos))

}

func MoveCalculate(actors []*Actor, d *DijkstraMap, points ...Point) {
	d.Calc(points...)
}

func FleeCalculate(actors []*Actor, d *DijkstraMap, points ...Point) {
	for x := range d.Points {
		for y := range d.Points[x] {
			if d.M.IsPassable(x, y) {
				d.Points[x][y] = 50 - d.Points[x][y]
			}
		}
	}
}

func RangeCalculate(actors []*Actor, d *DijkstraMap, points ...Point) {
	neighbors := DiagonalNeighboursAtRadius4(PPos.X, PPos.Y)
	tiles := make([]Point, 0)
	for i := range neighbors {
		if neighbors[i].X > 0 && neighbors[i].X < LevelW && neighbors[i].Y > 0 && neighbors[i].Y < LevelH {
			if d.M.IsPassable(neighbors[i].X, neighbors[i].Y) {
				tiles = append(tiles, neighbors[i])
			}
		}
	}
	d.Recalc(tiles...)
}

func FlankCalculate(actors []*Actor, d *DijkstraMap, points ...Point) {
	alreadyIncreased := make([]Position, 0)
	for i := range actors {
		if actors[i].State == FlankAi {
			neighbours := DiagonalNeighboursAtRadius2(d, actors[i].Pos.X, actors[i].Pos.Y)
			for j := range neighbours {
				if neighbours[j].X < LevelW && neighbours[j].Y < LevelH && neighbours[j].X > 0 && neighbours[j].Y > 0 {
					for k := range alreadyIncreased {
						if alreadyIncreased[k].X == neighbours[j].X && alreadyIncreased[k].Y == neighbours[j].Y {
							break
						} else {
							d.Points[neighbours[j].X][neighbours[j].Y] += 3
							alreadyIncreased = append(alreadyIncreased, Position{X: neighbours[j].X, Y: neighbours[j].Y})
						}
					}

				}
			}
		}
	}
}

func (a *Actor) MoveTransition() {
	if a.HP <= 50 {
		a.State = FleeAi
	}
	//if a.Type == RangeAi && AMMO_AVAILABLE { state = type }
}

func (a *Actor) FleeTransition() {
	if a.HP > 50 {
		a.State = a.Type
	}
}

func (a *Actor) RangeTransition() {
	if a.Ammo == 0 {
		a.State = MoveAi
	}
}

func (a *Actor) FlankTransition() {

}

func (a *Actor) MoveUnderlying() {

}

func (a *Actor) RangeUnderlying() {

}

func (a *Actor) FleeUnderlying() {

}

func (a *Actor) FlankUnderlying() {

}



