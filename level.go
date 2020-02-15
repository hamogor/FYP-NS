package main

import (
	"fmt"
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
	d := NewDungeon(LevelW+2, 20)
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
			if y == 0 || y == LevelH || y == LevelH-1 {
				l.Tiles[x][y] = wall()
			}
			if x == 0 || x == LevelW || x == LevelW -1 {
				l.Tiles[x][y] = wall()
			}
		}
	}

	for y := 0; y < LevelH; y++ {
		for x := 0; x < LevelW; x++ {
			if l.Tiles[x][y].Terrain == Wall {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	l.Spawn = l.Rooms[0].center()
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			pos := Position{X: x, Y: y}
			bit := make([]bool, 8)
			if pos.terrain(l) == Wall {
				if pos.SW().wall(l) {
					if pos.S().wall(l) && pos.W().wall(l) {
						bit[7] = true
					}
				}
				if pos.S().wall(l) {
					bit[6] = true
				}
				if pos.SE().wall(l) {
					if pos.S().wall(l) && pos.E().wall(l) {
						bit[5] = true
					}
				}
				if pos.W().wall(l) {
					bit[4] = true
				}
				if pos.E().wall(l) {
					bit[3] = true
				}
				if pos.NW().wall(l) {
					if pos.N().wall(l) && pos.W().wall(l) {
						bit[2] = true
					}
				}
				if pos.N().wall(l) {
					bit[1] = true
				}
				if pos.NE().wall(l) {
					if pos.N().wall(l) && pos.E().wall(l) {
						bit[0] = true
					}
				}

				bitmask := BoolListToMask(bit)
				mask := 0
				switch bitmask {
				case 2: mask = 1; break
				case 8: mask = 2; break
				case 10: mask = 3; break
				case 11: mask = 4; break
				case 16: mask = 5; break
				case 18: mask = 6; break
				case 22: mask = 7; break
				case 24: mask = 8; break
				case 26: mask = 9; break
				case 27: mask = 10; break
				case 30: mask = 11; break
				case 31: mask = 12; break
				case 64: mask = 13; break
				case 66: mask = 14; break
				case 72: mask = 15; break
				case 74: mask = 16; break
				case 75: mask = 17; break
				case 80: mask = 18; break
				case 82: mask = 19; break
				case 86: mask = 20; break
				case 88: mask = 21; break
				case 90: mask = 22; break
				case 91: mask = 23; break
				case 94: mask = 24; break
				case 95: mask = 25; break
				case 104: mask = 26; break
				case 106: mask = 27; break
				case 107: mask = 28; break
				case 120: mask = 29; break
				case 122: mask = 30; break
				case 123: mask = 31; break
				case 126: mask = 32; break
				case 127: mask = 33; break
				case 208: mask = 34; break
				case 210: mask = 35; break
				case 214: mask = 36; break
				case 216: mask = 37; break
				case 218: mask = 38; break
				case 219: mask = 39; break
				case 222: mask = 40; break
				case 223: mask = 41; break
				case 248: mask = 42; break
				case 250: mask = 43; break
				case 251: mask = 44; break
				case 254: mask = 45; break
				case 255: mask = 46; break
				case 0: mask = 47; break
				}



				l.Tiles[x][y].Bitmask = mask
				l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][mask])
				l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][mask])

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

	// Handle edge cases.
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			//pos := Position{X: x, Y: y}

			//if pos.W().terrain(l) == NilTerrain {
			//	l.Tiles[x][y].Bitmask = 29
			//	l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][28])
			//	l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][28])
			//}
			//if pos.W().terrain(l) == NilTerrain && pos.S().wall(l) && pos.E().wall(l) {
			//	l.Tiles[x][y].Bitmask = 45
			//	l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][44])
			//	l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][44])
			//}
			//if pos.N().wall(l) && pos.E().wall(l) && pos.W().terrain(l) == NilTerrain {
			//	l.Tiles[x][y].Bitmask = 34
			//	l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][33])
			//	l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][33])
			//}
			//if pos.N().terrain(l) == NilTerrain && pos.E().wall(l) && pos.W().wall(l) && pos.S().floor(l) {
			//	l.Tiles[x][y].Bitmask = 43
			//	l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][42])
			//	l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][42])
			//}
			//if pos.S().terrain(l) == NilTerrain && pos.W().wall(l) && pos.E().wall(l) && pos.N().floor(l) {
			//	l.Tiles[x][y].Bitmask = 13
			//	l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][12])
			//	l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][12])
			//}
			//if pos.E().terrain(l) == NilTerrain && pos.N().wall(l) && pos.S().wall(l) {
			//	l.Tiles[x][y].Bitmask = 37
			//	l.Tiles[x][y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["l_wall"][36])
			//	l.Tiles[x][y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["d_wall"][36])
			//}
		}
	}









	g.Level = l
}

func (pos Position) placeDoorIfPossible(l *Level, a *Assets) {
	for i := range l.Doors { // BUGGED, CHECK IS POS IS IN LEVEL
		if pos.inLevel() && pos.W().inLevel() && pos.N().inLevel() && pos.E().inLevel() && pos.S().inLevel() {
			if l.Tiles[l.Doors[i].X-1][l.Doors[i].Y].Terrain == Wall &&
				l.Tiles[l.Doors[i].X+1][l.Doors[i].Y].Terrain == Wall &&
				l.Tiles[l.Doors[i].X][l.Doors[i].Y-1].Terrain != Wall &&
				l.Tiles[l.Doors[i].X][l.Doors[i].Y+1].Terrain != Wall {
				l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.L = pixel.NewSprite(a.Sheets.Environment, a.Env["door_w_e"][0])
				l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.D = pixel.NewSprite(a.Sheets.Environment, a.Env["door_w_e"][1])
				l.Tiles[l.Doors[i].X][l.Doors[i].Y].Terrain = Door
			} else if l.Tiles[l.Doors[i].X][l.Doors[i].Y-1].Terrain == Wall &&
				l.Tiles[l.Doors[i].X][l.Doors[i].Y+1].Terrain == Wall &&
				l.Tiles[l.Doors[i].X-1][l.Doors[i].Y+1].Terrain != Wall &&
				l.Tiles[l.Doors[i].X+1][l.Doors[i].Y].Terrain != Wall {
				l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.L = pixel.NewSprite(a.Sheets.Environment, a.Env["door_n_s"][0])
				l.Tiles[l.Doors[i].X][l.Doors[i].Y].Sprites.D = pixel.NewSprite(a.Sheets.Environment, a.Env["door_n_s"][1])
				l.Tiles[l.Doors[i].X][l.Doors[i].Y].Terrain = Door
			}
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
