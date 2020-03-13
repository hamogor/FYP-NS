package ecs

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type Game struct {
	World *ecs.World
	Window *pixelgl.Window
	Config pixelgl.WindowConfig
	*AssetStore
}

func initGame() *Game {
	g := &Game{World: &ecs.World{}}
	g.AssetStore = InitAssetStore()
	g.World.AddSystem(InitRenderSystem(g))

	InitPlayer(g.World, g.AssetStore)

	return g
}

func run() {
	g := initGame()
	g.loop()
}

func (g *Game) loop() {
	for !g.Window.Closed() {
		dt = float32(time.Since(last).Seconds())
		last = time.Now()


		for i := range g.World.Systems() {
			g.World.Systems()[i].Update(dt)
		}

		frames++

		select {
		case <-second:
			g.Window.SetTitle(fmt.Sprintf("%s | FPS: %d",
				g.Config.Title, frames))
			frames = 0
		default:
		}
	}
}

//func main() {
//	pixelgl.Run(run)
//}