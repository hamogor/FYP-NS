package ecs

import (
	"github.com/EngoEngine/ecs"
	"github.com/faiface/pixel"
)

type PlayerEntity struct {
	ecs.BasicEntity
	*RenderComponent
}

func InitPlayer(world *ecs.World, a *AssetStore) {
	p := &PlayerEntity{
		BasicEntity:        ecs.NewBasic(),
		RenderComponent: &RenderComponent{
			Sprite: pixel.NewSprite(*a.Sheets.Sprites, a.Anims["player_idle"].Frames[0].Rect),
				},
	}
	for _, system := range world.Systems() {

		// Use a type-switch to figure out which System is which
		switch sys := system.(type) {
		// Create a case for each System you want to use
		case *RenderSystem:
			sys.Add(&p.BasicEntity, p.RenderComponent)
		}
	}
}
