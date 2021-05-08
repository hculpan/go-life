package pages

import (
	"github.com/hculpan/go-life/app/components"
	"github.com/hculpan/go-sdl-lib/component"
)

type GamePage struct {
	component.BasePage
}

func NewGamePage(name string, x, y, width, height int32) *GamePage {
	p := GamePage{}
	p.Name = "GamePage"
	p.SetPosition(0, 0)
	p.SetSize(width, height)

	p.AddChild(components.NewLifeBoardComponent(0, 0, width, height))

	return &p
}
