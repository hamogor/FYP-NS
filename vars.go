package main

import (
	"image"
	"time"
)

var (
	nextPos               float64 = 0
	rect                          = image.Rectangle{}
	jsonPath                      = "./assets/json"
	pngPath                       = "./assets/png"
	tileSheetOutputPath           = "./assets/tiles.png"
	spriteSheetOutputPath         = "./assets/sprites.png"
	frames                        = 0
	dt                            = 0.0
	PPos						  = Position{X: 0, Y: 0}
	edgeOffset            float64 = 4
	second                        = time.Tick(time.Second)
	last                          = time.Now()
	WHeight               float64 = 768
	WWidth                float64 = 1366
	tilePaths                     = []string{
		"/l_wall",
		"/d_wall",
		"/l_floor",
		"/d_floor",
		"/door_n_s",
		"/door_w_e",
	}
	spritePaths = []string{ // Name and filename of sprite assets
		"/player_idle",
	}
	uiPaths = map[string]string{
		"background": "/assets/bg/island_2.png",
		"logo":       "/assets/bg/title.png",

		"startBtn":    "/assets/png/gui/start_btn.png",
		"startBtnH":   "/assets/png/gui/start_btn_hover.png",
		"optionsBtn":  "/assets/png/gui/options_btn.png",
		"optionsBtnH": "/assets/png/gui/options_btn_hover.png",
		"aboutBtn":    "/assets/png/gui/about_btn.png",
		"aboutBtnH":   "/assets/png/gui/about_btn_hover.png",
		"exitBtn":     "/assets/png/gui/exit_btn.png",
		"exitBtnH":    "/assets/png/gui/exit_btn_hover.png",

		"bar":        "/assets/png/gui/bar.png",
		"buttons":    "/assets/png/gui/buttons.png",
		"hbar":       "/assets/png/gui/hbar.png",
		"look":       "/assets/png/gui/look.png",
		"objectives": "/assets/png/gui/objectives.png",
	}
)

const (
	Scaled           = 2
	TileW, TileH     = 32, 48
	SpriteW, SpriteH = 32, 48
	LevelW, LevelH   = 64, 64
)
