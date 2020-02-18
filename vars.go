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
	allowedInput                  = true
	second                        = time.Tick(time.Second)
	last                          = time.Now()
	WHeight               float64 = 1080
	WWidth                float64 = 1920
	tilePaths                     = []string{
		"/l_wall",
		"/d_wall",
		"/l_floor",
		"/d_floor",
		"/ammo",
		"/health",
		"/door_n_s",
		"/door_w_e",
	}
	spritePaths = []string{ // Name and filename of sprite assets
		"/player_idle",
	}
)

const (
	Scaled           = 1
	TileW, TileH     = 32, 48
	SpriteW, SpriteH = 32, 48
	LevelW, LevelH   = 40, 40
)
