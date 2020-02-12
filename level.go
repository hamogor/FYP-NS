package main

import (
	"github.com/faiface/pixel"
)

type Level struct {
	Tiles [LevelW][LevelH]*Tile
	Spawn Position
	Rooms []Rectangle
}

// Objective per floor, can't go to next without completing.
func (g *Game) initLevel() {
	d := NewDungeon(LevelW, 50)
	l := &Level{
		Tiles: [LevelW][LevelH]*Tile{},
		Spawn: Position{},
		Rooms: d.Rooms,
	}
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if d.Grid[x][y] == 1 {
				l.Tiles[x][y] = floor()
			} else if d.Grid[x][y] == 0 {
				l.Tiles[x][y] = wall()
			}
		}
	}
	l.Spawn = l.Rooms[0].center()
	//for x := 0; x < LevelW; x++ {
	//	for y := 0; y < LevelH; y++ {
	//		l.Tiles[x][y] = floor()
	//		if y == 0 || y == LevelH-1 {
	//			l.Tiles[x][y] = wall()
	//		}
	//		if x == 0 || x == LevelW-1 {
	//			l.Tiles[x][y]= wall()
	//		}
	//	}
	//}



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
				l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_floor"][bitmask])
				l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_floor"][bitmask])
			}
		}
	}
	g.Level = l

}



