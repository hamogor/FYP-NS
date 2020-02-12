package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
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
}

type Layer struct {
	Batch  *pixel.Batch
	Canvas *pixelgl.Canvas
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
		},
		Actors: Layer{
			Batch:  pixel.NewBatch(&pixel.TrianglesData{}, g.Assets.Sheets.Sprites),
			Canvas: pixelgl.NewCanvas(pixel.R(-WWidth/2, -WHeight/2, WWidth/2, WHeight/2)),
		},
	}
	g.Render = r
}

func (g *Game) render() {
	g.Render.Env.Batch.Clear()
	g.Render.Actors.Batch.Clear()
	g.Render.Env.Canvas.Clear(color.RGBA{
		R: 16,
		G: 10,
		B: 11,
		A: 255,
	})
	g.Render.Actors.Canvas.Clear(color.Transparent)
	g.Render.Ui = g.Render.Ui[:0]

	g.Render.Camera = pixel.Lerp(g.Render.Camera, g.Player.Actor.Pos.sToVec(), 1-math.Pow(1.0/128, dt))
	cam := pixel.IM.Moved(pixel.V(RoundFloat(g.Render.Camera.X), RoundFloat(g.Render.Camera.Y)).Scaled(-1))

	g.Render.Env.Batch.SetMatrix(cam)
	g.Render.Actors.Batch.SetMatrix(cam)

	g.Render.renderEnvironment(g.Level, g.Player)
	g.Render.renderActors(g.Player, g.Level, g.Assets)


	g.Render.Env.Batch.Draw(g.Render.Env.Canvas)
	g.Render.Actors.Batch.Draw(g.Render.Actors.Canvas)

	g.Render.Env.Canvas.SetMatrix(pixel.IM.Scaled(pixel.ZV, math.Min(Scaled, Scaled)).Moved(g.Render.Env.Canvas.Bounds().Center()))
	g.Render.Actors.Canvas.SetMatrix(pixel.IM.Scaled(pixel.ZV, math.Min(Scaled, Scaled)).Moved(g.Render.Actors.Canvas.Bounds().Center()))

	g.Render.Env.Canvas.Draw(g.Render.Window, pixel.IM.Moved(g.Render.Window.Bounds().Center()))
	g.Render.Actors.Canvas.Draw(g.Render.Window, pixel.IM.Moved(g.Render.Window.Bounds().Center()))

	g.Render.renderMiniMap(g.Player, g.Ui)
	g.Render.Window.Update()
}

func (r *Render) renderActors(p *Player, l *Level, a *Assets) {
	p.Actor.updateAnimState()
	p.Actor.CAnim.Sprite.Set(a.Sheets.Sprites, p.Actor.CAnim.Frame.Rect)
	p.Actor.CAnim.Sprite.Draw(r.Actors.Batch, pixel.IM.ScaledXY(pixel.ZV, pixel.V(-p.Actor.Direction, 1)).Moved(p.Actor.Pos.sToVec()))
}

func (r *Render) renderEnvironment(l *Level, p *Player) {
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

func (r *Render) renderMiniMap(p *Player, ui *Ui) {
	tr := pixel.V(WWidth-(LevelW*2), WHeight-(LevelH*2))
	ui.MiniMap.Sprite.Draw(r.Window, pixel.IM.Moved(tr).Scaled(pixel.V(tr.X+3, tr.Y+3), math.Min(4, 4)))
	ui.MiniMap.Msprite.Draw(r.Window, pixel.IM.Scaled(pixel.ZV, math.Min(3, 3)).ScaledXY(pixel.ZV, pixel.V(1, -1)).Moved(tr))
}
