package main

// The computer player attempts to keep his paddle aligned vertically with
// the ball. A delay could be added before the adjustment to reduce the
// difficulty.

// processAI moves the computer-controlled opponent in response to the ball's
// position.
// A perfect (unbeatable) AI would start adjusting its position as soon as
// the ball has matched its Y range. In order to make the AI defeatable, we
// wait until the ball is out of range before adjusting position.
func (g *game) processAI() {
	ball := g.ball
	opponent := g.paddles[1]
	if opponent.Y() > ball.Y()-ballHeight {
		opponent.SetYVelocity(-paddleVelocity)
	} else if opponent.Y()+int32(paddleHeight) < ball.Y() {
		opponent.SetYVelocity(paddleVelocity)
	}
}
