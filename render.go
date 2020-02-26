package main

import (
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
		Undecorated: true,
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
		//r.renderMiniMap(p, ui, s)
		//r.renderBar(ui, s)
		//r.renderPortrait(ui, s) // Top of bar
		//r.renderHealthBar(ui, p, s)
		r.renderMainMenu(ui, s)
		r.renderTest9Slice(ui, s)
}

func (r *Render) renderMiniMap(p *Player, ui *Ui, s *Scenes) {
	if s.ActiveElements[MiniMapActive] {
		//tr := pixel.V(WWidth-(LevelW*2), WHeight-(LevelH*2))
		//scale := (WWidth / ui.MiniMap.Sprite.Frame().Max.X) / 8
		tr := pixel.V(WWidth - (ui.MiniMap.Sprite.Frame().Max.X / 2), WHeight - (ui.MiniMap.Sprite.Frame().Max.Y / 2))
		ui.MiniMap.Sprite.Draw(r.Window, pixel.IM.Moved(tr))
		//ui.MiniMap.Sprite.Draw(r.Window, pixel.IM.Moved(tr).Scaled(pixel.V(tr.X+3, tr.Y+3), math.Min(4, 4)))
		//ui.MiniMap.Msprite.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, math.Min(2, 2)).ScaledXY(pixel.ZV, pixel.V(1, -1)).Moved(tr))
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
		mat := pixel.IM.ScaledXY(pixel.ZV, pixel.V(w, h))
		screenRect(mat, ui.MainMenu.StartButton.Sprite, r)
		ui.MainMenu.Logo.Draw(r.Window, anchorTL(mat, ui.MainMenu.Logo, percentW(25), percentH(20), 5, &ui.MainMenu.StartButton))

		if ui.MainMenu.StartButton.Hovering {
			ui.MainMenu.StartButton.HSprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.StartButton.HSprite, percentW(23), percentH(45), 2, &ui.MainMenu.StartButton))
		} else {
			ui.MainMenu.StartButton.Sprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.StartButton.Sprite, percentW(23), percentH(45), 2, &ui.MainMenu.StartButton))
		}

		if ui.MainMenu.OptionsButton.Hovering {
			ui.MainMenu.OptionsButton.HSprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.OptionsButton.Sprite, percentW(23), percentH(50), 2, &ui.MainMenu.OptionsButton))
		} else {
			ui.MainMenu.OptionsButton.Sprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.OptionsButton.Sprite, percentW(23), percentH(50), 2, &ui.MainMenu.OptionsButton))
		}

		if ui.MainMenu.AboutButton.Hovering {
			ui.MainMenu.AboutButton.HSprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.AboutButton.Sprite, percentW(23), percentH(55), 2, &ui.MainMenu.AboutButton))
		} else {
			ui.MainMenu.AboutButton.Sprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.AboutButton.Sprite, percentW(23), percentH(55), 2, &ui.MainMenu.AboutButton))
		}

		if ui.MainMenu.ExitButton.Hovering {
			ui.MainMenu.ExitButton.HSprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.ExitButton.Sprite, percentW(23), percentH(60), 2, &ui.MainMenu.ExitButton))
		} else {
			ui.MainMenu.ExitButton.Sprite.Draw(r.Window, anchorTL(mat, ui.MainMenu.ExitButton.Sprite, percentW(23), percentH(60), 2, &ui.MainMenu.ExitButton))
		}


	}
}

func (r *Render) renderTest9Slice(ui *Ui, s *Scenes) {
	if s.CurrentScene == GameScene {
		scale := (WWidth / ui.Test.Objectives.Frame().Max.X) / 8
		tr := pixel.V(WWidth - (ui.Test.Objectives.Frame().Max.X / 2) * scale - 4, WHeight - (ui.Test.Objectives.Frame().Max.Y / 2) * scale - 4)
		ui.Test.Objectives.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, scale).Moved(tr))

		belowTR := pixel.V(WWidth - (ui.Test.Look.Frame().Max.X / 2) * scale - 4, WHeight - ((ui.Test.Objectives.Frame().Max.Y) + (ui.Test.Look.Frame().Max.Y / 2)) * scale - 10)
		ui.Test.Look.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, scale).Moved(belowTR))

		bm := pixel.V(WWidth/2, (ui.Test.Bar.Frame().Max.Y / 2) * scale + 4)
		ui.Test.Bar.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, scale).Moved(bm))

		tl := pixel.V((ui.Test.HBar.Frame().Max.X/2) * scale + 4, WHeight - (ui.Test.HBar.Frame().Max.Y /2) * scale - 4)
		ui.Test.HBar.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, scale).Moved(tl))

		buttonPos := pixel.V((ui.Test.Buttons.Frame().Max.X / 2) * scale + 4, WHeight / 3 + (WHeight/3) * scale + 4)
		ui.Test.Buttons.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, scale).Moved(buttonPos))
	}

}

func ScreenToWorldSpace(r *Render, mat pixel.Matrix) pixel.Vec {
	return mat.Unproject(r.Window.MousePosition())
}