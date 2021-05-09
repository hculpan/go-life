package components

import (
	"github.com/hculpan/go-life/app/model"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/sdl"
)

type LifeBoardComponent struct {
	component.BaseComponent
}

func NewLifeBoardComponent(x, y, width, height int32) *LifeBoardComponent {
	result := &LifeBoardComponent{}

	result.SetPosition(x, y)
	result.SetSize(width, height)

	return result
}

func (c *LifeBoardComponent) DrawComponent(r *sdl.Renderer) error {
	r.SetDrawColor(0, 0, 0, 0)
	rect := sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height}
	r.FillRect(&rect)

	rect.W = 3
	rect.H = 3

	r.SetDrawColor(0, 255, 0, 255)
	for x := 0; x < model.Game.BoardWidth; x++ {
		for y := 0; y < model.Game.BoardHeight; y++ {
			if model.Game.GetCurrentBoardState(x, y) > 0 {
				rect.X = int32(x*4+1) + c.X
				rect.Y = int32(y*4+1) + c.Y
				r.FillRect(&rect)
			}
		}
	}

	return nil
}

func (c *LifeBoardComponent) Draw(r *sdl.Renderer) error {
	if err := component.DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
