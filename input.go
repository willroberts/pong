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
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return errors.New("quitting")
		case *sdl.KeyDownEvent:
			if t.Keysym.Sym != arrowUp && t.Keysym.Sym != arrowDown {
				continue
			}
			if t.Keysym.Sym == arrowUp {
				g.paddles[0].movingUp = true
			}
			if t.Keysym.Sym == arrowDown {
				g.paddles[0].movingDown = true
			}
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym != arrowUp && t.Keysym.Sym != arrowDown {
				continue
			}
			if t.Keysym.Sym == arrowUp {
				g.paddles[0].movingUp = false
			}
			if t.Keysym.Sym == arrowDown {
				g.paddles[0].movingDown = false
			}
		}
	}
	return nil
}
