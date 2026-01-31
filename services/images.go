package services

import (
	"fmt"
	"log"
	"os"

	"github.com/fogleman/gg"
)

type ImagesService struct{}

func newImagesService() *ImagesService {
	return &ImagesService{}
}

func (s *ImagesService) CreateWordImage(word string, color string) []byte {
	path := fmt.Sprintf(`./%s.png`, word)
	const w = 1000
	const h = 350
	dc := gg.NewContext(w, h)
	switch color {
	case "white":
		dc.SetRGB(1, 1, 1)
	case "red":
		dc.SetRGB(1, 0.8, 0.8)
	case "green":
		dc.SetRGB(0.8, 1, 0.8)
	default:
		dc.SetRGB(1, 1, 1)

	}
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./fonts/font.ttf", 100); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(word, 40, h/2, 0, 0.5)
	dc.SavePNG(path)

	photoBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}

	return photoBytes
}
