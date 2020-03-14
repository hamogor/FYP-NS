package main

func (g *Game) initData() {
	g.Player.Actor.Fov = initFov(g.Level)
	g.Player.Actor.Pos = g.Level.Spawn
	g.Render.Camera = g.Player.Actor.Pos.sToVec()
	g.Player.Actor.calculateFov()
	g.Player.updateMiniMap(g.Level, g.Ui)
}

func (a *Actor) calculateFov() {
	a.Fov.Fov(a.Pos, 10, true, FOVCircular)
}
