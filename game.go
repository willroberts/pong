package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerStartingX   int32 = 50
	playerStartingY   int32 = 235
	opponentStartingX int32 = 725
	opponentStartingY int32 = 235
)

type Game struct {
	surface *sdl.Surface
	paddles []*Paddle
	ball    *ball

	playerScore   uint
	opponentScore uint
}

func (g *Game) drawPaddle(x, y int32, color uint32) {
	rect := sdl.Rect{x, y, paddleWidth, paddleHeight}
	g.surface.FillRect(&rect, color)
}

func (g *Game) drawBall(x, y int32, color uint32) {
	rect := sdl.Rect{x, y, ballWidth, ballHeight}
	g.surface.FillRect(&rect, color)
}

func (g *Game) mainLoop() {
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

func (g *Game) processCollision() {
	ballRect := g.ball.Rect()
	if ballRect == nil {
		log.Fatal("error: ball rect is nil")
	}
	for _, p := range g.paddles {
		paddleRect := p.Rect()
		if paddleRect == nil {
			log.Fatal("error: paddle rect is nil")
		}
		if ballRect.HasIntersection(paddleRect) {
			g.ball.velocityX *= -1
		}
	}
}
