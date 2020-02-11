package main

import "github.com/faiface/pixel"

type Tile struct {
	Terrain Terrain
	Sprites Sprites
	Blocks  bool
	Bitmask  int
}

type Sprites struct {
	L *pixel.Sprite
	D *pixel.Sprite
}

type Terrain int

const (
	Wall Terrain = iota
	Floor
	NilTerrain
)

func (pos Position) terrain(l *Level) Terrain {
	if pos.inLevel() {
		return l.Tiles[pos.X][pos.Y].Terrain
	} else {
		return NilTerrain
	}

}

func floor() *Tile {
	return &Tile{
		Terrain: Floor,
		Sprites: Sprites{},
		Blocks:  false,
		Bitmask:  0,
	}
}

func wall() *Tile {
	return &Tile{
		Terrain: Wall,
		Sprites: Sprites{},
		Blocks:  true,
		Bitmask:  0,
	}
}
