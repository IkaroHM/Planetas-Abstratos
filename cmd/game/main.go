package main

import (
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/ikarohm/planetas-abstratos/internal/engine"
	"golang.org/x/image/colornames"
)

func run() {
	window, err := engine.InitWindow()

	if (err != nil){
		panic(err)
	}

	defer window.Destroy()

	
	for !window.Closed() {
		window.Clear(colornames.Skyblue)

		window.Update()
	}
}

func main(){
	opengl.Run(run)
}