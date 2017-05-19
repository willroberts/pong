package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// TODO: Move all collision logic here.
func (g *game) processCollision() {
	// Detect collision with paddles, and reverse horizontal course on hit.
	for _, p := range g.paddles {
		paddleRect := p.Rect()
		if g.ball.Rect().HasIntersection(paddleRect) {
			g.ball.velocityX *= -1
			if p.movingUp {
				g.ball.velocityY -= paddleVelocity
				if g.ball.velocityY < -6 {
					g.ball.velocityY = -6
				}
			}
			if p.movingDown {
				g.ball.velocityY += paddleVelocity
				if g.ball.velocityY > 6 {
					g.ball.velocityY = 6
				}
			}
		}
	}

	// Detect collision with top and bottom, reversing vertical course on hit.
	if g.ball.positionY <= 0 || g.ball.positionY >= int32(windowHeight)-ballHeight {
		g.ball.velocityY *= -1
	}

	// Detect collision with left and right, counting a score on hit.
	if g.ball.positionX >= int32(windowWidth) {
		log.Println("you scored!")
		sdl.Delay(2000)
		//paddles.reset()
		g.ball.reset()
	}
	if g.ball.positionX <= int32(0) {
		log.Println("opponent scored!")
		sdl.Delay(2000)
		//paddles.reset()
		g.ball.reset()
	}
}
