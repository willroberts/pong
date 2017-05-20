package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	surface *sdl.Surface
	paddles []Rect
	ball    Rect
	score   int
	pause   bool
}

func (g *game) drawRect(r Rect) error {
	err := g.surface.FillRect(r.Rect(), r.Color())
	if err != nil {
		return err
	}
	return nil
}

func (g *game) Reset() {
	g.ball.ResetPosition()
	for _, p := range g.paddles {
		p.ResetPosition()
	}
	g.ball.SetXVelocity(ballVelocity)
}

func (g *game) mainLoop() error {
	// Pause for one second if we just scored.
	if g.pause {
		sdl.Delay(1000)
		g.Reset()
		g.pause = false
	}

	// Input.
	err := g.processInput()
	if err != nil {
		return err
	}
	g.processAI()

	// Collision and Scoring.
	g.processCollision()

	// Movement.
	g.ball.Move()
	for _, p := range g.paddles {
		p.Move()
	}

	// Drawing.
	err = g.createLabel(&sdl.Rect{X: 690, Y: 0, W: 0, H: 0}, "PONG")
	if err != nil {
		return err
	}

	err = g.createLabel(&sdl.Rect{X: 0, Y: 0, W: 0, H: 0},
		fmt.Sprintf("SCORE: %d", g.score))
	if err != nil {
		return err
	}

	g.drawRect(g.ball)
	for _, p := range g.paddles {
		g.drawRect(p)
	}

	return nil
}
