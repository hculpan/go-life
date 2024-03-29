package pages

import (
	"github.com/hculpan/go-life/app/components"
	"github.com/hculpan/go-life/app/model"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/sdl"
)

type GamePage struct {
	component.BasePage
}

func NewGamePage(name string, x, y, width, height int32) *GamePage {
	p := GamePage{}
	p.Initialize()
	p.Name = "GamePage"
	p.SetPosition(0, 0)
	p.SetSize(width, height)

	p.AddChild(components.NewLifeBoardComponent(0, 80, width, height-80))
	p.AddChild(components.NewHeaderComponent(0, 0, width, 80))

	return &p
}

func (g *GamePage) KeyEvent(event *sdl.KeyboardEvent) bool {
	keycode := sdl.GetKeyFromScancode(event.Keysym.Scancode)
	if keycode == sdl.K_r {
		model.Game.Reset()
		return true
	}

	return component.PassKeyEventToChildren(event, g.Children)
}

func (c *GamePage) Draw(r *sdl.Renderer) error {
	return component.DrawParentAndChildren(r, c)
}

func (c *GamePage) MouseButtonEvent(event *sdl.MouseButtonEvent) bool {
	if c.IsPointInComponent(event.X, event.Y) {
		return component.PassMouseButtonEventToChildren(event, c.Children)
	}

	return false
}
