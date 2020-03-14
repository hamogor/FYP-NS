package main

type Scenes struct {
	CurrentScene   SceneName
	ActiveElements map[UiElement]bool
}

type SceneName int

const (
	MainMenuScene SceneName = iota
	GameScene               = 2
)

type UiElement int

const (
	MiniMapActive      UiElement = iota
	MenuBarActive                = 2
	PortraitActive               = 3
	MainMenuActive               = 4
	PlayerHealthActive           = 5
)

func (g *Game) initScenes() {
	ae := make(map[UiElement]bool, 0)
	ae[MainMenuActive] = true
	g.Scenes = &Scenes{
		CurrentScene:   MainMenuScene,
		ActiveElements: ae,
	}
}

func (s *Scenes) resetActiveElements() {
	s.ActiveElements[MainMenuActive] = false
	s.ActiveElements[MenuBarActive] = false
	s.ActiveElements[PortraitActive] = false
	s.ActiveElements[MiniMapActive] = false
	s.ActiveElements[PlayerHealthActive] = false
}

func (s *Scenes) setActiveUiElements() {
	s.resetActiveElements()
	switch s.CurrentScene {
	case MainMenuScene:
		s.ActiveElements[MainMenuActive] = true
		break
	case GameScene:
		s.ActiveElements[MiniMapActive] = true
		s.ActiveElements[PortraitActive] = true
		s.ActiveElements[MenuBarActive] = true
		s.ActiveElements[PlayerHealthActive] = true
		break
	}
}
