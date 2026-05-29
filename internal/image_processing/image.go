package image_processing

import (
	"image"
	"image/color"
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
	return i.Clone(img)
}

func (i ImageHandler) Clone(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds()
	dst := image.NewNRGBA(bounds)
	draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
	return dst
	//utils.ParallelTreatment()
}

func (i ImageHandler) ExtractPixels(img image.Image) []color.RGBA {
	bounds := img.Bounds()
	pixels := make([]color.RGBA, 0, bounds.Dx()*bounds.Dy())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels = append(pixels, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}
	return pixels
}
