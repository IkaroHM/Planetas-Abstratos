package main

import (
	"log"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/ikarohm/planetas-abstratos/internal/engine"
	"github.com/ikarohm/planetas-abstratos/internal/entities/nara"
	"golang.org/x/image/colornames"
)

func initBlocks() (map[int]nara.Block) {

		grass := engine.GetSprite("assets/sprites/nara.png", 32, 0, 32, 32)
		floatGrass := engine.GetSprite("assets/sprites/nara.png", 96, 0, 32, 32)
		grassLeft := engine.GetSprite("assets/sprites/nara.png", 0, 0, 32, 32)
		grassRight := engine.GetSprite("assets/sprites/nara.png", 64, 0, 32, 32)
		dirtMid := engine.GetSprite("assets/sprites/nara.png", 32, 32, 32, 32)
		dirtLeft := engine.GetSprite("assets/sprites/nara.png", 0, 32, 32, 32)
		dirtRight := engine.GetSprite("assets/sprites/nara.png", 64, 32, 32, 32)

	blocks := map[int]nara.Block{
		nara.GrassMid: {Name:"Grass", Sprite: grass, Solid: true}, 
		nara.DirtMid: {Name: "DirtMid", Sprite: dirtMid, Solid: true},
		nara.FloatGrass: {Name: "FloatGrass", Sprite: floatGrass, Solid: true},
		nara.GrassLeft: {Name:"GrassLeft", Sprite: grassLeft, Solid: true},
		nara.GrassRight: {Name:"GrassRight", Sprite: grassRight, Solid: true},
		nara.DirtLeft: {Name:"DirtLeft", Sprite: dirtLeft, Solid: true},
		nara.DirtRight: {Name:"DirtRight", Sprite: dirtRight, Solid: true},
	}

	return blocks
}

func initWorld()(nara.World){
	world := nara.World{
		Tiles: [][]int{
			{nara.Empty, nara.Empty, nara.Empty, nara.Empty, nara.Empty},
			{nara.Empty, nara.Empty, nara.Empty, nara.Empty, nara.Empty},
			{nara.Empty, nara.Empty, nara.Empty, nara.Empty, nara.Empty, },
			{nara.Empty, nara.GrassLeft, nara.GrassMid, nara.GrassMid, nara.GrassMid, nara.GrassMid, nara.GrassRight, nara.Empty, nara.FloatGrass},
			{nara.Empty, nara.DirtMid, nara.DirtMid, nara.DirtMid, nara.DirtMid, nara.DirtMid, nara.DirtMid},
			{nara.Empty, nara.DirtLeft, nara.DirtMid, nara.DirtMid, nara.DirtMid, nara.DirtMid, nara.DirtRight},
		},
	}
	return world
}

func drawNara(window pixel.Target, world nara.World, blocks map[int]nara.Block){
	for y := range world.Tiles {
		for x := range world.Tiles[y] {
			tile := world.Tiles[y][x]
			if (tile != nara.Empty){
				block := blocks[tile]
				sprite := block.Sprite
				tileX := float64(x) * 96
				tileY := float64(len(world.Tiles)-1-y) * 96
				sprite.Draw(window, 
					pixel.IM.Scaled(pixel.ZV, 3).Moved(
								pixel.V(
									tileX + 48,
									tileY + 48,
								),
							))
			}
		}
	}
}

func run() {
	window, err := engine.InitWindow()

	if (err != nil){
		log.Fatal(err)
	}

	defer window.Destroy()
	
	world := initWorld()
	blocks := initBlocks()

	for !window.Closed() {
		window.Clear(colornames.Cornflowerblue)

		drawNara(window, world, blocks)

		window.Update()
	}
}

func main(){
	opengl.Run(run)
}