package main

import (
	"github.com/faiface/pixel"
)

func anchorTL(mat pixel.Matrix, sprite *pixel.Sprite, offX, offY, scale float64, btn *Button) pixel.Matrix {
	v := pixel.V(sprite.Frame().Max.X/2 + offX,  (WHeight-sprite.Frame().Max.Y/2) -offY) // Origin
	rect := pixel.Rect{
		Min: pixel.Vec{X: v.X - (sprite.Frame().Max.X*scale) / scale, Y: v.Y - (sprite.Frame().Max.Y*scale) / scale},
		Max: pixel.Vec{X: v.X + (sprite.Frame().Max.X*scale) / scale, Y: v.Y + (sprite.Frame().Max.Y*scale) / scale},
	}
	if btn != nil {
		btn.Rect = rect
	}


	return pixel.IM.Scaled(pixel.ZV, scale).Moved(v)
}

func screenRect(mat pixel.Matrix, sprite *pixel.Sprite, r *Render) pixel.Rect {
	//rect := pixel.R(sprite.Frame().Max.Y, 0, sprite.Frame().Max.X, sprite.Frame().Max.X)
	//posWin := r.Window.MousePosition()
	//posGame := mat.Unproject(posWin)
	//fmt.Print(posWin, " ", posGame, " ", rect, "\n")

	//nMat := anchorTL(mat, sprite, percentW(0), percentH(0)) // Origin point

	//fmt.Print(nMat.String(), " ", mouse,  "\n")
	//bl := nMat[4]/nMat[5] * 100
	//fmt.Print(bl, " ", mouse,  "\n")
	return pixel.Rect{}
}

func percentW(percent float64) float64 {
	offset :=  (WWidth * percent) / 100
	return offset
}

func percentH(percent float64) float64 {
	offset :=  (WHeight * percent) / 100
	return offset
}

func contains(r pixel.Rect, u pixel.Vec) bool {
	return r.Min.X <= u.X && u.X <= r.Max.X && r.Min.Y <= u.Y && u.Y <= r.Max.Y
}

