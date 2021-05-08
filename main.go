package main

import (
	"fmt"

	"github.com/hculpan/go-life/app/controllers"
	"github.com/hculpan/go-life/app/model"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	component.SetupSDL()

	displayMode, err := sdl.GetCurrentDisplayMode(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	gameWidth := int32(float64(displayMode.W) * 0.75)
	gameWidth += gameWidth % 4
	gameHeight := int32(float64(displayMode.H) * 0.75)
	gameHeight += gameWidth % 4

	// TODO: Set to the desired default window background
	windowBackground := sdl.Color{R: 255, G: 255, B: 255, A: 255}

	game := model.NewGameOfLife(gameWidth, gameHeight, 0.1)
	window := component.NewWindow(gameWidth, gameHeight, "GoSDL", windowBackground)
	gamecontroller := controllers.NewLifeGameController(window, game)
	if err := gamecontroller.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
