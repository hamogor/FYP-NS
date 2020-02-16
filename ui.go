package main

import (
	"github.com/faiface/pixel"
	"image"
	"image/color"
)

type Ui struct {
	MiniMap  MiniMap
	MenuBar  MenuBar
	Portrait Portrait
}

type MiniMap struct {
	Sprite  *pixel.Sprite // Box sprite
	Msprite *pixel.Sprite // Map sprite
	Map     *image.RGBA   // Map pixels
}

type MenuBar struct {
	LSprite *pixel.Sprite // Left side
	RSprite *pixel.Sprite // Right side
	Sprite  *pixel.Sprite // Inbetween
}

type Portrait struct {
	Sprite *pixel.Sprite
}

func (g *Game) initUi() {
	mapSprite := GetPixelPicture("./assets/png/minimap.png")
	barSprite := GetPixelPicture("./assets/png/bar.png")
	lBarSprite := GetPixelPicture("./assets/png/l_bar.png")
	rBarSprite := GetPixelPicture("./assets/png/r_bar.png")
	portraitSprite := GetPixelPicture("./assets/png/portrait.png")
	ui := &Ui{
		MiniMap: MiniMap{
			Sprite: pixel.NewSprite(mapSprite, mapSprite.Bounds()),
			Map:    image.NewRGBA(image.Rect(0, 0, LevelW, LevelH)),
		},
		MenuBar: MenuBar{
			LSprite: pixel.NewSprite(lBarSprite, lBarSprite.Bounds()),
			Sprite:  pixel.NewSprite(barSprite, barSprite.Bounds()),
			RSprite: pixel.NewSprite(rBarSprite, rBarSprite.Bounds()),
		},
		Portrait: Portrait{
			Sprite: pixel.NewSprite(portraitSprite, portraitSprite.Bounds()),
		},
	}
	g.Ui = ui
}

func (p *Player) updateMiniMap(l *Level, ui *Ui) {
	for x := 0; x < LevelW; x++ {
		for y := 0; y < LevelH; y++ {
			if p.Actor.Fov.Look(x, y) {
				if l.Tiles[x][y].Terrain == Wall {
					ui.MiniMap.Map.Set(x, y, color.RGBA{R: 37, G: 43, B: 69, A: 255})
				}
				if l.Tiles[x][y].Terrain == Floor {
					ui.MiniMap.Map.Set(x, y, color.RGBA{R: 57, G: 68, B: 100, A: 255})
				}
				if l.Tiles[x][y].Terrain == Door || l.Tiles[x][y].Terrain == OpenDoor {
					ui.MiniMap.Map.Set(x, y, color.RGBA{R:209, G:119, B:8, A:255})
				}

			} else if p.Actor.Fov.explored[x][y] {
				if l.Tiles[x][y].Terrain == Wall {
					ui.MiniMap.Map.Set(x, y, color.RGBA{R: 18, G: 21, B: 34, A: 255})
				}
				if l.Tiles[x][y].Terrain == Floor {
					ui.MiniMap.Map.Set(x, y, color.RGBA{R: 37, G: 43, B: 67, A: 255})
				}
			}
		}
	}
	ui.MiniMap.Map.Set(p.Actor.Pos.X, p.Actor.Pos.Y, color.RGBA{R: 64, G: 136, B: 72, A: 255})
	minimapImage := pixel.PictureDataFromImage(ui.MiniMap.Map)
	ui.MiniMap.Msprite = pixel.NewSprite(minimapImage, minimapImage.Rect)
}
