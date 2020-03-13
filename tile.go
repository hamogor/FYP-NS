package NS

import "github.com/faiface/pixel"

type Tile struct {
	Terrain Terrain
	Sprites Sprites
	Blocks  bool
	Bitmask int
}

type TileAction func(l *Level, pos Position)

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
	Free
)

func (pos Position) terrain(l *Level) Terrain {
	if pos.inLevel() {
		return l.Tiles[pos.X][pos.Y].Terrain
	} else {
		return Free
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
		Terrain: Floor,
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

func (pos Position) resolveDoorTypeAndOpen(l *Level, a *Assets) {
	if l.Tiles[pos.X][pos.Y].Bitmask == 0 {
		l.Tiles[pos.X][pos.Y] = openDoor("door_w_e", a, 0)
	} else if l.Tiles[pos.X][pos.Y].Bitmask == 1 {
		l.Tiles[pos.X][pos.Y] = openDoor("door_n_s", a, 1)
	}
}
