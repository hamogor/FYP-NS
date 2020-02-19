package main

import "fmt"

type Room struct {
	Area Rectangle
	x1, y1, x2, y2 int
	Door Position
}

func generateLevel(g *Game) {
	l := &Level{
		Tiles: [64][64]*Tile{},
		Spawn: Position{X: 1, Y: 1},
		Rooms: nil,
		Doors: nil,
	}
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			l.Tiles[x][y] = tile()
			if x == 0 || x == LevelW || x == LevelW - 1 {
				l.Tiles[x][y].Terrain = Wall
			}
			if y == 0 || y == LevelH || y == LevelH - 1 {
				l.Tiles[x][y].Terrain = Wall
			}
		}
	}
	l.gen()


	l.applyBitmask(g.Assets)
	l.print()
	g.Level = l
}

func (l *Level) gen() {

	for i := 0; i < 1; i++ {
		w := random(8, 14)
		h := random(3, w / 2)
		x := random(0, LevelW - w - 1)
		y := random(0, 1)
		newRoom := &Room{
			Area: Rectangle{x, y, w, h},
			x1:   x,
			y1:   y,
			x2:   x + w,
			y2:   y + h,
		}
		l.drawRoom(newRoom)
		fmt.Print(newRoom, "\n")
		if i == 0 {
			l.Spawn = newRoom.Area.center()
		}
	}

}

func (l *Level) drawRoom(room *Room) {
	for x := room.x1; x < room.x2 + 1; x++ {
		for y := room.y1; y < room.y2 + 1; y++ {
			if x == room.Area.X || x == room.Area.Width + room.Area.X {
				l.Tiles[x][y].Terrain = Wall
			}
			if y == room.Area.Y || y == room.Area.Height + room.Area.Y {
				l.Tiles[x][y].Terrain = Wall
			}
			fmt.Print(room.x2 / 2, " ", room.y2, "\n")
			l.Tiles[room.x2][room.y2/2].Terrain = Door
			l.Doors = append(l.Doors, Position{X: room.x2, Y: room.y2/2})
		}
	}
}

func (l *Level) applyBitmask(a *Assets) {
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			pos := Position{X: x, Y: y}

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