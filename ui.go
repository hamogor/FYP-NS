package main

import (
	"github.com/faiface/pixel"
	"image"
	"image/color"
	"image/draw"
	"os"
)

type Ui struct {
	MiniMap  MiniMap
	MenuBar  MenuBar
	Portrait Portrait
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
	Background *pixel.Sprite
	StartButton Button
	OptionsButton Button
	AboutButton Button
	ExitButton Button
	Logo *pixel.Sprite
	Matrix pixel.Matrix
}

type Button struct {
	Pos pixel.Rect
	Sprite *pixel.Sprite
	HSprite *pixel.Sprite
	Handler buttonHandler
	Hovering bool
	Rect pixel.Rect
}

type MiniMap struct {
	Sprite  *pixel.Sprite // Objectives sprite
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
	optionsBtn, optionsBtnHover := GetPixelPicture("./assets/png/options_btn.png"), GetPixelPicture("./assets/png/options_btn_hover.png")
	aboutBtn, aboutBtnHover := GetPixelPicture("./assets/png/about_btn.png"), GetPixelPicture("./assets/png/about_btn_hover.png")
	exitBtn, exitBtnHover := GetPixelPicture("./assets/png/exit_btn.png"), GetPixelPicture("./assets/png/exit_btn_hover.png")
	logo := GetPixelPicture("./assets/bg/title.png")

	// NEW UI
	bar := GetPixelPicture("./assets/png/ui/sliced/bar.png")
	buttons := GetPixelPicture("./assets/png/ui/sliced/buttons.png")
	hbar := GetPixelPicture("./assets/png/ui/sliced/hbar.png")
	look := GetPixelPicture("./assets/png/ui/sliced/look.png")
	objectives := GetPixelPicture("./assets/png/ui/sliced/objectives.png")
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
			Logo: pixel.NewSprite(logo, logo.Bounds()),
			StartButton:    Button{
				Pos:     pixel.Rect{},
				Sprite:  pixel.NewSprite(startButtonSprite, startButtonSprite.Bounds()),
				HSprite: pixel.NewSprite(startButtonHSprite, startButtonHSprite.Bounds()),
				Handler: startButton,
			},
			OptionsButton: Button{
				Pos: pixel.Rect{},
				Sprite: pixel.NewSprite(optionsBtn, optionsBtn.Bounds()),
				HSprite: pixel.NewSprite(optionsBtnHover, optionsBtnHover.Bounds()),
				Handler: nil,
			},
			AboutButton: Button{
				Pos: pixel.Rect{},
				Sprite: pixel.NewSprite(aboutBtn, aboutBtn.Bounds()),
				HSprite: pixel.NewSprite(aboutBtnHover, aboutBtnHover.Bounds()),
				Handler: nil,
			},
			ExitButton: Button{
				Pos: pixel.Rect{},
				Sprite: pixel.NewSprite(exitBtn, exitBtn.Bounds()),
				HSprite: pixel.NewSprite(exitBtnHover, exitBtnHover.Bounds()),
				Handler: exitButton,
			},
		},
		Test: Test{
			Objectives: pixel.NewSprite(objectives, objectives.Bounds()),
			Look:       pixel.NewSprite(look, look.Bounds()),
			Bar:        pixel.NewSprite(bar, bar.Bounds()),
			HBar:       pixel.NewSprite(hbar, hbar.Bounds()),
			Buttons:    pixel.NewSprite(buttons, buttons.Bounds()),
		},
	}
	g.Ui = ui
}

func buildTest(x, y int) *pixel.Sprite {
	boxBL := GetUnparsedImg("/ui/box_BL")
	boxBM := GetUnparsedImg("/ui/box_BM")
	boxBR := GetUnparsedImg("/ui/box_BR")
	boxL := GetUnparsedImg("/ui/box_L")
	boxM := GetUnparsedImg("/ui/box_M")
	boxR := GetUnparsedImg("/ui/box_R")
	boxTL := GetUnparsedImg("/ui/box_TL")
	boxTM := GetUnparsedImg("/ui/box_TM")
	boxTR := GetUnparsedImg("/ui/box_TR")

	w, h := x, y
	r := image.Rectangle{Max: image.Point{X: w, Y: h}}
	img := image.NewRGBA(r)

	// BOTTOM LEFT
	blRect := image.Rectangle{
		Min: image.Point{X: boxBL.Bounds().Min.X, Y:  h - boxBL.Bounds().Max.Y},
		Max: image.Point{X: boxBL.Bounds().Max.X, Y: boxBL.Bounds().Max.Y*2 + h},
	}
	draw.Draw(img, blRect, boxBL, image.Point{X: 0, Y: 0}, draw.Src)
	// BOTTOM LEFT

	// BOTTOM RIGHT
	brRect := image.Rectangle{
		Min: image.Point{X: w - boxBR.Bounds().Max.X, Y: h - boxBR.Bounds().Max.Y},
		Max: image.Point{X: boxBR.Bounds().Max.X*2 + w, Y: boxBR.Bounds().Max.Y*2 + h},
	}
	draw.Draw(img, brRect, boxBR, image.Point{X: 0, Y: 0}, draw.Src)


	// TOP LEFT
	tlRect := image.Rectangle{
		Min: image.Point{X: boxTL.Bounds().Min.X, Y: boxTL.Bounds().Min.Y},
		Max: image.Point{X: boxTL.Bounds().Max.X, Y: boxTL.Bounds().Max.Y},
	}
	draw.Draw(img, tlRect, boxTL, image.Point{X: 0, Y: 0}, draw.Src)


	// TOP RIGHT
	trRect := image.Rectangle{
		Min: image.Point{X: w - boxTR.Bounds().Max.X, Y: boxTR.Bounds().Min.Y},
		Max: image.Point{X: boxTR.Bounds().Max.X*2 + w, Y: boxTR.Bounds().Max.Y},
	}
	draw.Draw(img, trRect, boxTR, image.Point{X: 0, Y: 0}, draw.Src)


	// LEFT
	vertH := h - (boxBL.Bounds().Max.Y - boxTL.Bounds().Max.Y)
	horiW := w - boxTR.Bounds().Max.X
	for i := 0; i < vertH; i++ {
		nextPos := image.Point{X: 0, Y: boxL.Bounds().Max.Y + i}
		rect = image.Rectangle{
			Min: nextPos,
			Max: nextPos.Add(boxL.Bounds().Max),
		}
		if i == h - boxBL.Bounds().Max.Y {break}
		draw.Draw(img, rect, boxL, image.Point{X: 0, Y: 0}, draw.Src)
	}

	// RIGHT
	for i := 0; i < vertH; i++ {
		nextPos := image.Point{X: w - boxTR.Bounds().Max.X, Y: boxL.Bounds().Max.Y + i}
		rect = image.Rectangle{
			Min: nextPos,
			Max: nextPos.Add(boxR.Bounds().Max),
		}
		if i == h - boxBR.Bounds().Max.Y {break}
		draw.Draw(img, rect, boxR, image.Point{X: 0, Y: 0}, draw.Src)
	}

	// TOP
	for i := 0; i < horiW; i++ {
		nextPos := image.Point{X: boxTL.Bounds().Max.X + i, Y: 0}
		rect = image.Rectangle{
			Min: nextPos,
			Max: nextPos.Add(boxTM.Bounds().Max),
		}
		if i == horiW- 1 {break}
		draw.Draw(img, rect, boxTM, image.Point{X: 0, Y: 0}, draw.Src)
	}

	// BOTTOM
	for i := 0; i < horiW; i++ {
		nextPos := image.Point{X: boxBL.Bounds().Max.X + i, Y: h - boxBR.Bounds().Max.Y}
		rect = image.Rectangle{
			Min: nextPos,
			Max: nextPos.Add(boxBM.Bounds().Max),
		}
		if i == horiW- 1 {break}
		draw.Draw(img, rect, boxBM, image.Point{X: 0, Y: 0}, draw.Src)
	}

	middleRect := image.Rectangle{
		Min: image.Point{
			X: boxBL.Bounds().Max.X,
			Y: boxBL.Bounds().Max.Y,
		},
		Max: image.Point{
			X: horiW,
			Y: vertH - boxBL.Bounds().Max.X,
		},
	}
	for i := middleRect.Min.X; i < middleRect.Max.X; i++ {
		for j := middleRect.Min.Y; j < middleRect.Max.Y; j++ {
			nextPos := image.Point{X: i, Y: j}
			rect = image.Rectangle{
				Min: nextPos,
				Max: nextPos.Add(boxM.Bounds().Max),
			}
			draw.Draw(img, rect, boxM, image.Point{X: 0, Y: 0}, draw.Src)
		}
	}




	pic := pixel.PictureDataFromImage(img)
	return  pixel.NewSprite(pic, rectToRectangle(img.Bounds()))

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

func exitButton(s *Scenes) {
	if s.CurrentScene == MainMenuScene {
		os.Exit(0)
	}
}


//middleW := w - (boxBL.Bounds().Max.X + boxBR.Bounds().Max.X)
//leftH := h - (boxBL.Bounds().Max.Y + boxTL.Bounds().Max.Y)
//r := image.Rectangle{Max: image.Point{X: w, Y: h}}
//img := image.NewRGBA(r)
//
//
//
////BOTTOM
//draw.Draw(img, boxBL.Bounds(), boxBL, image.Point{X: 0, Y: 0}, draw.Src) // Draw BL
//for i := 0; i < middleW; i++ {
//	nextPos := image.Point{X: boxBL.Bounds().Max.X + i, Y: 0}
//	rect = image.Rectangle{
//		Min: nextPos,
//		Max: nextPos.Add(boxBL.Bounds().Max),
//	}
//	draw.Draw(img, rect, boxBM, image.Point{X: 0, Y: 0}, draw.Src)
//}
//draw.Draw(img, rect, boxBR, image.Point{X: 0, Y: 0}, draw.Src)
//
//for i := 0; i < leftH; i++ {
//	nextPos := image.Point{X: 0, Y: boxBL.Bounds().Max.Y + i}
//	rect = image.Rectangle{
//		Min: nextPos,
//		Max: nextPos.Add(boxBL.Bounds().Max),
//	}
//	draw.Draw(img, rect, boxL, image.Point{X: 0, Y: 0}, draw.Src)
//}


