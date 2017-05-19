package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ballWidth    int32 = 25
	ballHeight   int32 = 25
	ballVelocity int32 = 6

	ballStartingX         int32 = 390
	ballStartingY         int32 = 290
	ballStartingVelocityX int32 = 5
	ballStartingVelocityY int32 = 0
)

type ball struct {
	color uint32

	positionX int32
	positionY int32

	velocityX int32
	velocityY int32
}

func (b *ball) move() {
	b.positionX += b.velocityX
	b.positionY += b.velocityY
}

func (b *ball) reset() {
	b.positionX = ballStartingX
	b.positionY = ballStartingY
	b.velocityX = ballStartingVelocityX
	b.velocityY = ballStartingVelocityY
}

func (b *ball) Rect() *sdl.Rect {
	return &sdl.Rect{
		X: b.positionX,
		Y: b.positionY,
		W: ballWidth,
		H: ballHeight,
	}
}
