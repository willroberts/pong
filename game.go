package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	surface *sdl.Surface
	paddles []*paddle
	ball    *ball

	playerScore   uint
	opponentScore uint
}

func (g *game) drawPaddle(x, y int32, color uint32) {
	rect := sdl.Rect{
		X: x,
		Y: y,
		W: paddleWidth,
		H: paddleHeight,
	}
	g.surface.FillRect(&rect, color)
}

func (g *game) drawBall(x, y int32, color uint32) {
	rect := sdl.Rect{
		X: x,
		Y: y,
		W: ballWidth,
		H: ballHeight,
	}
	g.surface.FillRect(&rect, color)
}

func (g *game) mainLoop() {
	b := g.ball
	g.processCollision()
	b.move()
	g.drawBall(b.positionX, b.positionY, b.color)

	err := g.processInput()
	if err != nil {
		log.Println("error: ", err)
		sdl.Quit()
	}
	for _, p := range g.paddles {
		p.move()
		g.drawPaddle(p.positionX, p.positionY, p.color)
	}
}
