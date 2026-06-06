package engine

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/gopxl/pixel/v2"
)

func LoadAssets(path string) (pixel.Picture, error){
	file, err := os.Open(path)
	if (err != nil){
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if (err != nil){
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func CreateSprite(path string) (*pixel.Sprite) {
	picture, err := LoadAssets(path)
	if (err != nil) {
		log.Fatal(err)
	}

	sprite := pixel.NewSprite(picture, pixel.R(0, 0, 32, 32))

	return sprite
}