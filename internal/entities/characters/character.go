package characters

import (

	"github.com/ikarohm/planetas-abstratos/internal/physics"
)

type Character struct{
	Body *physics.Body
	Health int
}
