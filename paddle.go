package main

import "github.com/veandco/go-sdl2/sdl"

const (
	paddleWidth    int32 = 25
	paddleHeight   int32 = 150
	paddleVelocity int32 = 4 // Pixels per frame.

	playerStartingX   int32 = 50
	playerStartingY   int32 = 235
	opponentStartingX int32 = 725
	opponentStartingY int32 = 235
)

type paddle struct {
	color uint32

	positionX int32
	positionY int32

	movingUp   bool
	movingDown bool
}

func (p *paddle) move() {
	if p.movingUp && p.positionY > 0 {
		p.positionY -= paddleVelocity
	}
	if p.movingDown && p.positionY < (int32(windowHeight)-paddleHeight) {
		p.positionY += paddleVelocity
	}
}

func (p *paddle) Rect() *sdl.Rect {
	return &sdl.Rect{p.positionX, p.positionY, paddleWidth, paddleHeight}
}
