package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	surface *sdl.Surface
	paddles []*paddle
	ball    *ball
	score   uint
}

func (g *game) drawPaddle(x, y int32, color uint32) {
	rect := sdl.Rect{
		X: x,
		Y: y,
		W: paddleWidth,
		H: paddleHeight,
	}
	err := g.surface.FillRect(&rect, color)
	if err != nil {
		log.Println("error drawing paddle")
	}
}

func (g *game) drawBall(x, y int32, color uint32) {
	rect := sdl.Rect{
		X: x,
		Y: y,
		W: ballWidth,
		H: ballHeight,
	}
	err := g.surface.FillRect(&rect, color)
	if err != nil {
		log.Println("error drawing ball")
	}
}

func (g *game) mainLoop() error {
	// Collision and Scoring.
	g.processCollision()

	// Input.
	err := g.processInput()
	if err != nil {
		return err
	}

	// Movement.
	g.ball.move()
	g.processAI()
	for _, p := range g.paddles {
		p.move()
	}

	// Drawing.
	err = g.createLabel(&sdl.Rect{690, 0, 0, 0}, "PONG")
	if err != nil {
		return err
	}

	err = g.createLabel(&sdl.Rect{0, 0, 0, 0}, fmt.Sprintf("SCORE: %d", g.score))
	if err != nil {
		return err
	}

	g.drawBall(g.ball.positionX, g.ball.positionY, g.ball.color)
	for _, p := range g.paddles {
		g.drawPaddle(p.positionX, p.positionY, p.color)
	}

	return nil
}
