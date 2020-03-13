package NS

import (
	"github.com/faiface/pixel"
)

type Animation struct {
	Frames []*Frame
	Loop   bool
	Active bool
}

type CAnim struct {
	Counter    float64
	State      animState
	FrameIndex int
	Frame      *Frame
	Sprite     *pixel.Sprite
}

type Frame struct {
	Rect     pixel.Rect
	Duration float64
}

type animState int

const (
	idle   animState = iota
	attack           = 1
)

func buildCAnim() *CAnim {
	return &CAnim{
		Counter:    0,
		State:      idle,
		FrameIndex: 0,
		Frame:      nil,
		Sprite:     pixel.NewSprite(nil, pixel.Rect{}),
	}
}

func (a *Actor) updateAnimState() {
	switch a.CAnim.State {
	case idle:
		a.play(a.Anims["player_idle"])
	}
}

func (a *Actor) play(anim *Animation) {
	a.CAnim.Counter += dt                           // Add to counter
	a.CAnim.Frame = anim.Frames[a.CAnim.FrameIndex] // Set to first frame
	if a.CAnim.Counter > anim.Frames[a.CAnim.FrameIndex].Duration {
		a.CAnim.FrameIndex++ // Go to next Frame
		a.CAnim.Counter = 0  // Reset frame timing
	}
	if a.CAnim.FrameIndex >= len(anim.Frames) {
		if anim.Loop {
			a.CAnim.FrameIndex = 0
		} else {
			a.CAnim.State = idle
		}
	}
}

