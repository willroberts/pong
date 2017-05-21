package main

// processAI moves the computer-controlled opponent in response to the ball's
// vertical position.
// A perfect (unbeatable) AI would start adjusting its position as soon as
// the ball has matched its Y range. In order to make the AI defeatable, we
// wait until the ball is out of range before adjusting position.
func (g *GameEngine) processAI() {
	opponent := g.paddles[1]
	if opponent.Y() > g.ball.Y()-ballHeight {
		opponent.SetYVelocity(-paddleVelocity)
	} else if opponent.Y()+int32(paddleHeight) < g.ball.Y() {
		opponent.SetYVelocity(paddleVelocity)
	}
}
