package ecs

import "github.com/faiface/pixel"

type Animation struct {
	Frames []*Frame
	Loop bool
	Active bool
}

type AnimationComponent struct {
	Counter float64
	Index int
	Frame *Frame
	Sprite *pixel.Sprite
}

type Frame struct {
	Rect pixel.Rect
	Duration float64
}
