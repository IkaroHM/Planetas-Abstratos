package main

import (
	"log"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/ikarohm/planetas-abstratos/internal/engine"
	"github.com/ikarohm/planetas-abstratos/internal/entities"
	"golang.org/x/image/colornames"
)

func run() {
	window, err := engine.InitWindow()

	if (err != nil){
		log.Fatal(err)
	}

	defer window.Destroy()


	assets := entities.Assets{
		Grass: engine.GetSprite("assets/sprites/Planeta-Inicial.png", 32.0, 0.0, 32.0, 32.0),
		Dirt: engine.GetSprite("assets/sprites/Planeta-Inicial.png", 32.0, 32.0, 32.0, 32.0),
	}

	nara := entities.World{
		Tiles: [][]int{
			{0, 0, 0,0},
			{0, 0, 0, 0},
			{entities.Grass, entities.Grass, entities.Grass, entities.Grass, entities.Grass, entities.Grass, entities.Grass, entities.Grass},

			{entities.Dirt, entities.Dirt, entities.Dirt, entities.Dirt,
			 entities.Dirt, entities.Dirt, entities.Dirt, entities.Dirt},

			{entities.Dirt, entities.Dirt, entities.Dirt, entities.Dirt,
			 entities.Dirt, entities.Dirt, entities.Dirt, entities.Dirt},
		},
	} 
	astronautaPicture, _ := engine.LoadAsset("assets/sprites/walkingAstronauta-Sheet.png")

	astronauta := engine.SliceSpriteSheet(astronautaPicture, 32, 32)

	for !window.Closed() {
		window.Clear(colornames.Cornflowerblue)

		astronauta[0].Draw(window, pixel.IM.Scaled(pixel.ZV, 6).Moved(pixel.V(96, 370)))
		
		for y := range nara.Tiles {
			for x := range nara.Tiles[y] {
				tile := nara.Tiles[y][x]

				tileX := float64(x) * 96
				tileY := float64(len(nara.Tiles)-1-y) * 96

				switch tile {
					case entities.Grass:
						assets.Grass.Draw(window, pixel.IM.Scaled(pixel.ZV, 3).Moved(
							pixel.V(
								tileX + 48,
								tileY + 48,
							),
						))
					case entities.Dirt:
						assets.Dirt.Draw(window, pixel.IM.Scaled(pixel.ZV, 3).Moved(
							pixel.V(
								tileX + 48,
								tileY + 48,
							),
						))
				}
			}
		}

		window.Update()
	}
}

func main(){
	opengl.Run(run)
}