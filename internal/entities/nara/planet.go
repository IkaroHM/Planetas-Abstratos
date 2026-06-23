package nara

import "github.com/gopxl/pixel/v2"

const (
	TileSize  = 32
	Scale     = 3
	FinalSize = TileSize * Scale 
)
type World struct {
	Tiles [][]int
}

func InitWorld() World {
	return World{
		Tiles: [][]int{
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, GrassMid, Empty, Empty, Empty, Empty},
			{Empty, GrassLeft, GrassMid, GrassMid, DirtMid, GrassMid, GrassRight, Empty, FloatGrass},
			{Empty, DirtMid,   DirtMid,  DirtMid,  DirtMid,  DirtMid,   DirtMid,    Empty, Empty},
			{Empty, DirtLeft,  DirtMid,  DirtMid,  DirtMid,  DirtMid,   DirtRight,  Empty, Empty},
		},
	}
}

func DrawNara(window pixel.Target, world World, blocks map[int]Block) {
	numRows := len(world.Tiles)

	for y := 0; y < numRows; y++ {
		for x := 0; x < len(world.Tiles[y]); x++ {
			tile := world.Tiles[y][x]
			if tile == Empty {
				continue
			}

			block, exists := blocks[tile]
			if !exists || block.Sprite == nil {
				continue
			}

			screenX := float64(x) * FinalSize
			screenY := float64(numRows-1-y) * FinalSize

			centerOffset := FinalSize / 2
			pos := pixel.V(screenX+float64(centerOffset), screenY+float64(centerOffset))

			mat := pixel.IM.Scaled(pixel.ZV, Scale).Moved(pos)
			block.Sprite.Draw(window, mat)
		}
	}
}
