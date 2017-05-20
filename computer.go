package main

// The computer player attempts to keep his paddle aligned vertically with
// the ball. A delay could be added before the adjustment to reduce the
// difficulty.

func (g *game) processAI() {
	ball := g.ball
	opponent := g.paddles[1]

	if opponent.Y() >= ball.Y()+int32(ballHeight) {
		opponent.SetYVelocity(-paddleVelocity)
	} else if opponent.Y()+int32(paddleHeight) < ball.Y() {
		opponent.SetYVelocity(paddleVelocity)
	} else {
		opponent.SetYVelocity(0)
	}
}
