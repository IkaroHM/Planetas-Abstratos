package main

import (
	"fmt"
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


	astronauta, _ := engine.LoadAsset("assets/sprites/walkingAstronauta-Sheet.png")
	var spriti pixel.Sprite
	fmt.Println(astronauta, spriti, "Agora estao sendo usadas KKKKKKKKKK")
	for !window.Closed() {
		window.Clear(colornames.Cornflowerblue)

		window.Update()
	}
}

func main(){
	opengl.Run(run)
}