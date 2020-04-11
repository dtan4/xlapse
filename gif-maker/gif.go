package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Gif struct {
	gif *gif.GIF
}

func NewGif() *Gif {
	return &Gif{
		gif: &gif.GIF{},
	}
}

func (g *Gif) Append(body []byte, delay int) error {
	m, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("cannot decode image: %w", err)
	}

	b := bytes.Buffer{}

	if err := gif.Encode(&b, m, nil); err != nil {
		return fmt.Errorf("cannot encode image as GIF: %w", err)
	}

	i, err := gif.Decode(&b)
	if err != nil {
		return fmt.Errorf("cannot decode GIF-encoded image: %w", err)
	}

	g.gif.Image = append(g.gif.Image, i.(*image.Paletted))
	g.gif.Delay = append(g.gif.Delay, delay)

	return nil
}

func (g *Gif) SaveToFile(name string) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("cannot open %q: %w", name, err)
	}
	defer f.Close()

	if err := gif.EncodeAll(f, g.gif); err != nil {
		return fmt.Errorf("cannot write GIF image to %q: %w", name, err)
	}

	return nil
}
