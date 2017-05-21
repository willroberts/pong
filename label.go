package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (g *game) createLabel(position *sdl.Rect, text string) error {
	label, err := g.font.RenderUTF8_Solid(text,
		sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return err
	}
	defer label.Free()

	err = label.Blit(nil, g.surface, position)
	if err != nil {
		return err
	}

	return nil
}
