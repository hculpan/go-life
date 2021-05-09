package components

import (
	"fmt"

	"github.com/hculpan/go-life/app/model"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/resources"
	"github.com/veandco/go-sdl2/sdl"
)

type CyclesComponent struct {
	component.BaseComponent
}

func NewCyclesComponent(x, y, width, height int32) *CyclesComponent {
	result := &CyclesComponent{}

	result.SetPosition(x, y)
	result.SetSize(width, height)

	return result
}

func (c *CyclesComponent) DrawComponent(r *sdl.Renderer) error {
	msg := fmt.Sprintf("Cycle : %d", model.Game.Cycle)
	text, err := resources.Fonts.CreateTexture(msg, sdl.Color{R: 50, G: 255, B: 50, A: 255}, "HackBold-24", r)
	if err != nil {
		return err
	}
	_, _, w, h, err := text.Query()
	if err != nil {
		return err
	}
	r.Copy(text, &sdl.Rect{X: 0, Y: 0, W: w, H: h}, &sdl.Rect{X: c.X + 5, Y: c.Y, W: int32(w), H: int32(h)})

	return nil
}

func (c *CyclesComponent) Draw(r *sdl.Renderer) error {
	if err := component.DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
