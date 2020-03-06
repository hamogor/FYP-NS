package main

import (
	"fmt"
	"github.com/EngoEngine/ecs"
)

func initWorld() ecs.World {
	// Rewrite the whole thing with ecs
	world := ecs.World{}
	fmt.Print(world.Systems(), "\n")
}
