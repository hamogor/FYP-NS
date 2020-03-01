package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NumRooms    int = 50
	MaxNodeSize int = 8
	MinNodeSize int = 2
)

type tree []*BSPNode

type BSPNode struct {
	Area       Rectangle
	Hall       Rectangle
	Left       *BSPNode
	Right      *BSPNode
	X, Y, W, H int
}

func generateLevelold(g *Game) {
	nodes := make(tree, 0)
	l := &Level{
		Tiles: [64][64]*Tile{},
		Spawn: Position{},
		Rooms: nil,
		Doors: nil,
	}
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			l.Tiles[x][y] = floor(0, g.Assets)
		}
	}

	rootNode := newNodeOld(0, 0, LevelW+2, LevelH+2)
	nodes = append(nodes, rootNode)
	split := true

	for i := 0; i < len(nodes); i++ {
		if split {
			split = false
			if nodes[i].Left == nil && nodes[i].Right == nil {
				if nodes[i].W > MaxNodeSize || nodes[i].H > MaxNodeSize {
					nodes[i].split(&nodes)
					split = true
				}
			}
		} else {
			break
		}
	}
	nodes[0].createRooms(nodes, l)

	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if x == 0 || x == LevelW || x == LevelW-1 {
				l.Tiles[x][y].Terrain = Wall
			}
			if y == 0 || y == LevelH || y == LevelH-1 {
				l.Tiles[x][y].Terrain = Wall
			}
			pos := Position{X: x, Y: y}
			if pos.terrain(l) == Wall {
				mask := BoolListToMask(pos.ResolveBitMaskWall(l))
				l.Tiles[x][y] = wall(mask, g.Assets)
			} else if pos.terrain(l) == Floor {
				mask := BoolListToMask(pos.ResolveBitMaskFloor(l))
				l.Tiles[x][y] = floor(mask, g.Assets)
			}
		}
	}

	l.print()
	l.Spawn = Position{X: 1, Y: 1}
	g.Level = l
}

func (node *BSPNode) split(tree *tree) bool {
	if node.Left != nil || node.Right != nil {
		return false // Already split
	}
	splitHorizontally := randomBool()
	if float64(node.W/node.H) >= 1.25 {
		splitHorizontally = false
	} else if float64(node.H/node.W) >= 1.25 {
		splitHorizontally = true
	}
	maxS := 0
	if splitHorizontally {
		maxS = node.H - MinNodeSize
	} else {
		maxS = node.W - MinNodeSize
	}

	if maxS <= MinNodeSize {
		return false // Too small to split further
	}

	split := random(MinNodeSize, maxS)

	if splitHorizontally {
		node.Left = newNodeOld(node.X, node.Y, node.W, split)
		node.Right = newNodeOld(node.X, node.Y+split, node.W, node.H-split)
	} else {
		node.Left = newNodeOld(node.X, node.Y, split, node.H)
		node.Right = newNodeOld(node.X+split, node.Y, node.W-split, node.H)
	}
	*tree = append(*tree, node.Left, node.Right)
	return true
}

func (node *BSPNode) createRooms(tree tree, l *Level) {
	if node.Left != nil || node.Right != nil {
		if node.Left != nil {
			node.Left.createRooms(tree, l)
		}
		if node.Right != nil {
			node.Right.createRooms(tree, l)
		}
		if node.Left != nil && node.Right != nil {
			tree.createWall(node.Left.room(), node.Right.room(), l)
		} else {
			w := random(MinNodeSize, min(MaxNodeSize, node.W-1))
			h := random(MinNodeSize, min(MaxNodeSize, node.H-1))
			x := random(node.X, node.X+(node.H-1)-node.W)
			y := random(node.Y, node.Y+(node.H-1)-node.H)
			node.Area = Rectangle{
				X:      x,
				Y:      y,
				Width:  w,
				Height: h,
			}
			tree.createRoom(node.Area, l)
		}
	}
}

func (tree tree) createWall(left, right Rectangle, l *Level) {
	pos1 := left.center()
	pos2 := right.center()
	if random(0, 1) == 1 {
		tree.createHorWall(pos1.X, pos2.X, pos1.Y, l)
		tree.createVirWall(pos1.Y, pos2.Y, pos2.X, l)
	} else {
		tree.createVirWall(pos1.Y, pos2.Y, pos1.X, l)
		tree.createHorWall(pos1.X, pos2.X, pos2.Y, l)

	}
}

func (tree tree) createHorWall(x1, x2, y int, l *Level) {
	minS, maxS := min(x1, x2), max(x1, x2)+1
	for x := minS; x < maxS; x++ {
		l.Tiles[x][y].Terrain = Wall
	}
}

func (tree tree) createVirWall(y1, y2, x int, l *Level) {
	minS, maxS := min(y1, y2), max(y1, y2)+1
	for y := minS; y < maxS; y++ {
		l.Tiles[x][y].Terrain = Wall
	}
}

func (tree tree) createRoom(room Rectangle, l *Level) {
	for x := room.X + 1; x < room.Width; x++ {
		for y := room.Y + 1; y < room.Height; y++ {
			l.Tiles[x][y].Terrain = Floor
		}
	}
}

func randomBool() bool {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(2) != 0
}

func (node BSPNode) room() Rectangle {
	if node.Area.Width != 0 {
		return node.Area
	} else {
		if node.Left != nil {
			node.Left.Area = node.Left.room()
		}
		if node.Right != nil {
			node.Right.Area = node.Right.room()
		}
		if node.Left == nil && node.Right == nil {
			return Rectangle{}
		} else if node.Right.Area.Width != 0 {
			return node.Left.Area
		} else if node.Left.Area.Width != 0 {
			return node.Left.Area
		} else if random(0, 1) > 0 {
			return node.Left.Area
		} else {
			return node.Right.Area
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func newNodeOld(x, y, w, h int) *BSPNode {
	return &BSPNode{
		Area: Rectangle{
			X:      x,
			Y:      y,
			Width:  w,
			Height: h,
		},
		Hall:  Rectangle{},
		Left:  nil,
		Right: nil,
		W:     w,
		H:     h,
		X:     x,
		Y:     y,
	}
}

func (l *Level) print() {
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if l.Tiles[x][y].Terrain == Wall {
				fmt.Print("#")
			} else if l.Tiles[x][y].Terrain == Floor {
				fmt.Print(".")
			} else if l.Tiles[x][y].Terrain == Door {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

/*
Need a recursive function that gets a split coordinate that doesnt already exist
A function to store the the left and right / north and south area of room produced
Place doors
Choose a room along the bottom x coord and set it as spawn
Pathfind to furthest room and set it as objective

Fill level with floors
For range numrooms, make a split that doesn't already exist
store the left and right / north and south areas as children.

*/

//func splitRandom(l *Level, a *Assets, dir bool, nodes []*BSPNode) BSPNode {
//	if dir { //Vertical
//		start := random(LevelW/2, LevelW)
//		for y := 0; y < LevelH; y++ {
//			l.Tiles[start][y] = wall(0, a)
//		}
//	} else if !dir { // Horizontal
//		start := random(LevelH/2, LevelH)
//		for x := 0; x < LevelW; x++ {
//			l.Tiles[x][start] = wall(0, a)
//		}
//	}
//}
