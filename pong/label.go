package pong

import (
	"github.com/veandco/go-sdl2/sdl"
)

// createLabel draws text on the game window.
func (g *GameEngine) createLabel(position *sdl.Rect, text string) error {
	label, err := g.Font.RenderUTF8Solid(text,
		sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return err
	}
	defer label.Free()

	err = label.Blit(nil, g.Surface, position)
	if err != nil {
		return err
	}

	return nil
}
