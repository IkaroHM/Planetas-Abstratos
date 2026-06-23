package physics

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/ikarohm/planetas-abstratos/internal/entities/nara"
)

const (
	Gravity = 9.8 * 100
	StartVel = 300
	JumpForce = 500
)

type Body struct{
	Pos pixel.Vec
	Vel pixel.Vec
	Bounds pixel.Rect
	OnGround bool
}

func (b * Body) Update_and_move(dt float64, blocks map[int]nara.Block, worldMat nara.World, window *opengl.Window){
	directionX := 0.0
	b.Vel.X = 0
	b.Vel.Y -= Gravity * dt

	if (window.Pressed(pixel.KeyA)){
		directionX = -1
		b.Vel.X = StartVel * directionX
	}
	

	if (window.Pressed(pixel.KeyD)){
		directionX = 1
		b.Vel.X = StartVel * directionX
	}

	if window.Pressed(pixel.KeySpace) && b.OnGround{
		b.Vel.Y = JumpForce
	}

	newX := b.Pos.X + (b.Vel.X * dt)
	newY := b.Pos.Y + (b.Vel.Y * dt)

	newPosX := pixel.V(newX, b.Pos.Y)
	newPosY := pixel.V(b.Pos.X, newY)

	if !b.CheckColision(blocks, worldMat, newPosX){
		b.Pos.X = newX
	}

	if b.CheckColision(blocks, worldMat, newPosY){
		if b.Vel.Y <= 0{
			b.OnGround = true
		}
		b.Vel.Y = 0
	} else {
		b.Pos.Y = newY
		b.OnGround = false
	}
}

func NewBody(pos pixel.Vec, bounds pixel.Rect) *Body{
	return &Body{
		Pos: pos,
		Vel: pixel.ZV,
		Bounds: bounds,
		OnGround: false,
	}
}

func (b *Body) GetRect() (pixel.Rect){
	return b.Bounds.Moved(b.Pos)
}

func (b *Body) hitBoxPading(pos pixel.Vec) (pixel.Rect){

	fullRect := b.Bounds.Moved(pos)

	paddingX := 25.0
	paddingY := 0.0

	return pixel.R(
        fullRect.Min.X + paddingX,
        fullRect.Min.Y + paddingY,
        fullRect.Max.X - paddingX,
        fullRect.Max.Y - paddingY,
    )

}

func (b *Body) CheckColision(blocks map[int]nara.Block, worldMat nara.World, newPos pixel.Vec) (bool){
	
	newBounds := b.hitBoxPading(newPos)

	minX   := int(newBounds.Min.X / nara.FinalSize)
    maxX  := int(newBounds.Max.X / nara.FinalSize)
    
    maxY := len(worldMat.Tiles) - 1 - int(newBounds.Min.Y / nara.FinalSize)
    minY := len(worldMat.Tiles) - 1 - int(newBounds.Max.Y / nara.FinalSize)

	for y := minY; y <= maxY; y++{
		for x := minX; x <= maxX; x++ {
			if (y >= 0 && y < len(worldMat.Tiles) && x > 0 && x < len(worldMat.Tiles[y])) {
				tileId := worldMat.Tiles[y][x]

				if block, exists := blocks[tileId]; exists && block.Solid{
					return true
				}
			}
		}
	}
	return false
}
