package main

import (
	"embed"
	"fmt"

	"github.com/hculpan/go-life/app/controllers"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/resources"
	"github.com/veandco/go-sdl2/sdl"
)

//go:embed resources/fonts/*

var appFonts embed.FS

func main() {
	component.SetupSDL()

	displayMode, err := sdl.GetCurrentDisplayMode(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := resources.FontsInit(appFonts); err != nil {
		fmt.Println(err)
		return
	}

	//	resources.Fonts.RegisterFont("HackBold-24", "resources/fonts/HackBold-Pdjd.ttf", 24)
	resources.Fonts.RegisterFont("HackBold-48", "built-in-fonts/TruenoLight.otf", 48)

	// Since our cells are all 3 pixels with a 1 pixel barrier
	// around them, we want to make sure our widht/height is
	// a divisor of 4
	gameWidth := int32(float64(displayMode.W) * 0.75)
	gameWidth += gameWidth % 4
	gameHeight := int32(float64(displayMode.H) * 0.75)
	gameHeight += gameWidth % 4

	gamecontroller := controllers.NewLifeGameController(gameWidth, gameHeight)
	if err := gamecontroller.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
