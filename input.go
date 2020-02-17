package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func (g *Game) handleInput() {
	g.Player.Actor.ActionTaken = false

	if g.Scenes.CurrentScene == GameScene {
		if g.Render.Window.JustPressed(pixelgl.KeyW) || g.Render.Window.Repeated(pixelgl.KeyW) {
			g.Player.Actor.move(g.Player.Actor.Pos.N(), g)
			g.Player.Actor.ActionTaken = true
		}

		if g.Render.Window.JustPressed(pixelgl.KeyS) || g.Render.Window.Repeated(pixelgl.KeyS) {
			g.Player.Actor.move(g.Player.Actor.Pos.S(), g)
			g.Player.Actor.ActionTaken = true
		}

		if g.Render.Window.JustPressed(pixelgl.KeyA) || g.Render.Window.Repeated(pixelgl.KeyA) {
			g.Player.Actor.move(g.Player.Actor.Pos.W(), g)
			g.Player.Actor.ActionTaken = true
		}

		if g.Render.Window.JustPressed(pixelgl.KeyD) || g.Render.Window.Repeated(pixelgl.KeyD) {
			g.Player.Actor.move(g.Player.Actor.Pos.E(), g)
			g.Player.Actor.ActionTaken = true
		}

		if g.Render.Window.JustPressed(pixelgl.KeyQ) || g.Render.Window.Repeated(pixelgl.KeyQ) {
			g.Player.Actor.move(g.Player.Actor.Pos.NW(), g)
			g.Player.Actor.ActionTaken = true
		}

		if g.Render.Window.JustPressed(pixelgl.KeyE) || g.Render.Window.Repeated(pixelgl.KeyE) {
			g.Player.Actor.move(g.Player.Actor.Pos.NE(), g)
			g.Player.Actor.ActionTaken = true
		}

		if g.Render.Window.JustPressed(pixelgl.KeyZ) || g.Render.Window.Repeated(pixelgl.KeyZ) {
			g.Player.Actor.move(g.Player.Actor.Pos.SW(), g)
			g.Player.Actor.ActionTaken = true
		}

		if g.Render.Window.JustPressed(pixelgl.KeyC) || g.Render.Window.Repeated(pixelgl.KeyC) {
			g.Player.Actor.move(g.Player.Actor.Pos.SE(), g)
			g.Player.Actor.ActionTaken = true
		}
	}


	if g.Scenes.CurrentScene == MainMenuScene {
		mouse := g.Render.Window.MousePosition()
		if mouse.X > 163 && mouse.X < 589 &&
			mouse.Y > 181 && mouse.Y < 275 {
			g.Ui.MainMenu.StartButton.Hovering = true
		} else {
			g.Ui.MainMenu.StartButton.Hovering = false
		}
		if g.Ui.MainMenu.StartButton.Hovering && g.Render.Window.JustPressed(pixelgl.MouseButton1) {
			startButton(g.Scenes)
		}
	}



	if g.Player.Actor.ActionTaken {
		g.Player.Actor.calculateFov()
		g.Player.updateMiniMap(g.Level, g.Ui)
	}

}
