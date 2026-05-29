package image_processing

import (
	"image"
	"image/draw"
	"os"
	"reflect"
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

func (i ImageHandler) Resize(img *image.NRGBA, dst Dimension) *image.NRGBA {
	src := NewDimensionFromImg(*img)
	if reflect.DeepEqual(src, dst) {
		return i.Clone(img)
	}
}

func (i ImageHandler) Clone(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds()
	dst := image.NewNRGBA(bounds)
	draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
	return dst
	//utils.ParallelTreatment()
}
