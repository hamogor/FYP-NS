package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NumRooms int = 50

)


func generateLevel(g *Game) {
	dir := true
	//rooms := make([]Rectangle, 0)
	l := &Level{
		Tiles: [40][40]*Tile{},
		Spawn: Position{},
		Rooms: nil,
		Doors: nil,
	}
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			l.Tiles[x][y] = floor(0, g.Assets)
		}
	}
	for i := 0; i < NumRooms; i++ {
		splitRandom(l, g.Assets, dir)
		dir = !dir
	}


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

func splitRandom(l *Level, a *Assets, dir bool) {
	if dir {
		start := random(LevelW/2, LevelW)
		for y := 0; y < LevelH; y++ {
			l.Tiles[start][y] = wall(0, a)
		}
	} else if !dir {
		start := random(LevelH/2, LevelH)
		for x := 0; x < LevelW; x++ {
			l.Tiles[x][start] = wall(0, a)
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

/*
Need a recursive function that gets a split coordinate that doesnt already exist
A function to store the the left and right / north and south area of room produced
Place doors
Choose a room along the bottom x coord and set it as spawn
Pathfind to furthest room and set it as objective

*/
