package characters

import (

	"github.com/gopxl/pixel/v2"
	"github.com/ikarohm/planetas-abstratos/internal/physics"
)

type Character struct{
	Sprite *pixel.Sprite
	Body *physics.Body
	Health int
}
