package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	fontPath string = "font/SFPixelate-Bold.ttf"
	fontSize int    = 40
)

var (
	font *ttf.Font
)

func init() {
	var err error
	if font == nil {
		if err := ttf.Init(); err != nil {
			log.Fatal(err)
		}

		font, err = ttf.OpenFont(fontPath, fontSize)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// 340,0,0,0 and PONG
func (g *game) createLabel(position *sdl.Rect, text string) error {
	label, err := font.RenderUTF8_Solid(text,
		sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return err
	}
	defer label.Free()

	label.Blit(nil, g.surface, position)
	if err != nil {
		return err
	}

	return nil
}
