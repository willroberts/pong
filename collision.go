package main

// determineSpin provides a new Y velocity for y2 given its relative position
// to y1. Contact with the top of a paddle sends the ball up, and contact with
// the bottom of a paddle sends the ball down. Contact with outer edges results
// in the highest "spin" values.
func determineSpin(y1, y2 int32) int32 {
	diff := y2 - y1
	if diff > 100 {
		return 4
	} else if diff > 75 {
		return 2
	} else if diff > 50 {
		return 0
	} else if diff > 25 {
		return -2
	} else {
		return -4
	}
}

// handleBounces reverses the ball direction on paddle collision.
func (g *game) handleBounces() {
	player := g.paddles[0]
	opponent := g.paddles[1]

	if g.ball.Rect().HasIntersection(player.Rect()) {
		g.ball.SetXVelocity(ballVelocity)
		g.ball.SetYVelocity(determineSpin(player.Y(), g.ball.Y()))
	}

	if g.ball.Rect().HasIntersection(opponent.Rect()) {
		g.ball.SetXVelocity(-ballVelocity)
		g.ball.SetYVelocity(determineSpin(opponent.Y(), g.ball.Y()))
	}
}

// handleWalls handles collision of the paddles and ball with window edges. On
// contact with left or right walls, modifies score and signals for pause time.
func (g *game) handleWalls() {
	// Prevent the ball from leaving the game window.
	if g.ball.Y() <= 0 || g.ball.Y() >= int32(windowHeight)-ballHeight {
		g.ball.SetYVelocity(g.ball.GetYVelocity() * -1)
	}

	// Prevent paddles from leaving the game window.
	for _, p := range g.paddles {
		if p.Y() <= 0 && p.GetYVelocity() < 0 {
			p.SetYVelocity(0)
		}
		if p.Y() >= int32(windowHeight)-paddleHeight && p.GetYVelocity() > 0 {
			p.SetYVelocity(0)
		}
	}

	// Detect scores by the player or the opponent.
	if g.ball.X() >= int32(windowWidth)-ballWidth {
		g.score++
		g.pause = true
	}
	if g.ball.X() <= ballWidth {
		g.score--
		g.pause = true
	}
}
