package main

const (
	ballWidth     int32 = 25
	ballHeight    int32 = 25
	ballStartingX int32 = 390
	ballStartingY int32 = 290
	ballVelocity  int32 = 6
)

// FIXME: Accept rectParams instead of creating them inside.
func createBall() Rect {
	params := rectParams{
		color:     colorWhite,
		width:     ballWidth,
		height:    ballHeight,
		startingX: ballStartingX,
		startingY: ballStartingY,
	}
	r := NewRect(params)
	r.SetXVelocity(ballVelocity)
	return r
}
