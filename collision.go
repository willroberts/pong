package main

func (g *game) processCollision() {
	g.handleWalls()
	g.handleBounces()
}

// determineSpin provides a new Y velocity for y2 given its relative position
// to y1.
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
		spin := determineSpin(player.Y(), g.ball.Y())
		g.ball.SetYVelocity(spin)
	}

	if g.ball.Rect().HasIntersection(opponent.Rect()) {
		g.ball.SetXVelocity(-ballVelocity)
		spin := determineSpin(opponent.Y(), g.ball.Y())
		g.ball.SetYVelocity(spin)
	}
}

// handleWalls handles collision of the paddles and ball with window edges.
func (g *game) handleWalls() {
	// Prevent the ball from leaving the game window.
	if g.ball.Y() <= 0 || g.ball.Y() >= int32(windowHeight)-ballHeight {
		g.ball.SetYVelocity(g.ball.GetYVelocity() * -1)
	}
	if g.ball.X() >= int32(windowWidth)-ballWidth {
		g.score++
		g.pause = true
	}
	if g.ball.X() <= ballWidth {
		g.score--
		g.pause = true
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
}
