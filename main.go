package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type Game struct {
	Render *Render
	Assets *Assets
	Level  *Level
	Player *Player
	Ai     *Ai
	Ui     *Ui
	Scenes *Scenes
}

func initialiseGame() *Game {
	g := &Game{}
	g.buildAssets()
	g.initRender()
	g.initPlayer()
	g.initLevel()
	g.initUi()
	generateLevel(g) //TESTING
	g.initData()
	g.initScenes()
	return g
}

func run() {
	g := initialiseGame()
	g.loop()
}

func (g *Game) loop() {
	for !g.Render.Window.Closed() {
		dt = time.Since(last).Seconds()
		last = time.Now()

		g.render()
		g.handleInput()
		frames++

		select {
		case <-second:
			g.Render.Window.SetTitle(fmt.Sprintf("%s | FPS: %d",
				g.Render.Config.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
