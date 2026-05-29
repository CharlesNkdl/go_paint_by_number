package image_processing

import (
	"image"
	"os"
)

type ImageHandler struct{}

func (i ImageHandler) Open(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
