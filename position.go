package main

import (
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

func (pos Position) sToVec() pixel.Vec {
	return pixel.V(float64(pos.X*SpriteW), float64(pos.Y*SpriteH))
}

func (pos Position) inLevel() bool {
	if pos.X < 0 || pos.Y < 0 || pos.X >= LevelW || pos.Y >= LevelH {
		return false
	}
	return true
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
