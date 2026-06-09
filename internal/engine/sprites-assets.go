package engine

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/gopxl/pixel/v2"
)

func LoadAsset(path string) (pixel.Picture, error) {
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
	picture, err := LoadAsset(path)
	if (err != nil) {
		log.Fatal(err)
	}

	sprite := pixel.NewSprite(picture, picture.Bounds())

	return sprite
}

func SliceSpriteSheet(picture pixel.Picture, frameW, frameH float64) ([]*pixel.Sprite) {
	pictureW := picture.Bounds().Max.X
	var frames []*pixel.Sprite

	for f := 0.0; f < pictureW; f += frameW{
		sprite := pixel.NewSprite(picture, pixel.R(f, 0, frameW + f, frameH))
		frames = append(frames, sprite)
	}
	return frames
}

func GetSprite(path string, x float64, y float64, width float64, height float64) (*pixel.Sprite) {
	picture, _ := LoadAsset(path)
	sprite := pixel.NewSprite(picture, pixel.R(x, y, x + width, y + height))
	return sprite
}
