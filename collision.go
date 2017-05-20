package main

func (g *game) processCollision() {
	g.handleWalls()
	g.handleBounces()
}

// handleBounces reverses the ball direction on paddle collision.
func (g *game) handleBounces() {
	for _, p := range g.paddles {
		if p.Rect().HasIntersection(g.ball.Rect()) {
			g.ball.SetXVelocity(g.ball.GetXVelocity() * -1)
			g.ball.SetYVelocity(p.GetYVelocity()) // Apply "spin".
		}
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
		g.Reset()
	}
	if g.ball.X() <= ballWidth {
		g.Reset()
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
