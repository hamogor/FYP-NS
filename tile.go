package main

import "github.com/faiface/pixel"

type Tile struct {
	Terrain Terrain
	Sprites Sprites
	Blocks  bool
	Bitmask int
}

type Sprites struct {
	L *pixel.Sprite
	D *pixel.Sprite
}

type Terrain int

const (
	Wall Terrain = iota
	Floor
	Door
	OpenDoor
	NilTerrain
)

func (pos Position) terrain(l *Level) Terrain {
	if pos.inLevel() {
		return l.Tiles[pos.X][pos.Y].Terrain
	} else {
		return NilTerrain
	}
}

func (pos Position) bitmask(l *Level) int {
	if pos.inLevel() {
		return l.Tiles[pos.X][pos.Y].Bitmask
	}
	return 0
}

func tile() *Tile {
	return &Tile{
		Terrain: 0,
		Sprites: Sprites{},
		Blocks:  false,
		Bitmask: 0,
	}
}

func floor(mask int, a *Assets) *Tile {
	return &Tile{
		Terrain: Floor,
		Sprites: Sprites{
			L: pixel.NewSprite(a.Sheets.Environment, a.Env["l_floor"][mask]),
			D: pixel.NewSprite(a.Sheets.Environment, a.Env["d_floor"][mask]),
		},
		Blocks:  false,
		Bitmask: mask,
	}
}

func wall(mask int, a *Assets) *Tile {
	return &Tile{
		Terrain: Wall,
		Sprites: Sprites{
			L: pixel.NewSprite(a.Sheets.Environment, a.Env["l_wall"][mask]),
			D: pixel.NewSprite(a.Sheets.Environment, a.Env["d_wall"][mask]),
		},
		Blocks:  true,
		Bitmask: mask,
	}
}

// Doors bitmask is 0 == w_e && 1 == n_s
func door(dir string, mask int, a *Assets) *Tile {
	return &Tile{
		Terrain: Door,
		Sprites: Sprites{
			L: pixel.NewSprite(a.Sheets.Environment, a.Env[dir][0]),
			D: pixel.NewSprite(a.Sheets.Environment, a.Env[dir][1]),
		},
		Blocks:  true,
		Bitmask: mask,
	}
}

func openDoor(dir string, a *Assets, mask int) *Tile {
	return &Tile{
		Terrain: OpenDoor,
		Sprites: Sprites{
			L: pixel.NewSprite(a.Sheets.Environment, a.Env[dir][2]),
			D: pixel.NewSprite(a.Sheets.Environment, a.Env[dir][3]),
		},
		Blocks:  false,
		Bitmask: mask,
	}
}

//func openDoor(pos Position, g *Game) {
//	if g.Level.Tiles[pos.X][pos.Y].Bitmask == 0 {
//		g.Level.Tiles[pos.X][pos.Y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["door_w_e"][2])
//		g.Level.Tiles[pos.X][pos.Y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["door_w_e"][3])
//		g.Level.Tiles[pos.X][pos.Y].Terrain = OpenDoor
//		g.Level.Tiles[pos.X][pos.Y].Blocks = false
//	} else if g.Level.Tiles[pos.X][pos.Y].Bitmask == 1 {
//		g.Level.Tiles[pos.X][pos.Y].Sprites.L = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["door_n_s"][2])
//		g.Level.Tiles[pos.X][pos.Y].Sprites.D = pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["door_n_s"][3])
//		g.Level.Tiles[pos.X][pos.Y].Terrain = OpenDoor
//		g.Level.Tiles[pos.X][pos.Y].Blocks = false
//	}
//	g.Player.Actor.Fov.Block(pos.X, pos.Y, false)
//	g.Player.Actor.calculateFov()
//}
