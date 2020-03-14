package main

import (
	"image"
	"math"
)

type FovMap struct {
	w, h                    int
	blocked, seen, explored [][]bool
}

type FOVAlgo func(*FovMap, int, int, int, bool)

func NewMap(width, height int) *FovMap {
	blocked := make([][]bool, height)
	seen := make([][]bool, height)
	explored := make([][]bool, height)

	for y := 0; y < height; y++ {
		blocked[y] = make([]bool, width)
		seen[y] = make([]bool, width)
		explored[y] = make([]bool, width)
	}

	return &FovMap{width, height, blocked, seen, explored}
}

func initFov(l *Level) *FovMap {
	fov := NewMap(LevelW, LevelH)
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			pos := Position{X: x, Y: y}
			if pos.terrain(l) == Wall {
				fov.Block(x, y, true)
			} else if pos.terrain(l) == Floor {
				fov.Block(x, y, false)
			} else if pos.terrain(l) == Door {
				fov.Block(x, y, true)
			}
		}
	}
	return fov
}

func initialiseFovToLevel(l *Level) *FovMap {
	fov := NewMap(LevelW, LevelH)
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			pos := Position{X: x, Y: y}

			if pos.terrain(l) == Wall {
				fov.Block(x, y, true)
			} else if pos.terrain(l) == Floor {
				fov.Block(x, y, false)
			}
		}
	}
	return fov
}

func InitFovs() []*FovMap {
	var fovs []*FovMap
	for i := 0; i < 10; i++ {
		f := NewMap(LevelW, LevelH)
		fovs = append(fovs, f)
	}
	return fovs
}

func (f *FovMap) Fov(pos Position, radius int, includeWalls bool, algo FOVAlgo) {
	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			f.seen[y][x] = false
		}
	}
	algo(f, pos.X, pos.Y, radius, includeWalls)
}

func fovCircularCastRay(fov *FovMap, xo int, yo int, xd int, yd int, r2 int, walls bool) {
	curx := xo
	cury := yo
	in := false
	blocked := false
	if fov.In(curx, cury) {
		in = true
		fov.seen[cury][curx] = true
	}
	for _, p := range Line(xo, yo, xd, yd) {
		curx = p.X
		cury = p.Y
		if r2 > 0 {
			curRadius := (curx-xo)*(curx-xo) + (cury-yo)*(cury-yo)
			if curRadius > r2 {
				break
			}
		}
		if fov.In(curx, cury) {
			in = true
			if !blocked && fov.blocked[cury][curx] {
				blocked = true
			} else if blocked {
				break
			}
			if walls || !blocked {
				fov.seen[cury][curx] = true
			}
		} else if in {
			break
		}
	}

}

func fovCircularPostProc(fov *FovMap, x0, y0, x1, y1, dx, dy int) {
	for cx := x0; cx <= x1; cx++ {
		for cy := y0; cy <= y1; cy++ {
			x2 := cx + dx
			y2 := cy + dy
			if fov.In(cx, cy) && fov.Look(cx, cy) && !fov.blocked[cy][cx] {
				if x2 >= x0 && x2 <= x1 {
					if fov.In(x2, cy) && fov.blocked[cy][x2] {
						fov.seen[cy][x2] = true
					}
				}
				if y2 >= y0 && y2 <= y1 {
					if fov.In(cx, y2) && fov.blocked[y2][cx] {
						fov.seen[y2][cx] = true
					}
				}
				if x2 >= x0 && x2 <= x1 && y2 >= y0 && y2 <= y1 {
					if fov.In(x2, y2) && fov.blocked[y2][x2] {
						fov.seen[y2][x2] = true
					}
				}
			}
		}
	}
}

// FOVCicular raycasts out from the vantage in a circle.
func FOVCircular(fov *FovMap, x, y, r int, walls bool) {
	xo := 0
	yo := 0
	xmin := 0
	ymin := 0
	xmax := fov.Width()
	ymax := fov.Height()
	r2 := r * r
	if r > 0 {
		xmin = max(0, x-r)
		ymin = max(0, y-r)
		xmax = min(fov.Width(), x+r+1)
		ymax = min(fov.Height(), y+r+1)
	}
	xo = xmin
	yo = ymin
	for xo < xmax {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		xo++
	}
	xo = xmax - 1
	yo = ymin + 1
	for yo < ymax {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		yo++
	}
	xo = xmax - 2
	yo = ymax - 1
	for xo >= 0 {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		xo--
	}
	xo = xmin
	yo = ymax - 2
	for yo > 0 {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		yo--
	}
	if walls {
		fovCircularPostProc(fov, xmin, ymin, x, y, -1, -1)
		fovCircularPostProc(fov, x, ymin, xmax-1, y, 1, -1)
		fovCircularPostProc(fov, xmin, y, x, ymax-1, -1, 1)
		fovCircularPostProc(fov, x, y, xmax-1, ymax-1, 1, 1)
	}
}

func (f *FovMap) Block(x, y int, blocked bool) {
	f.blocked[y][x] = blocked
}

func (f *FovMap) Look(x, y int) bool {
	return f.seen[y][x]
}

func (f *FovMap) Width() int {
	return f.w
}

func (f *FovMap) Height() int {
	return f.h
}

func Line(x0, y0, x1, y1 int) []image.Point {
	dx := int(math.Abs(float64(x1 - x0)))
	dy := int(math.Abs(float64(y1 - y0)))
	sx := 0
	sy := 0
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}

	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy

	ps := make([]image.Point, 0)
	for {
		ps = append(ps, image.Pt(x0, y0))
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
	return ps
}

func (f *FovMap) In(x, y int) bool {
	return x >= 0 && x < f.w && y >= 0 && y < f.h
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (l *Level) calculateAllFovs() {
	for i := range l.Actors {
		l.Actors[i].calculateFov()
	}
}
