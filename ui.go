package main

import (
	"github.com/faiface/pixel"
	"os"
)

type Ui struct {
	MainMenu MainMenu
	Test     Test
}

type buttonHandler func(s *Scenes)

type Test struct {
	Objectives *pixel.Sprite
	Look       *pixel.Sprite
	Bar        *pixel.Sprite
	HBar       *pixel.Sprite
	Buttons    *pixel.Sprite
}

type MainMenu struct {
	Background    *pixel.Sprite
	StartButton   *Button
	OptionsButton *Button
	AboutButton   *Button
	ExitButton    *Button
	Logo          *pixel.Sprite
	Matrix        pixel.Matrix
}

type Button struct {
	Sprite   *pixel.Sprite
	HSprite  *pixel.Sprite
	Handler  buttonHandler
	Hovering bool
	Rect     *pixel.Rect
	Offset   pixel.Vec
	Pos      pixel.Vec
	Scale    float64
	Corner   corner
}

func (g *Game) initUi() {
	mMYOffset := BuildSprite(uiPaths["startBtn"]).Frame().Max.Y
	ui := &Ui{
		MainMenu: MainMenu{
			Background: BuildSprite(uiPaths["background"]),
			Logo:       BuildSprite(uiPaths["logo"]),
			StartButton: &Button{
				Sprite:  BuildSprite(uiPaths["startBtn"]),
				HSprite: BuildSprite(uiPaths["startBtnH"]),
				Handler: startButton,
				Offset:  pixel.V(percentW(8), percentH(20)),
				Scale:   2,
				Corner:  TL,
				Hovering: false,
				Rect:    &pixel.Rect{},
			},
			OptionsButton: &Button{
				Sprite:  BuildSprite(uiPaths["optionsBtn"]),
				HSprite: BuildSprite(uiPaths["optionsBtnH"]),
				Handler: nil,
				Offset:  pixel.V(percentW(8), percentH(20) + mMYOffset),
				Scale:   2,
				Corner:  TL,
				Hovering: false,
				Rect:    &pixel.Rect{},
			},
			AboutButton: &Button{
				Sprite:  BuildSprite(uiPaths["aboutBtn"]),
				HSprite: BuildSprite(uiPaths["aboutBtnH"]),
				Handler: nil,
				Offset:  pixel.V(percentW(8), percentH(20) + mMYOffset * 2),
				Scale:   2,
				Corner:  TL,
				Hovering: false,
				Rect:    &pixel.Rect{},
			},
			ExitButton: &Button{
				Sprite:  BuildSprite(uiPaths["exitBtn"]),
				HSprite: BuildSprite(uiPaths["exitBtnH"]),
				Handler: exitButton,
				Offset:  pixel.V(percentW(8), percentH(20) + mMYOffset * 3),
				Scale:   2,
				Corner:  TL,
				Hovering: false,
				Rect:    &pixel.Rect{},
			},
		},
		Test: Test{
			Objectives: BuildSprite(uiPaths["objectives"]),
			Look:       BuildSprite(uiPaths["look"]),
			Bar:        BuildSprite(uiPaths["bar"]),
			HBar:       BuildSprite(uiPaths["hbar"]),
			Buttons:    BuildSprite(uiPaths["buttons"]),
		},
	}


	g.Ui = ui
}

func startButton(s *Scenes) {
	if s.CurrentScene == MainMenuScene {
		s.CurrentScene = GameScene
	} else if s.CurrentScene == GameScene {
		s.CurrentScene = MainMenuScene
	}
}

func exitButton(s *Scenes) {
	if s.CurrentScene == MainMenuScene {
		os.Exit(0)
	}
}

func (p *Player) updateMiniMap(l *Level, ui *Ui) {
	//for x := 0; x < LevelW; x++ {
	//	for y := 0; y < LevelH; y++ {
	//		if p.Actor.Fov.Look(x, y) {
	//			if l.Tiles[x][y].Terrain == Wall {
	//				ui.MiniMap.Map.Set(x, y, color.RGBA{R: 37, G: 43, B: 69, A: 255})
	//			}
	//			if l.Tiles[x][y].Terrain == Floor {
	//				ui.MiniMap.Map.Set(x, y, color.RGBA{R: 57, G: 68, B: 100, A: 255})
	//			}
	//			if l.Tiles[x][y].Terrain == Door || l.Tiles[x][y].Terrain == OpenDoor {
	//				ui.MiniMap.Map.Set(x, y, color.RGBA{R:209, G:119, B:8, A:255})
	//			}
	//
	//		} else if p.Actor.Fov.explored[x][y] {
	//			if l.Tiles[x][y].Terrain == Wall {
	//				ui.MiniMap.Map.Set(x, y, color.RGBA{R: 18, G: 21, B: 34, A: 255})
	//			}
	//			if l.Tiles[x][y].Terrain == Floor {
	//				ui.MiniMap.Map.Set(x, y, color.RGBA{R: 37, G: 43, B: 67, A: 255})
	//			}
	//		}
	//	}
	//}
	//ui.MiniMap.Map.Set(p.Actor.Pos.X, p.Actor.Pos.Y, color.RGBA{R: 64, G: 136, B: 72, A: 255})
	//minimapImage := pixel.PictureDataFromImage(ui.MiniMap.Map)
	//ui.MiniMap.Msprite = pixel.NewSprite(minimapImage, minimapImage.Rect)
}