package components

import (
	"github.com/hculpan/go-life/app/model"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/sdl"
)

type HeaderComponent struct {
	component.BaseComponent
}

func NewHeaderComponent(x, y, width, height int32) *HeaderComponent {
	result := &HeaderComponent{}
	result.Initialize()

	result.SetPosition(x, y)
	result.SetSize(width, height)

	result.AddChild(NewCyclesComponent(x+5, y+10, 500, 60))
	result.AddChild(component.NewButtonComponent(width-220, y+10, 200, 60,
		"Restart",
		sdl.Color{R: 25, G: 25, B: 25, A: 255},
		sdl.Color{R: 50, G: 255, B: 50, A: 255},
		sdl.Color{R: 0, G: 0, B: 0, A: 255},
		func() {
			model.Game.Reset()
		}))

	return result
}

func (c *HeaderComponent) DrawComponent(r *sdl.Renderer) error {
	r.SetDrawColor(25, 25, 25, 255)
	rect := sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height}
	r.FillRect(&rect)

	r.SetDrawColor(60, 60, 60, 0)
	r.DrawLine(c.X, c.Height, c.Width, c.Height)

	return nil
}

func (c *HeaderComponent) Draw(r *sdl.Renderer) error {
	if err := component.DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
