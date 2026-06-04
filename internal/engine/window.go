package engine

import (

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

func InitWindow() (*opengl.Window, error) {
	cfg := opengl.WindowConfig{
		Title: "Abstract Planets",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	}
	window, err := opengl.NewWindow(cfg)

	if (err  != nil) {
		return nil, err
	}

	return window, nil
}