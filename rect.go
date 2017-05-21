package main

import "github.com/veandco/go-sdl2/sdl"

// Rect is the common interface used by the ball and paddles. It is based on
// *sdl.Rect, with additional fields related to game logic.
type Rect interface {
	Color() uint32
	X() int32
	Y() int32

	GetXVelocity() int32
	GetYVelocity() int32
	SetXVelocity(int32)
	SetYVelocity(int32)

	Move()
	ResetPosition()
	Rect() *sdl.Rect
}

// rectParams are the attributes needed to create a new Rect.
type rectParams struct {
	color     uint32
	width     int32
	height    int32
	startingX int32
	startingY int32
}

// rect is the implementation of the Rect interface.
type rect struct {
	color     uint32
	width     int32
	height    int32
	startingX int32
	startingY int32

	positionX int32
	positionY int32
	velocityX int32
	velocityY int32
}

func (r *rect) Color() uint32 {
	return r.color
}

func (r *rect) X() int32 {
	return r.positionX
}

func (r *rect) Y() int32 {
	return r.positionY
}

func (r *rect) GetXVelocity() int32 {
	return r.velocityX
}

func (r *rect) GetYVelocity() int32 {
	return r.velocityY
}

func (r *rect) SetXVelocity(v int32) {
	r.velocityX = v
}

func (r *rect) SetYVelocity(v int32) {
	r.velocityY = v
}

func (r *rect) Move() {
	r.positionX += r.velocityX
	r.positionY += r.velocityY
}

func (r *rect) ResetPosition() {
	r.positionX = r.startingX
	r.positionY = r.startingY
	r.velocityX = 0
	r.velocityY = 0
}

func (r *rect) Rect() *sdl.Rect {
	return &sdl.Rect{X: r.positionX, Y: r.positionY, W: r.width, H: r.height}
}

// NewRect creates a new Rect with the given parameters.
func NewRect(p rectParams) Rect {
	return &rect{
		color:     p.color,
		width:     p.width,
		height:    p.height,
		startingX: p.startingX,
		startingY: p.startingY,
		positionX: p.startingX,
		positionY: p.startingY,
	}
}
