package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle  string = "Pong"
	windowWidth  int    = 800
	windowHeight int    = 600

	frameTime uint32 = 16 // 62.5 FPS.
)

func init() {
	sdl.Init(sdl.INIT_EVERYTHING)
}

func main() {
	// Create the game instance.
	g := &game{}

	// Create the game window.
	window, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Get a pointer to the surface.
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	g.surface = surface

	// Initialize paddles and ball.
	g.paddles = []*paddle{
		&paddle{colorGreen, playerStartingX, playerStartingY, false, false},
		&paddle{colorRed, opponentStartingX, opponentStartingY, false, false},
	}
	g.ball = &ball{colorWhite, ballStartingX, ballStartingY,
		ballStartingVelocityX, ballStartingVelocityY}

	// Get a handle to the renderer.
	renderer, err := sdl.CreateRenderer(window, -1,
		sdl.RENDERER_SOFTWARE)

	// Main loop, including input and drawing.
	var running = true
	for running {
		renderer.Clear()
		g.mainLoop()
		window.UpdateSurface()
		sdl.Delay(frameTime)
	}

	// Destroy the game window.
	sdl.Quit()
}
