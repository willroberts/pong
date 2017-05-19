package main

// The computer player attempts to keep his paddle aligned vertically with
// the ball. A delay could be added before the adjustment to reduce the
// difficulty.

func (g *game) processAI() {
	ball := g.ball
	opponent := g.paddles[1]

	if opponent.positionY >= ball.positionY+int32(ballHeight) {
		opponent.movingUp = true
	} else {
		opponent.movingUp = false
	}

	if opponent.positionY+int32(paddleHeight) < ball.positionY {
		opponent.movingDown = true
	} else {
		opponent.movingDown = false
	}
}
