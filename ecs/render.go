package ecs

import (
	"github.com/EngoEngine/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"log"
	"math"
)

type RenderSystem struct {
	ecs.System
	Entities []*Renderable
	Window *pixelgl.Window
	Config *pixelgl.WindowConfig
	Env Layer
	Actors Layer
}

type Layer struct {
	Batch *pixel.Batch
	Canvas *pixelgl.Canvas
	Matrix pixel.Matrix
}

type Renderable struct {
	ecs.BasicEntity
	*RenderComponent
}

type RenderComponent struct {
	Sprite *pixel.Sprite
}

func InitRenderSystem(g *Game) *RenderSystem {
	cfg := &pixelgl.WindowConfig{
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
	win, err := pixelgl.NewWindow(*cfg)
	if err != nil {
		log.Println(err)
	}

	g.Window = win
	g.Config = *cfg

	return &RenderSystem{
		Entities:   make([]*Renderable, 0),
		Window:     win,
		Config:     cfg,
		Env: Layer{
			Batch:  pixel.NewBatch(&pixel.TrianglesData{}, *g.AssetStore.Sheets.Env),
			Canvas: pixelgl.NewCanvas(pixel.R(-WWidth/2, -WHeight/2, WWidth/2, WHeight/2)),
			Matrix: pixel.IM.Scaled(pixel.ZV, math.Min(Scaled, Scaled)).Moved(pixel.V(WWidth/2, WHeight/2)),
		},
		Actors: Layer{
			Batch:  pixel.NewBatch(&pixel.TrianglesData{}, *g.AssetStore.Sheets.Sprites),
			Canvas: pixelgl.NewCanvas(pixel.R(-WWidth/2, -WHeight/2, WWidth/2, WHeight/2)),
			Matrix: pixel.IM.Scaled(pixel.ZV, math.Min(Scaled, Scaled)).Moved(pixel.V(WWidth/2, WHeight/2)),
		},
	}
}

func (r *RenderSystem) Add(basic *ecs.BasicEntity, render *RenderComponent) {
	r.Entities = append(r.Entities, &Renderable{*basic, render})
}



func (r *RenderSystem) Update(dt float32) {
	r.Actors.Batch.Clear()
	r.Actors.Canvas.Clear(color.RGBA{R: 8, G: 8, B: 12, A: 255})
	for i := range r.Entities {
		r.Entities[i].Sprite.Draw(r.Actors.Batch, pixel.IM.Moved(r.Window.Bounds().Center()))
	}
	r.Actors.Batch.Draw(r.Actors.Canvas)
	r.Actors.Canvas.Draw(r.Window, pixel.IM.Moved(r.Window.Bounds().Center()))
	r.Window.Update()
}

func (r *RenderSystem) Remove(e ecs.BasicEntity) {

}
