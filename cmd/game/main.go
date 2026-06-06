package main

import (
	"log"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/ikarohm/planetas-abstratos/internal/engine"
	"golang.org/x/image/colornames"
)

func run() {
	window, err := engine.InitWindow()

	if (err != nil){
		log.Fatal(err)
	}

	defer window.Destroy()


	astronauta := engine.CreateSprite("assets/sprites/walkingAstronauta-Sheet.png")
	
	for !window.Closed() {
		window.Clear(colornames.Cornflowerblue)

		astronauta.Draw(window, pixel.IM.Scaled(pixel.ZV, 4).Moved(window.Bounds().Center()))
		window.Update()
	}
}

func main(){
	opengl.Run(run)
}