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
	MainMenu MainMenu
}

type buttonHandler func(s *Scenes)

type MainMenu struct {
	Background *pixel.Sprite
	StartButton Button
}

type Button struct {
	Pos pixel.Rect
	Sprite *pixel.Sprite
	HSprite *pixel.Sprite
	Handler buttonHandler
	Hovering bool
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
	mainMenuSprite := GetPixelPicture("./assets/bg/island_2.png")
	startButtonSprite := GetPixelPicture("./assets/png/start_btn.png")
	startButtonHSprite := GetPixelPicture("./assets/png/start_btn_hover.png")
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
		MainMenu: MainMenu{
			Background: pixel.NewSprite(mainMenuSprite, mainMenuSprite.Bounds()),
			StartButton:    Button{
				Pos:     pixel.Rect{},
				Sprite:  pixel.NewSprite(startButtonSprite, startButtonSprite.Bounds()),
				HSprite: pixel.NewSprite(startButtonHSprite, startButtonHSprite.Bounds()),
				Handler: startButton,
			},
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

func startButton(s *Scenes) {
	if s.CurrentScene == MainMenuScene {
		s.CurrentScene = GameScene
	} else if s.CurrentScene == GameScene {
		s.CurrentScene = MainMenuScene
	}
}


