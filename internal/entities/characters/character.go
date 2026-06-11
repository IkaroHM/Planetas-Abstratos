package entities

import "github.com/gopxl/pixel/v2"

type Character struct{
	Sprite *pixel.Sprite
	Physic int
	Health int
	Speed int
}