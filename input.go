package main

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	arrowDown sdl.Keycode = 1073741905
	arrowUp   sdl.Keycode = 1073741906
)

func (g *game) processInput() error {
	player := g.paddles[0]
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return errors.New("quitting")
		case *sdl.KeyDownEvent:
			if t.Keysym.Sym != arrowUp && t.Keysym.Sym != arrowDown {
				continue
			}
			if t.Keysym.Sym == arrowUp {
				player.SetYVelocity(-paddleVelocity)
			}
			if t.Keysym.Sym == arrowDown {
				player.SetYVelocity(paddleVelocity)
			}
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym != arrowUp && t.Keysym.Sym != arrowDown {
				continue
			}
			if t.Keysym.Sym == arrowUp {
				player.SetYVelocity(0)
			}
			if t.Keysym.Sym == arrowDown {
				player.SetYVelocity(0)
			}
		}
	}
	return nil
}
