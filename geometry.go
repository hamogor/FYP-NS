package main

import (
	"github.com/faiface/pixel"
)

type corner int

const (
	TL corner = iota
	TR        = 1
	BL        = 2
	BR        = 3
)

func screenRect(mat pixel.Matrix, sprite *pixel.Sprite, r *Render) pixel.Rect {
	//rect := pixel.R(sprite.Frame().Max.Y, 0, sprite.Frame().Max.X, sprite.Frame().Max.X)
	//posWin := r.Window.MousePosition()
	//posGame := mat.Unproject(posWin)
	//fmt.Print(posWin, " ", posGame, " ", rect, "\n")
	//
	//nMat := anchorTL(mat, sprite, percentW(0), percentH(0)) // Origin point
	//
	//fmt.Print(nMat.String(), " ", mouse,  "\n")
	//bl := nMat[4]/nMat[5] * 100
	//fmt.Print(bl, " ", mouse,  "\n")
	return pixel.Rect{}
}

func percentW(percent float64) float64 {
	offset := (WWidth * percent) / 100
	return offset
}

func percentH(percent float64) float64 {
	offset := (WHeight * percent) / 100
	return offset
}

func contains(r pixel.Rect, u pixel.Vec) bool {
	return r.Min.X <= u.X && u.X <= r.Max.X && r.Min.Y <= u.Y && u.Y <= r.Max.Y
}

func anchorAndScale(sprite *pixel.Sprite, corner corner, offset pixel.Vec, scale float64) pixel.Matrix {
	v := pixel.Vec{}
	switch corner {
	case TL:
		v = pixel.V((sprite.Frame().Max.X/2+offset.X)*scale+4,
			WHeight-(sprite.Frame().Max.Y/2+offset.Y)*scale-4)
		break
	case TR:
		v = pixel.V(WWidth-(sprite.Frame().Max.X/2+offset.X)*scale-4,
			WHeight-(sprite.Frame().Max.Y/2-offset.Y)*scale-4)
		break
	case BL:
		v = pixel.V((sprite.Frame().Max.X/2+offset.X)*scale+4,
			(sprite.Frame().Max.Y/2+offset.Y)*scale-4)
	case BR:
		v = pixel.V(WWidth-(sprite.Frame().Max.X/2+offset.X)*scale-4,
			(sprite.Frame().Max.Y/2+offset.Y)*scale-4)
	}
	return pixel.IM.Scaled(pixel.ZV, scale).Moved(v)
}

func anchor(button *Button) pixel.Matrix {
	v := pixel.Vec{}
	r := button.Sprite.Frame()
	off := &button.Offset
	switch button.Corner {
	case TL:
		v = pixel.V((r.Max.X/2 + off.X) * button.Scale,
			WHeight - (r.Max.Y/2 + off.Y) * button.Scale - edgeOffset)
		break
	case TR:
		v = pixel.V(WWidth - (r.Max.X / 2 + off.X) * button.Scale - edgeOffset,
			WHeight - (r.Max.Y / 2 + off.Y) * button.Scale - edgeOffset)
		break
	case BL:
		v = pixel.V(WWidth - (r.Max.X / 2 + off.X) * button.Scale + edgeOffset,
			(r.Max.Y / 2 + off.Y) * button.Scale - edgeOffset)
		break
	case BR:
		v = pixel.V(WWidth - (r.Max.X / 2 + off.X) * button.Scale - edgeOffset,
			(r.Max.Y / 2 + off.Y) * button.Scale - edgeOffset)
		break
	}
	button.Pos = v
	button.setRect()
	return pixel.IM.Scaled(pixel.ZV, button.Scale).Moved(v)
}

func (button *Button) setRect() {

	button.Rect = &pixel.Rect{
		Min: pixel.Vec{X: button.Pos.X - (button.Sprite.Frame().Max.X/2 * button.Scale) , Y: button.Pos.Y - (button.Sprite.Frame().Max.Y/2 * button.Scale)},
		Max: pixel.Vec{X: button.Pos.X + (button.Sprite.Frame().Max.X/2 * button.Scale), Y: button.Pos.Y + (button.Sprite.Frame().Max.Y/2 * button.Scale)},
	}
	//fmt.Print(button.Rect, "\n")


}

func anchorTL(mat pixel.Matrix, sprite *pixel.Sprite, offX, offY, scale float64, btn *Button) pixel.Matrix {
	v := pixel.V(sprite.Frame().Max.X/2+offX, (WHeight-sprite.Frame().Max.Y/2)-offY) // Origin
	rect := pixel.Rect{
		Min: pixel.Vec{X: v.X - (sprite.Frame().Max.X*scale) / scale, Y: v.Y - (sprite.Frame().Max.Y*scale) / scale},
		Max: pixel.Vec{X: v.X + (sprite.Frame().Max.X*scale) / scale, Y: v.Y + (sprite.Frame().Max.Y*scale) / scale},
	}
	if btn != nil {
		btn.Rect = &rect
	}

	return pixel.IM.Scaled(pixel.ZV, scale).Moved(v)
}
