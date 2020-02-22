package main

type LNode struct {

}

const (
	MaxRoomSize = 6
	MinRoomSize = 3
)


func generateLevel(g *Game) {
	l := &Level{
		Tiles: [64][64]*Tile{},
		Spawn: Position{X: 1, Y: 1},
		Rooms: []Rectangle{},
		Doors: []Position{},
		Actors: make([]*Actor, 0),
	}
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			l.Tiles[x][y] = floor(0, g.Assets)
			if x == 0 || x == LevelW || x == LevelW - 1 {
				l.Tiles[x][y] = wall(0, g.Assets)
			}
			if y == 0 || y == LevelH || y == LevelH - 1 {
				l.Tiles[x][y] = wall(0, g.Assets)
			}
		}
	}

	generateRoom(l)



	l.print()
	l.applyBitmask(g.Assets)
	g.Level = l
}

func generateRoom(l *Level) {
	if len(l.Rooms) == 0 {

	}
}

func (l *Level) applyBitmask(a *Assets) {
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			pos := Position{X: x, Y: y}
			if pos.inLevel() {
				if pos.terrain(l) == Wall {
					mask := BoolListToMask(pos.ResolveBitMaskWall(l))
					l.Tiles[x][y] = wall(mask, a)
				} else if pos.terrain(l) == Floor {
					mask := BoolListToMask(pos.ResolveBitMaskFloor(l))
					l.Tiles[x][y] = floor(mask, a)
				} else if pos.terrain(l) == Door {
					if pos.W().wall(l) && pos.E().wall(l) {
						l.Tiles[x][y] = door("door_w_e", 0, a)
					} else if pos.N().wall(l) && pos.S().wall(l) {
						l.Tiles[x][y] = door("door_n_s", 1, a)
					}
				}
			}

		}
	}
}

