package main

import (
	"github.com/faiface/pixel"
)

type Level struct {
	Tiles [LevelW][LevelH]*Tile
	Spawn Position
	Rooms []Rectangle
	Doors []Position
}

// Objective per floor, can't go to next without completing.
func (g *Game) initLevel() {
	d := NewDungeon(LevelW, 20)
	l := &Level{
		Tiles: [LevelW][LevelH]*Tile{},
		Spawn: Position{},
		Rooms: d.Rooms,
		Doors: d.Doors,
	}
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if d.Grid[x][y] == 1 {
				l.Tiles[x][y] = floor()
			} else if d.Grid[x][y] == 0 {
				l.Tiles[x][y] = wall()
			}
			if y == 0 || y == LevelH-1 {
				l.Tiles[x][y] = wall()
			}
			if x == 0 || x == LevelW-1 {
				l.Tiles[x][y] = wall()
			}
		}
	}
	l.Spawn = l.Rooms[0].center()

	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			pos := Position{X: x, Y: y}

			if pos.terrain(l) == Wall {
				bitmask := 0

				if pos.S().terrain(l) == Wall {
					bitmask += 1
				}
				if pos.E().terrain(l) == Wall {
					bitmask += 2
				}
				if pos.N().terrain(l) == Wall {
					bitmask += 4
				}
				if pos.W().terrain(l) == Wall {
					bitmask += 8
				}

				l.Tiles[x][y].Bitmask = bitmask
				l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][bitmask])
				l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][bitmask])

			} else if pos.terrain(l) == Floor {
				bitmask := 0
				if pos.S().terrain(l) == Wall {
					bitmask += 1
				}
				if pos.E().terrain(l) == Wall {
					bitmask += 2
				}
				if pos.N().terrain(l) == Wall {
					bitmask += 4
				}
				if pos.W().terrain(l) == Wall {
					bitmask += 8
				}
				l.Tiles[x][y].Bitmask = bitmask
				l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_floor"][bitmask])
				l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_floor"][bitmask])
			}
			pos.placeDoorIfPossible(l, g.Assets)
		}
	}

	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			pos := Position{X: x, Y: y}
			if pos.cardinalWalls(l) {
				if pos.SW().floor(l) && pos.NW().wall(l) && pos.NE().wall(l) && pos.SE().wall(l) {
					l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][16])
					l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][16])

				}
				if pos.cardinalWalls(l) {
					if pos.NW().wall(l) && pos.SE().wall(l) && pos.SW().floor(l) && pos.NE().floor(l) {
						l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][17])
						l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][17])
					}
				}
			}
		}
	}

	//for y := 0; y < LevelH; y++ {
	//	for x := 0; x < LevelW; x++ {
	//		if l.Tiles[x][y].Terrain == Wall {
	//			fmt.Print(l.Tiles[x][y].Bitmask, " ")
	//		} else {
	//			fmt.Print(" . ")
	//		}
	//	}
	//	fmt.Println()
	//}

	g.Level = l
}

func (pos Position) placeDoorIfPossible(l *Level, a *Assets) {
	for i := range l.Doors {
		if l.Tiles[l.Doors[i].X-1][l.Doors[i].Y].Terrain == Wall &&
			l.Tiles[l.Doors[i].X+1][l.Doors[i].Y].Terrain == Wall &&
			l.Tiles[l.Doors[i].X][l.Doors[i].Y-1].Terrain != Wall &&
			l.Tiles[l.Doors[i].X][l.Doors[i].Y+1].Terrain != Wall {
			l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.L = pixel.NewSprite(a.Sheets.Environment, a.Env["door_w_e"][0])
			l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.D = pixel.NewSprite(a.Sheets.Environment, a.Env["door_w_e"][1])
		} else if l.Tiles[l.Doors[i].X][l.Doors[i].Y-1].Terrain == Wall &&
			l.Tiles[l.Doors[i].X][l.Doors[i].Y+1].Terrain == Wall &&
			l.Tiles[l.Doors[i].X-1][l.Doors[i].Y+1].Terrain != Wall &&
			l.Tiles[l.Doors[i].X+1][l.Doors[i].Y].Terrain != Wall {
			l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.L = pixel.NewSprite(a.Sheets.Environment, a.Env["door_n_s"][0])
			l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.D = pixel.NewSprite(a.Sheets.Environment, a.Env["door_n_s"][1])
		}

	}
}

func (pos Position) setWallBitmask(a *Assets, l *Level, bit int) {
	l.Tiles[pos.X][pos.Y].Sprites.L = pixel.NewSprite(a.Sheets.Environment, a.Env["l_wall"][bit])
	l.Tiles[pos.X][pos.Y].Sprites.D = pixel.NewSprite(a.Sheets.Environment, a.Env["d_wall"][bit])
}

func BoolListToMask(bit []bool) int {
	var n int
	for i, r := range bit {
		if r {
			n += 1 << (8 - i - 1)
		}
	}
	return n
}
