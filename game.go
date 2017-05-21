package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// GameEngine stores our global variables and shared components.
type GameEngine struct {
	font    *ttf.Font
	surface *sdl.Surface
	paddles []Rect
	ball    Rect
	score   int
	pause   bool
}

// Setup configures our game engine and populates it with the necessary
// components.
func Setup() (*GameEngine, error) {
	g := &GameEngine{}

	err := ttf.Init()
	if err != nil {
		return &GameEngine{}, err
	}
	font, err := ttf.OpenFont("font/SFPixelate-Bold.ttf", 40)
	if err != nil {
		return &GameEngine{}, err
	}
	g.font = font

	g.paddles = []Rect{
		NewRect(rectParams{
			color:     colorGreen,
			width:     paddleWidth,
			height:    paddleHeight,
			startingX: 50,
			startingY: 230,
		}),
		NewRect(rectParams{
			color:     colorRed,
			width:     paddleWidth,
			height:    paddleHeight,
			startingX: 725,
			startingY: 230,
		}),
	}

	g.ball = NewRect(rectParams{
		color:     colorWhite,
		width:     ballWidth,
		height:    ballHeight,
		startingX: 390,
		startingY: 290,
	})
	g.ball.SetXVelocity(ballVelocity)

	return g, nil
}

// Reset brings components back to their initial states.
func (g *GameEngine) Reset() {
	g.ball.ResetPosition()
	for _, p := range g.paddles {
		p.ResetPosition()
	}
	g.ball.SetXVelocity(ballVelocity)
}

// Loop contains the core game logic which runs once per frame.
func (g *GameEngine) Loop() error {
	// Pause for one second if we just scored.
	if g.pause {
		sdl.Delay(1000)
		g.Reset()
		g.pause = false
	}

	err := g.processInput()
	if err != nil {
		return err
	}
	g.processAI()

	// Process collision and Scoring.
	g.handleWalls()
	g.handleBounces()

	// Process movement.
	g.ball.Move()
	for _, p := range g.paddles {
		p.Move()
	}

	// Draw all objects on the surface.
	err = g.createLabel(&sdl.Rect{X: 690, Y: 0, W: 0, H: 0}, "PONG")
	if err != nil {
		return err
	}
	err = g.createLabel(&sdl.Rect{X: 0, Y: 0, W: 0, H: 0},
		fmt.Sprintf("SCORE: %d", g.score))
	if err != nil {
		return err
	}
	if err = g.drawRect(g.ball); err != nil {
		return err
	}
	for _, p := range g.paddles {
		if err = g.drawRect(p); err != nil {
			return err
		}
	}

	return nil
}

// drawRect is a helper function for drawing any *sdl.Rect on the game window.
func (g *GameEngine) drawRect(r Rect) error {
	err := g.surface.FillRect(r.Rect(), r.Color())
	if err != nil {
		return err
	}
	return nil
}
