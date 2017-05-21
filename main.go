package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/willroberts/pong/pong"
)

const (
	windowTitle  string = "Pong"
	windowWidth  int    = 800
	windowHeight int    = 600
	frameTime    uint32 = 16 // 62.5 FPS.
)

func init() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal("failed to initialize sdl: ", err)
	}
}

func main() {
	// Create the game window and drawable surface.
	window, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal("failed to initialize window: ", err)
	}
	defer window.Destroy()
	surface, err := window.GetSurface()
	if err != nil {
		log.Fatal("failed to initialize surface: ", err)
	}

	// Get a pointer to the renderer for screen clearing.
	renderer, err := sdl.CreateRenderer(window, -1,
		sdl.RENDERER_SOFTWARE)
	if err != nil {
		log.Fatal("failed to initialize renderer: ", err)
	}

	// Initialize the game engine.
	g, err := pong.Setup(windowWidth, windowHeight)
	if err != nil {
		log.Fatal("failed to initialize game engine: ", err)
	}
	g.Surface = surface
	defer g.Font.Close()

	// Run the game's main loop until we close the window.
	var running = true
	for running {
		err = renderer.Clear()
		if err != nil {
			log.Println("failed to clear the screen: ", err)
		}

		err = g.Loop()
		if err != nil && err.Error() == "quitting" {
			running = false
		}

		err := window.UpdateSurface()
		if err != nil {
			log.Println("failed to update window surface: ", err)
		}

		sdl.Delay(frameTime)
	}

	// Destroy the game window.
	sdl.Quit()
}
