package main

import (
	"math"
	"math/rand"
)

type LNode struct {
}

const (
	MaxRoomSize = 6
	MinRoomSize = 3
)

func generateLevel(g *Game) {
	l := &Level{
		Tiles:  [64][64]*Tile{},
		Spawn:  Position{X: 2, Y: 2},
		Rooms:  []Rectangle{},
		Doors:  []Position{},
		Actors: make([]*Actor, 0),
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
