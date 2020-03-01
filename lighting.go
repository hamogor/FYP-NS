package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type Light struct {
	color pixel.RGBA
	point pixel.Vec
	angle float64
	radius float64
	spread float64
	imd *imdraw.IMDraw
	Canvas *pixelgl.Canvas

}

func createLight(r *Render) *Light {
	light :=  &Light{
		color:  pixel.RGBA{},
		point:  r.Window.Bounds().Center(),
		angle:  math.Pi / 4,
		radius: 800,
		spread: math.Pi / math.E,
		imd:    imdraw.New(nil),
		Canvas: pixelgl.NewCanvas(r.Window.Bounds()),
	}
	light.imd.Color = pixel.Alpha(1)
	light.imd.Push(pixel.ZV)
	light.imd.Color = pixel.Alpha(0)
	for angle := -light.spread / 2; angle <= light.spread/2; angle += light.spread / 64 {
		light.imd.Push(pixel.V(1, 0).Rotated(angle))
	}
	light.imd.Polygon(0)
	light.point = r.Window.Bounds().Center()

	return light
}

func (light *Light) applyLight(r *Render) *Light {
	light.Canvas.SetMatrix(pixel.IM.Scaled(pixel.ZV, light.radius).Rotated(pixel.ZV, light.angle).Moved(light.point))
	light.Canvas.SetColorMask(pixel.Alpha(1))
	light.Canvas.SetComposeMethod(pixel.ComposePlus)
	light.imd.Draw(light.Canvas)



	return light

}
