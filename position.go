package main

import (
	"fmt"
	"github.com/faiface/pixel"
)

type Position struct {
	X, Y int
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

func (pos Position) cardinalWalls(l *Level) bool {
	n, e, s, w := false, false, false, false
	if pos.inLevel() {
		if pos.N().inLevel() && l.Tiles[pos.N().X][pos.N().Y].Terrain == Wall {
			n = true
		}
		if pos.E().inLevel() && l.Tiles[pos.E().X][pos.E().Y].Terrain == Wall {
			e = true
		}
		if pos.S().inLevel() && l.Tiles[pos.S().X][pos.S().Y].Terrain == Wall {
			s = true
		}
		if pos.W().inLevel() && l.Tiles[pos.W().X][pos.W().Y].Terrain == Wall {
			w = true
		}
		if n == true && e == true && s == true && w == true {
			return true
		}
	}
	return false
}

func (pos Position) sToVec() pixel.Vec {
	return pixel.V(float64(pos.X*SpriteW), float64(pos.Y*SpriteH))
}

func (pos Position) wall(l *Level) bool {
	if pos.inLevel() {
		fmt.Print(pos, "\n")
		if l.Tiles[pos.X][pos.Y].Terrain == Wall {
			return true
		} else {
			return false // Handles edge bitmask
		}
	}

	return true
}

func (pos Position) floor(l *Level) bool {
	if l.Tiles[pos.X][pos.Y].Terrain == Floor {
		return true
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
