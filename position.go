package main

import (
	"github.com/faiface/pixel"
)

type Position struct {
	X, Y int
}

func (p Position) GetXY() (int, int) {
	if p.X < LevelW && p.Y < LevelH && p.X > 0 && p.Y > 0 {
		return p.X, p.Y
	}
	return 0, 0
}

func (pos Position) E() Position {
	return Position{X: pos.X + 1, Y: pos.Y}
}

func (pos Position) SE() Position {
	return Position{X: pos.X + 1, Y: pos.Y - 1}
}

func (pos Position) NE() Position {
	return Position{X: pos.X + 1, Y: pos.Y + 1}
}

func (pos Position) N() Position {
	return Position{X: pos.X, Y: pos.Y + 1}
}

func (pos Position) S() Position {
	return Position{X: pos.X, Y: pos.Y - 1}
}

func (pos Position) W() Position {
	return Position{X: pos.X - 1, Y: pos.Y}
}

func (pos Position) SW() Position {
	return Position{X: pos.X - 1, Y: pos.Y - 1}
}

func (pos Position) NW() Position {
	return Position{X: pos.X - 1, Y: pos.Y + 1}
}

func (pos Position) sToVec() pixel.Vec {
	return pixel.V(float64(pos.X*SpriteW), float64(pos.Y*SpriteH))
}

func (pos Position) wall(l *Level) bool {
	if pos.inLevel() {
		if l.Tiles[pos.X][pos.Y].Terrain == Wall {
			return true
		} else {
			return false // Handles edge bitmask
		}
	}

	return true
}

func isItem(l *Level, x, y int) bool {
	if l.Items[x][y].Active {
		return true
	} else {
		return false
	}
}

func (pos Position) floor(l *Level) bool {
	if l.Tiles[pos.X][pos.Y].Terrain == Floor {
		return true
	}
	return false
}

func (pos Position) actor(ai *AiManager) bool {
	for i := range ai.Actors {
		if pos == ai.Actors[i].Pos {
			return true
		}
	}
	return false
}

func (pos Position) inLevel() bool {
	if pos.X < 0 || pos.Y < 0 || pos.X >= LevelW || pos.Y >= LevelH {
		return false
	} else {
		return true
	}
}

func (r *Render) mouseTranslate() Position {
	return Position{}
}

func tilePos(x, y int) pixel.Matrix {
	return pixel.IM.Moved(pixel.V(float64(x)*TileW, float64(y)*TileH))
}

func RoundFloat(f float64) float64 {
	return float64(int64(f))
}

func (a *Actor) seenItem(iType ItemType, l *Level) bool {
	a.Points = make([]Point, 0)
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if l.Items[x][y].Active && l.Items[x][y].Type == iType && a.Fov.seen[x][y] {
				a.Points = append(a.Points, Position{X: x, Y: y})
			}
		}
	}

	if len(a.Points) > 0 {
		return true
	} else {
		return false
	}
}


