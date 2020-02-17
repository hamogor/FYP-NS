package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"log"
	"math"
)

type Render struct {
	Window *pixelgl.Window
	Config pixelgl.WindowConfig
	Env    Layer
	Actors Layer
	Camera pixel.Vec
	Ui     []*imdraw.IMDraw
	Mouse  pixel.Vec
}

type Layer struct {
	Batch  *pixel.Batch
	Canvas *pixelgl.Canvas
	Matrix pixel.Matrix
}

func (g *Game) initRender() {
	cfg := pixelgl.WindowConfig{
		Title:       "Null Spire",
		Icon:        nil,
		Bounds:      pixel.R(0, 0, WWidth, WHeight),
		Monitor:     nil,
		Resizable:   false,
		Undecorated: false,
		NoIconify:   false,
		AlwaysOnTop: false,
		VSync:       false,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Println(err)
	}

	r := &Render{
		Window: win,
		Config: cfg,
		Env: Layer{
			Batch:  pixel.NewBatch(&pixel.TrianglesData{}, g.Assets.Sheets.Environment),
			Canvas: pixelgl.NewCanvas(pixel.R(-WWidth/2, -WHeight/2, WWidth/2, WHeight/2)),
			Matrix: pixel.IM.Scaled(pixel.ZV, math.Min(Scaled, Scaled)).Moved(pixel.V(WWidth/2, WHeight/2)),
		},
		Actors: Layer{
			Batch:  pixel.NewBatch(&pixel.TrianglesData{}, g.Assets.Sheets.Sprites),
			Canvas: pixelgl.NewCanvas(pixel.R(-WWidth/2, -WHeight/2, WWidth/2, WHeight/2)),
			Matrix: pixel.IM.Scaled(pixel.ZV, math.Min(Scaled, Scaled)).Moved(pixel.V(WWidth/2, WHeight/2)),
		},
	}
	g.Render = r
	g.Render.Env.Canvas.SetMatrix(g.Render.Env.Matrix)
	g.Render.Actors.Canvas.SetMatrix(g.Render.Actors.Matrix)

}

func (g *Game) render() {
	g.Scenes.resetActiveElements()
	g.Scenes.setActiveUiElements()
	g.Render.Env.Batch.Clear()
	g.Render.Actors.Batch.Clear()
	g.Render.Env.Canvas.Clear(color.RGBA{R: 8, G: 8, B: 12, A: 255,})
	g.Render.Actors.Canvas.Clear(color.Transparent)
	g.Render.Ui = g.Render.Ui[:0]



	g.Render.Camera = pixel.Lerp(g.Render.Camera, g.Player.Actor.Pos.sToVec(), 1-math.Pow(1.0/128, dt))
	cam := pixel.IM.Moved(pixel.V(RoundFloat(g.Render.Camera.X), RoundFloat(g.Render.Camera.Y)).Scaled(-1))


	g.Render.Env.Batch.SetMatrix(cam)
	g.Render.Actors.Batch.SetMatrix(cam)

	g.Render.renderEnvironment(g.Level, g.Player, g.Scenes)
	g.Render.renderActors(g.Player, g.Level, g.Assets, g.Scenes)

	g.Render.Env.Batch.Draw(g.Render.Env.Canvas)
	g.Render.Actors.Batch.Draw(g.Render.Actors.Canvas)

	canvasMatrix := pixel.IM.Scaled(pixel.ZV, math.Min(Scaled, Scaled)).Moved(g.Render.Env.Canvas.Bounds().Center())
	g.Render.Env.Canvas.SetMatrix(canvasMatrix)
	g.Render.Actors.Canvas.SetMatrix(canvasMatrix)

	worldMatrix := pixel.IM.Moved(g.Render.Window.Bounds().Center())
	g.Render.Env.Canvas.Draw(g.Render.Window, worldMatrix)
	g.Render.Actors.Canvas.Draw(g.Render.Window, worldMatrix)

	g.Render.renderUi(g.Player, g.Ui, g.Scenes)
	g.Render.Window.Update()
}

func (r *Render) renderActors(p *Player, l *Level, a *Assets, s *Scenes) {
	if s.CurrentScene == GameScene {
		p.Actor.updateAnimState()
		p.Actor.CAnim.Sprite.Set(a.Sheets.Sprites, p.Actor.CAnim.Frame.Rect)
		p.Actor.CAnim.Sprite.Draw(r.Actors.Batch, pixel.IM.ScaledXY(pixel.ZV, pixel.V(-p.Actor.Direction, 1)).Moved(p.Actor.Pos.sToVec()))
	}

}

func (r *Render) renderEnvironment(l *Level, p *Player, s *Scenes) {
	if s.CurrentScene == GameScene {
		for x := 0; x < LevelW; x++ {
			for y := 0; y < LevelH; y++ {
				if p.Actor.Fov.Look(x, y) {
					l.Tiles[x][y].Sprites.L.Draw(r.Env.Batch, tilePos(x, y))
					p.Actor.Fov.explored[x][y] = true
				} else if p.Actor.Fov.explored[x][y] {
					l.Tiles[x][y].Sprites.D.Draw(r.Env.Batch, tilePos(x, y))
				}

			}
		}
	}

}

func (r *Render) renderUi(p *Player, ui *Ui, s *Scenes) {
		r.renderMiniMap(p, ui, s)
		r.renderBar(ui, s)
		r.renderPortrait(ui, s) // Top of bar
		r.renderHealthBar(ui, p, s)
		r.renderMainMenu(ui, s)
}

func (r *Render) renderMiniMap(p *Player, ui *Ui, s *Scenes) {
	if s.ActiveElements[MiniMapActive] {
		tr := pixel.V(WWidth-(LevelW*2), WHeight-(LevelH*2))
		ui.MiniMap.Sprite.Draw(r.Window, pixel.IM.Moved(tr).Scaled(pixel.V(tr.X+3, tr.Y+3), math.Min(4, 4)))
		ui.MiniMap.Msprite.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, math.Min(3, 3)).ScaledXY(pixel.ZV, pixel.V(1, -1)).Moved(tr))
	}
}

func (r *Render) renderBar(ui *Ui, s *Scenes) {
	if s.ActiveElements[MenuBarActive] {
		ui.MenuBar.LSprite.Draw(r.Window, pixel.IM.Moved(pixel.V(ui.MenuBar.LSprite.Frame().Max.X, ui.MenuBar.LSprite.Frame().Max.Y/2)))
		ui.MenuBar.RSprite.Draw(r.Window, pixel.IM.Moved(pixel.V(WWidth-ui.MenuBar.RSprite.Frame().Max.X, ui.MenuBar.RSprite.Frame().Max.Y/2)))
		mat := pixel.IM
		mat = mat.ScaledXY(pixel.ZV, pixel.V(WWidth-(ui.MenuBar.LSprite.Frame().Max.X+ui.MenuBar.RSprite.Frame().Max.X)-10, 1))
		mat = mat.Moved(pixel.V(WWidth/2, ui.MenuBar.Sprite.Frame().Max.Y/2))
		ui.MenuBar.Sprite.Draw(r.Window, mat)
	}

}

func (r *Render) renderPortrait(ui *Ui, s *Scenes) {
	if s.ActiveElements[PortraitActive] {
		blOfBar := pixel.V(ui.MenuBar.LSprite.Frame().Max.X+ui.Portrait.Sprite.Frame().Max.X+10, ui.MenuBar.LSprite.Frame().Max.Y/2)
		ui.Portrait.Sprite.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, math.Min(2, 2)).Moved(blOfBar))
	}
}

func (r *Render) renderHealthBar(ui *Ui, p *Player, s *Scenes) {
	if s.ActiveElements[PlayerHealthActive] {
		red := imdraw.New(nil)
		red.Clear()

		red.Color = colornames.Darkred
		red.Push(pixel.V(100, 60))
		red.Push(pixel.V(300, 70))
		red.Rectangle(0)
		red.Draw(r.Window)

		green := imdraw.New(nil)
		green.Clear()
		percentHP := float64(p.Actor.HP) / float64(100) * 200
		green.Color = colornames.Forestgreen
		green.Push(pixel.V(100, 60))
		green.Push(pixel.V(100+percentHP, 70))
		green.Rectangle(0)
		green.Draw(r.Window)
	}

}

func (r *Render) renderMainMenu(ui *Ui, s *Scenes) {
	if s.ActiveElements[MainMenuActive] {
		w, h := WWidth/ui.MainMenu.Background.Frame().Max.X, WHeight/ui.MainMenu.Background.Frame().Max.Y
		ui.MainMenu.Background.Draw(r.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(w, h)).Moved(r.Window.Bounds().Center()))
		center := r.Window.MousePosition()
		center.X = center.X - WWidth * 0.5
		center.Y = center.Y - WHeight * 0.5
		fmt.Print(center, "\n")
		if ui.MainMenu.StartButton.Hovering {
			ui.MainMenu.StartButton.HSprite.Draw(r.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(w, h)).Scaled(pixel.ZV, 2).Moved(pixel.V(380, 220)))
		} else {
			ui.MainMenu.StartButton.Sprite.Draw(r.Window, pixel.IM.ScaledXY(pixel.ZV, pixel.V(w, h)).Scaled(pixel.ZV, 2).Moved(pixel.V(380, 220)))
		}
		ScreenToWorldSpace(r, pixel.IM)

	}
}

func ScreenToWorldSpace(r *Render, mat pixel.Matrix) pixel.Vec {
	return mat.Project(r.Window.MousePosition())
}