package controllers

import (
	"github.com/hculpan/go-life/app/model"
	"github.com/hculpan/go-life/app/pages"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/game"
)

type LifeGameController struct {
	game.GameController
}

func NewLifeGameController(window *component.Window, game *model.GameOfLife) LifeGameController {
	result := LifeGameController{}
	result.Game = game
	result.Window = window

	result.RegisterPages()

	return result
}

func (g *LifeGameController) RegisterPages() {
	component.RegisterPage(pages.NewGamePage("GamePage", 0, 0, g.Window.Width, g.Window.Height))
}
