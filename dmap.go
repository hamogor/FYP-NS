package main // Package dmap implements a Brogue-style Dijkstra Map data structure
import (
	"bytes"
	"fmt"
	"math"
)

// based on the map it is given. For more information on this data
// structure, read the article on RogueBasin:
// http://www.roguebasin.com/index.php?title=The_Incredible_Power_of_Dijkstra_Maps

// Point is a representation of a point on a map
type Point interface {
	GetXY() (int, int)
}

// Map is a map to calculate dmaps with. All methods should be linear
// time, as the algorithm will call them a lot. Also your map should
// be statically sized; if the map is dynamically sized, a new dmap
// will need to be created for it whenever it changes size.
type Map interface {
	SizeX() int
	SizeY() int
	IsPassable(int, int) bool
	OOB(int, int) bool
}

// Rank is the rank of a tile - lower is closer to the target
type Rank uint16

func (r Rank) GetXY() (int, int) {
	return r.GetXY()
}

// RankMax is the highest possible rank. It takes a value a little
// below its implementations maximum to prevent overflow
const RankMax = math.MaxUint16 - 10

// DijkstraMap is a representation of a Brogue-style 'Dijkstra'
// map. To reach a target, an AI should try to minimize the rank of
// the tile it's standing on (targets have a value of zero)
type DijkstraMap struct {
	Points       [][]Rank
	M            Map
	NeigbourFunc func(d *DijkstraMap, x, y int) []WeightedPoint
}

// WeightedPoint is a Point that also has a rank
type WeightedPoint struct {
	X   int
	Y   int
	Val Rank
}

// GetXY implements the Point interface
func (d *WeightedPoint) GetXY() (int, int) {
	return d.X, d.Y
}

// BlankDMap creates a blank Dijkstra map to be used with the map passed to it
func BlankDMap(m Map, neigbourfunc func(d *DijkstraMap, x, y int) []WeightedPoint) *DijkstraMap {
	ret := make([][]Rank, m.SizeX())
	for i := range ret {
		ret[i] = make([]Rank, m.SizeY())
		for j := range ret[i] {
			ret[i][j] = RankMax
		}
	}
	return &DijkstraMap{ret, m, neigbourfunc}
}

// ManhattanNeighbours returns the neighbours of the block x, y to the
// north, south, east, and west
func ManhattanNeighbours(d *DijkstraMap, x, y int) []WeightedPoint {
	return []WeightedPoint{
		d.GetValPoint(x+1, y),
		d.GetValPoint(x-1, y),
		d.GetValPoint(x, y-1),
		d.GetValPoint(x, y+1),
	}
}

// DiagonalNeighbours returns the neighbours of the block x, y to the
// north, south, east, west, NE, SE, NW, and SW
func DiagonalNeighbours(d *DijkstraMap, x, y int) []WeightedPoint {
	return []WeightedPoint{
		d.GetValPoint(x+1, y),
		d.GetValPoint(x-1, y),
		d.GetValPoint(x, y-1),
		d.GetValPoint(x, y+1),
		d.GetValPoint(x+1, y+1),
		d.GetValPoint(x+1, y-1),
		d.GetValPoint(x-1, y+1),
		d.GetValPoint(x-1, y-1),
	}
}

// Calc calculates the Dijkstra map with points given as targets. You
// need to blank the map before using this method. It's recommended to
// use this one initially, but to use a Recalc instead for subsequent
// moves, since Recalc, unlike BlankDMap, doesn't allocate memory.
func (d *DijkstraMap) Calc(points ...Point) {
	for _, point := range points {
		x, y := point.GetXY()
		if x <= LevelW || y <= LevelH || x < 0 || y < 0 {

		}
		d.Points[x][y] = 0
	}
	mademutation := true
	for mademutation {
		mademutation = false
		for x := range d.Points {
			for y := range d.Points[x] {
				if d.M.IsPassable(x, y) {
					ln := d.LowestNeighbour(x, y).Val
					if d.Points[x][y] > ln+1 {
						d.Points[x][y] = ln + 1
						mademutation = true
					}
				}
				x1, y1 := (d.M.SizeX()-1)-x, (d.M.SizeY()-1)-y
				if d.M.IsPassable(x1, y1) {
					ln := d.LowestNeighbour(x1, y1).Val
					if d.Points[x1][y1] > ln+1 {
						d.Points[x1][y1] = ln + 1
						mademutation = true
					}
				}
			}
		}
	}
}

// Recalc recalculates the Dijkstra map with points given as
// targets. It's essentially equivalent to a blank followed by a calc,
// but should be a bit faster because it doesn't reallocate the
// memory. As per the note for DijkstraMap, don't use this method if
// your map is dynamically sized; you'll just have to use BlankDMap
// and Calc as if creating a new dmap every update.
func (d *DijkstraMap) Recalc(points ...Point) {
	for i := range d.Points {
		for j := range d.Points[i] {
			d.Points[i][j] = RankMax
		}
	}
	d.Calc(points...)
}

// GetValPoint gets the weighted point at X, Y of the Dijkstra
// map. Points that are out of bounds count as maximum rank (so
// shouldn't be targeted)
func (d *DijkstraMap) GetValPoint(x, y int) WeightedPoint {
	if d.M.OOB(x, y) {
		return WeightedPoint{x, y, RankMax}
	}
	if x >= LevelW || y >= LevelH || x < 0 || y < 0 {
		return WeightedPoint{
			X:   x,
			Y:   y,
			Val: RankMax,
		}
	}
	return WeightedPoint{x, y, d.Points[x][y]}
}

func (d *DijkstraMap) SetValPoint(x, y, val int) {
	d.Points[x][y] = Rank(val)
}

func (d *DijkstraMap) getRankAsUInt8(x, y int) uint8 {
	if x <= 19 && y <= 19 && x > 0 && y > 0 {
		return uint8(d.Points[x][y])
	} else {
		return 100
	}

}

// LowestNeighbour returns the neighbour of the point at x, y with the
// lowest rank.
func (d *DijkstraMap) LowestNeighbour(x, y int) WeightedPoint {
	vals := d.NeigbourFunc(d, x, y)
	var lv Rank = RankMax
	ret := vals[0]
	for _, val := range vals {
		if val.Val < lv {
			lv = val.Val
			ret = val
		}
	}
	return ret
}



func (d *DijkstraMap) GetLowestValue(l *Level) Rank {
	lv := Rank(RankMax)
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if d.M.IsPassable(x, y) {
				if d.Points[x][y] < lv {
					lv = d.Points[x][y]
				}
			}
		}
	}
	return lv
}

func (d *DijkstraMap) GetLowestValuePos(l *Level) Position {
	pos := Position{X: 0, Y: 0}
	var lv = Rank(RankMax)
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if d.M.IsPassable(x, y) {
				if d.Points[x][y] < lv {
					lv = d.Points[x][y]
					pos = Position{X: x, Y: y}
				}
			}
		}
	}
	return pos
}

// String returns a string representation of a Dijkstra Map
func (d *DijkstraMap) String() string {
	buf := bytes.Buffer{}
	for x := range d.Points {
		for y := range d.Points[x] {
			buf.WriteString(fmt.Sprintf("%6d", d.Points[x][y]))
			buf.WriteString(", ")
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

func DiagonalNeighboursAtRadius2(d *DijkstraMap, x, y int) []WeightedPoint {
	return []WeightedPoint{
		d.GetValPoint(x-2, y+2),
		d.GetValPoint(x-1, y+2),
		d.GetValPoint(x, y+2),
		d.GetValPoint(x+1, y+2),
		d.GetValPoint(x+2, y+2),
		d.GetValPoint(x-2, y+1),
		d.GetValPoint(x-1, y+1),
		d.GetValPoint(x, y+1),
		d.GetValPoint(x+1, y+1),
		d.GetValPoint(x+2, y+1),
		d.GetValPoint(x-2, y),
		d.GetValPoint(x-1, y),
		d.GetValPoint(x+1, y),
		d.GetValPoint(x+2, y),
		d.GetValPoint(x-2, y-1),
		d.GetValPoint(x-1, y-1),
		d.GetValPoint(x, y-1),
		d.GetValPoint(x+1, y-1),
		d.GetValPoint(x+2, y-1),
		d.GetValPoint(x-2, y-2),
		d.GetValPoint(x-1, y-2),
		d.GetValPoint(x, y-2),
		d.GetValPoint(x+1, y-2),
		d.GetValPoint(x+2, y-2),
	}
}

func DiagonalNeighboursAtRadius4(x, y int) []Position {
	return []Position{
		{X: x - 4, Y: y + 4},
		{X: x-3, Y: y+4},
		{X: x-2, Y: y+4},
		{X: x-1, Y: y+4},
		{X: x, Y: y+4},
		{X: x+1, Y: y+4},
		{X: x+2, Y: y+4},
		{X: x+3, Y: y+4},
		{X: x+4, Y: y+4},

		{X: x+4, Y: y+3},
		{X: x+4, Y: y+2},
		{X: x+4, Y: y+1},
		{X: x+4, Y: y},
		{X: x+4, Y: y-1},
		{X: x+4, Y: y-2},
		{X: x+4, Y: y-3},
		{X: x+4, Y: y-4},

		{X: x+3, Y: y-4},
		{X: x+2, Y: y-4},
		{X: x+1, Y: y-4},
		{X: x, Y: y-4},
		{X: x-1, Y: y+4},
		{X: x-2, Y: y+4},
		{X: x-3, Y: y+4},
		{X: x-4, Y: y+4},
		{X: x-4, Y: y-3},
		{X: x-4, Y: y-2},
		{X: x-4, Y: y-1},
		{X: x-4, Y: y},
		{X: x-4, Y: y+1},
		{X: x-4, Y: y+2},
		{X: x-4, Y: y+3},

		{X: x-4, Y: y-4},
		{X: x-3, Y: y-4},
		{X: x-2, Y: y-4},
		{X: x-1, Y: y-4},
	}
}

func (d *DijkstraMap) CalculateFleeMap(p *Player, l *Level) {

	for x := range d.Points {
		for y := range d.Points[x] {
			if d.M.IsPassable(x, y) {
				d.Points[x][y] = 50 - d.Points[x][y]
			}
		}
	}

	lvPos := d.GetLowestValuePos(l)
	d.Calc(lvPos)
}




func (d *DijkstraMap) copyDmap(l *Level, a *Actor) *DijkstraMap {
	copyDmap := &DijkstraMap{
		Points:       d.Points,
		M:            d.M,
		NeigbourFunc: d.NeigbourFunc,
	}
	return copyDmap
}

func (d *DijkstraMap) MutateActorPositions(l *Level, j int) {
	for i := range l.Actors {
		if i == j {
			break
		} else if l.Actors[i].State != PlayerAi {
			d.Points[l.Actors[i].Pos.X][l.Actors[i].Pos.Y] = RankMax
		}
	}
}


func (d *DijkstraMap) CalculateRangeMap(p *Player, l *Level) {
	neighbours := DiagonalNeighboursAtRadius4(p.Actor.Pos.X, p.Actor.Pos.Y)
	points := make([]Point, 0)
	for i := range neighbours {
		if neighbours[i].X > 0 && neighbours[i].X < LevelW && neighbours[i].Y > 0 && neighbours[i].Y < LevelH {
			if l.Tiles[neighbours[i].X][neighbours[i].Y].Terrain == Floor ||
				l.Tiles[neighbours[i].X][neighbours[i].Y].Terrain == OpenDoor {
				points = append(points, neighbours[i])

			}
		}

	}
	d.Recalc(points...)
}

func (d *DijkstraMap) CalculateFlankMap(p *Player, l *Level) {
	alreadyIncreased := make([]Position, 0)
	for i := range l.Actors {
		if l.Actors[i].State == FlankAi {
			neighbours := DiagonalNeighboursAtRadius2(d, l.Actors[i].Pos.X, l.Actors[i].Pos.Y)
			for j := range neighbours {
				if neighbours[j].X < LevelW && neighbours[j].Y < LevelH && neighbours[j].X > 0 && neighbours[j].Y > 0 {
					for k := range alreadyIncreased {
						if alreadyIncreased[k].X == neighbours[j].X && alreadyIncreased[k].Y == neighbours[j].Y {
							break
						} else {
							d.Points[neighbours[j].X][neighbours[j].Y] += 3
							alreadyIncreased = append(alreadyIncreased, Position{X: neighbours[j].X, Y: neighbours[j].Y})
						}
					}

				}
			}
			//d.Points[l.Actors[i].Pos.X][l.Actors[i].Pos.Y] += 2
		}
		//if l.Actors[i].State == FlankAi {
		//	neighbours := DiagonalNeighbours(l.Actors[i].DMap, l.Actors[i].Pos.X, l.Actors[i].Pos.Y)
		//	for j := range neighbours {
		//		if neighbours[j].X < LevelW && neighbours[j].Y < LevelH && neighbours[j].X > 0 && neighbours[j].Y > 0 {
		//			l.Actors[i].DMap.Points[neighbours[j].X][neighbours[j].Y] += 3
		//		}
		//	}
		//}

	}
}


