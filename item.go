package main

import (
	"github.com/faiface/pixel"
)

type Item struct {
	Name    string
	Sprites Sprites
	Active  bool
	Type    ItemType
	Handler ItemHandler
}

type ItemType int

const (
	Health ItemType = iota
	Ammo            = 1
	Nil             = 2
)

type ItemHandler func()

func ammo(g *Game) *Item {
	return &Item{
		Name: "Ammo",
		Sprites: Sprites{
			L: pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["ammo"][0]),
			D: pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["ammo"][1]),
		},
		Active:  true,
		Type:    1,
		Handler: ammoHandle,
	}
}

func health(g *Game) *Item {
	return &Item{
		Name: "Health",
		Sprites: Sprites{
			L: pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["health"][0]),
			D: pixel.NewSprite(g.Assets.Sheets.Environment, g.Assets.Env["health"][1]),
		},
		Active:  true,
		Type:    0,
		Handler: healthHandle,
	}
}

func ammoHandle() {

}

func healthHandle() {

}
