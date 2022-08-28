package controllers

import (
	"github.com/hculpan/go-life/app/model"
	"github.com/hculpan/go-life/app/pages"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/game"
	"github.com/veandco/go-sdl2/sdl"
)

type LifeGameController struct {
	game.GameController
}

func NewLifeGameController(gameWidth, gameHeight int32) LifeGameController {
	result := LifeGameController{}

	windowBackground := sdl.Color{R: 0, G: 0, B: 0, A: 0}

	result.Window = component.NewWindow(gameWidth, gameHeight, "Conway's Game of Life", windowBackground)
	result.Game = model.NewGameOfLife(result.Window.Width, result.Window.Height, 0.1)

	result.RegisterPages()

	return result
}

func (g *LifeGameController) RegisterPages() {
	component.RegisterPage(pages.NewGamePage("GamePage", 0, 0, g.Window.Width, g.Window.Height))
}
