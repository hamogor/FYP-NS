package main

import (
	"math"
	"math/rand"
	"time"
)


type Level struct {
	Tiles  [LevelW][LevelH]*Tile
	Spawn  Position
	Rooms  []Rectangle
	Doors  []Position
	Items  [LevelW][LevelH]*Item
}

func (l *Level) SizeX() int {
	return LevelW
}

func (l *Level) SizeY() int {
	return LevelH
}

func (l *Level) IsPassable(x, y int) bool {
	if l.Tiles[x][y].Terrain == Wall || l.Tiles[x][y].Terrain == Door {
		return false
	} else {
		return true
	}
}

func (l *Level) OOB(x, y int) bool {
	if x > 0 && y > 0 && x <= LevelW && y <= LevelH {
		return false
	} else {
		return true
	}
}

// Objective per floor, can't go to next without completing.
func (g *Game) initAiSystem() {
	g.Ai = NewAiManager()
}

func resolveDoors(l *Level, a *Assets) {
	for i := range l.Doors {
		pos := Position{X: l.Doors[i].X, Y: l.Doors[i].Y}
		if pos.inLevel() {
			if pos.W().wall(l) && pos.E().wall(l) {
				l.Tiles[pos.X][pos.Y] = door("door_w_e", 0, a)
			} else if pos.N().wall(l) && pos.S().wall(l) {
				l.Tiles[pos.X][pos.Y] = door("door_n_s", 1, a)
			}
		}
	}
}



func (pos Position) ResolveBitMaskWall(l *Level) []bool {
	bit := make([]bool, 8)
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
	return bit
}

func (pos Position) ResolveBitMaskFloor(l *Level) []bool {
	bit := make([]bool, 8)
	if pos.SW().floor(l) {
		if pos.S().floor(l) && pos.W().floor(l) {
			bit[7] = true
		}
	}
	if pos.S().floor(l) {
		bit[6] = true
	}
	if pos.SE().floor(l) {
		if pos.S().floor(l) && pos.E().floor(l) {
			bit[5] = true
		}
	}
	if pos.W().floor(l) {
		bit[4] = true
	}
	if pos.E().floor(l) {
		bit[3] = true
	}
	if pos.NW().floor(l) {
		if pos.N().floor(l) && pos.W().floor(l) {
			bit[2] = true
		}
	}
	if pos.N().floor(l) {
		bit[1] = true
	}
	if pos.NE().floor(l) {
		if pos.N().floor(l) && pos.E().floor(l) {
			bit[0] = true
		}
	}
	return bit
}

func BoolListToMask(bit []bool) int {
	var n int
	for i, r := range bit {
		if r {
			n += 1 << (8 - i - 1)
		}
	}
	mask := 0
	switch n {
	case 2:
		mask = 1
		break
	case 8:
		mask = 2
		break
	case 10:
		mask = 3
		break
	case 11:
		mask = 4
		break
	case 16:
		mask = 5
		break
	case 18:
		mask = 6
		break
	case 22:
		mask = 7
		break
	case 24:
		mask = 8
		break
	case 26:
		mask = 9
		break
	case 27:
		mask = 10
		break
	case 30:
		mask = 11
		break
	case 31:
		mask = 12
		break
	case 64:
		mask = 13
		break
	case 66:
		mask = 14
		break
	case 72:
		mask = 15
		break
	case 74:
		mask = 16
		break
	case 75:
		mask = 17
		break
	case 80:
		mask = 18
		break
	case 82:
		mask = 19
		break
	case 86:
		mask = 20
		break
	case 88:
		mask = 21
		break
	case 90:
		mask = 22
		break
	case 91:
		mask = 23
		break
	case 94:
		mask = 24
		break
	case 95:
		mask = 25
		break
	case 104:
		mask = 26
		break
	case 106:
		mask = 27
		break
	case 107:
		mask = 28
		break
	case 120:
		mask = 29
		break
	case 122:
		mask = 30
		break
	case 123:
		mask = 31
		break
	case 126:
		mask = 32
		break
	case 127:
		mask = 33
		break
	case 208:
		mask = 34
		break
	case 210:
		mask = 35
		break
	case 214:
		mask = 36
		break
	case 216:
		mask = 37
		break
	case 218:
		mask = 38
		break
	case 219:
		mask = 39
		break
	case 222:
		mask = 40
		break
	case 223:
		mask = 41
		break
	case 248:
		mask = 42
		break
	case 250:
		mask = 43
		break
	case 251:
		mask = 44
		break
	case 254:
		mask = 45
		break
	case 255:
		mask = 46
		break
	case 0:
		mask = 47
		break
	}
	return mask
}

type Rectangle struct {
	X, Y, Width, Height int
}

const (
	MaxRoomSize = 6
	MinRoomSize = 3
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func (r *Rectangle) center() Position {
	return Position{
		X: (r.X + (r.X + r.Width)) / 2,
		Y: (r.Y + (r.Y + r.Height)) / 2,
	}
}

func generateLevel(g *Game) {
	l := &Level{
		Tiles:  [LevelW][LevelH]*Tile{},
		Spawn:  Position{X: 2, Y: 2},
		Rooms:  []Rectangle{},
		Doors:  []Position{},
	}
	roomSize := 6
	l.fillVoid(g.Assets)
	l.createRoomGrid(float64(roomSize), g.Assets)
	l.createDoors(roomSize, g.Assets)
	l.crushWalls(roomSize, 0.45, g.Assets)
	l.clearWallPoints(g.Assets)
	l.applyBitmask(g.Assets)
	g.Level = l
}

func (l *Level) createRoomGrid(roomSize float64, a *Assets) {
	xBorderIndex := ((LevelW - 1) / math.Floor(roomSize+1)) * (roomSize + 1)
	yBorderIndex := ((LevelH - 1) / math.Floor(roomSize+1)) * (roomSize + 1)

	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if float64(x) > xBorderIndex || float64(y) > yBorderIndex {
				l.Tiles[x][y] = tile()
			} else if x%(int(roomSize)+1) == 0 || y%(int(roomSize)+1) == 0 {
				l.Tiles[x][y] = wall(0, a)
			}
		}
	}
}

func (l *Level) fillVoid(a *Assets) {
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			l.Tiles[x][y] = floor(0, a)
		}
	}
}

func (l *Level) createDoors(roomSize int, a *Assets) {
	xRooms := (LevelW - 1) / math.Floor(float64(roomSize)+1)
	yRooms := (LevelH - 1) / math.Floor(float64(roomSize)+1)
	for i := 1; i < int(xRooms); i++ {
		for j := 1; j < int(yRooms); j++ {
			if i == 1 {
				y := i*(roomSize+1) - random(1, roomSize)
				x := j * (roomSize + 1)
				l.Tiles[x][y] = door("door_n_s", 1, a)
			}
			if j == 1 {
				y := i * (roomSize + 1)
				x := j*(roomSize+1) - random(1, roomSize)
				l.Tiles[x][y] = door("door_n_s", 1, a)
			}
			y := i * (roomSize + 1)
			x := j*(roomSize+1) + random(1, roomSize)
			l.Tiles[x][y] = door("door_n_s", 1, a)

			y = i*(roomSize+1) + random(1, roomSize)
			x = j * (roomSize + 1)
			l.Tiles[x][y] = door("door_n_s", 1, a)
		}
	}
}

func (l *Level) crushWalls(roomSize int, deleteChance float64, a *Assets) {
	xRooms := (LevelW - 1) / math.Floor(float64(roomSize)+1)
	yRooms := (LevelH - 1) / math.Floor(float64(roomSize)+1)
	for i := 1; i < int(xRooms); i++ {
		for j := 0; j < int(yRooms); j++ {
			y := i * (roomSize + 1)
			x := j*(roomSize+1) + 1
			chance := 0.0 + rand.Float64()*(1.0-0.0)
			if chance < deleteChance {
				for k := 0; k < roomSize; k++ {
					l.Tiles[x][y] = floor(0, a)
					x += 1
				}
			}
		}
	}
	for i := 1; i < int(yRooms); i++ {
		for j := 1; j < int(xRooms); j++ {
			y := i*(roomSize+1) + 1
			x := j * (roomSize + 1)
			chance := 0.0 + rand.Float64()*(1.0-0.0)
			if chance < deleteChance {
				for k := 0; k < roomSize; k++ {
					l.Tiles[x][y] = floor(0, a)
					y += 1
				}
			}
		}
	}
}

func (l *Level) clearWallPoints(a *Assets) {
	for x := 0; x < LevelW-1; x++ {
		for y := 0; y < LevelH-1; y++ {
			if l.Tiles[x][y].Terrain == Wall {
				pos := Position{X: x, Y: y}
				if pos.N().terrain(l) == Floor &&
					pos.S().terrain(l) == Floor &&
					pos.W().terrain(l) == Floor &&
					pos.E().terrain(l) == Floor &&
					pos.NW().terrain(l) == Floor &&
					pos.NE().terrain(l) == Floor &&
					pos.SW().terrain(l) == Floor &&
					pos.SE().terrain(l) == Floor {
					l.Tiles[x][y] = floor(0, a)
				}

			}
		}
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


func (g *Game) BuildStaticLevel(level string) {
	l := &Level{
		Tiles:  [LevelW][LevelH]*Tile{},
		Spawn:  Position{X: 1, Y: 1},
	}

	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			l.Items[x][y] = &Item{
				Name:    "Nil",
				Sprites: Sprites{},
				Active:  false,
				Type:    2,
				Handler: nil,
			}
		}
	}

	x, y := 0, 0
	for _, c := range level {
		parsed := false
		if c == '#' {
			l.Tiles[x][y] = wall(0, g.Assets)
			parsed = true
		} else if c == '.' {
			l.Tiles[x][y] = floor(0, g.Assets)
			parsed = true
		} else if c == 'd' {
			l.Tiles[x][y] = door("door_n_s", 0, g.Assets)
			parsed = true
		} else if c == 'm' {
			l.Tiles[x][y] = floor(0, g.Assets)
			g.newActor(MoveAi, Position{X: x, Y: y})
			parsed = true
		} else if c == 'r' {
			l.Tiles[x][y] = floor(0, g.Assets)
			parsed = true
		} else if c == 'f' {
			l.Tiles[x][y] = floor(0, g.Assets)
			parsed = true
		} else if c == 'a' {
			l.Tiles[x][y] = floor(0, g.Assets)
			parsed = true
		} else if c == 'h' {
			l.Tiles[x][y] = floor(0, g.Assets)
			parsed = true
		} else if c == 's' {
			l.Spawn = Position{X: x, Y: y}
		}

		if parsed {
			if x >= LevelW-1 {
				x = 0
				y++
			} else {
				x++
			}
		}
	}
	l.applyBitmask(g.Assets)
	g.Level = l
}


