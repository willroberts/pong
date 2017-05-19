package main

import "github.com/veandco/go-sdl2/sdl"

// TODO:
// 	Instead of drawing in weird places, have all objects have a Rect()
//  function which returns its *sdl.Rect for drawing and collision.
//  Use interfaces, too.

const (
	paddleWidth    int32 = 25
	paddleHeight   int32 = 150
	paddleVelocity int32 = 4 // Pixels per frame.
)

type Paddle struct {
	color uint32

	positionX int32
	positionY int32

	movingUp   bool
	movingDown bool
}

func (p *Paddle) move() {
	if p.movingUp && p.positionY > 0 {
		p.positionY -= paddleVelocity
	}
	if p.movingDown && p.positionY < (int32(windowHeight)-paddleHeight) {
		p.positionY += paddleVelocity
	}
}

func (p *Paddle) Rect() *sdl.Rect {
	return &sdl.Rect{p.positionX, p.positionY, paddleWidth, paddleHeight}
}
