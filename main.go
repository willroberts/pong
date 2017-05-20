package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle  string = "Pong"
	windowWidth  int    = 800
	windowHeight int    = 600

	frameTime uint32 = 16 // 62.5 FPS.
)

func init() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal("failed to initialize sdl")
	}
}

func main() {
	// Create the game instance.
	g := &game{}

	// Remember to tear down font resources.
	defer font.Close()

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

	// Get a pointer to the renderer.
	renderer, err := sdl.CreateRenderer(window, -1,
		sdl.RENDERER_SOFTWARE)
	if err != nil {
		log.Println("error creating renderer")
	}

	// Main loop, including input and drawing.
	var running = true
	for running {
		err = renderer.Clear()
		if err != nil {
			log.Println("error clearing the screen")
		}
		err = g.mainLoop()
		if err != nil && err.Error() == "quitting" {
			running = false
		}
		err := window.UpdateSurface()
		if err != nil {
			log.Println("error updating window surface")
		}
		sdl.Delay(frameTime)
	}

	// Destroy the game window.
	sdl.Quit()
}
