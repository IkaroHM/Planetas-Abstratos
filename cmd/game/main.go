package main

import (
	"log"
	"time"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/ikarohm/planetas-abstratos/internal/engine"
	"github.com/ikarohm/planetas-abstratos/internal/entities/characters"
	"github.com/ikarohm/planetas-abstratos/internal/entities/nara"
	"github.com/ikarohm/planetas-abstratos/internal/physics"
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

func run() {
	window, err := engine.InitWindow()

	if (err != nil){
		log.Fatal(err)
	}

	defer window.Destroy()
	
	world := nara.InitWorld()
	blocks := initBlocks()
	player := &characters.Character{}
	player.Body = physics.NewBody(pixel.V(192, 2000), pixel.R(-48, -48, 48, 48))
	player.Sprite = engine.GetSprite("assets/sprites/walkingAstronauta-Sheet.png", 0, 0, 32, 32 )
	lastFrame := time.Now()
	dt := time.Since(lastFrame).Seconds()

	drawPlayer := func(window pixel.Target, player *characters.Character, dt float64) {
		if glWindow, ok := window.(*opengl.Window); ok{
			player.Body.Update_and_move(dt, blocks, world, glWindow)
		}else{
			log.Fatal("A janela nao pode ser convertida pra *opengl.Window")
		}

		player.Sprite.Draw(window, pixel.IM.Scaled(pixel.ZV, 4).Moved(player.Body.Pos))
	}

	for !window.Closed() {
		dt = time.Since(lastFrame).Seconds()

		lastFrame = time.Now()

		window.Clear(colornames.Darkslategray)
		drawPlayer(window, player, dt)
		nara.DrawNara(window, world, blocks)
		

		window.Update()
	}
}

func main(){
	opengl.Run(run)
}