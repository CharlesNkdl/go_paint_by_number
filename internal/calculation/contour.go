package calculation

import (
	"image"
	"image/color"
)

type Contour struct {
	Mask [][]bool
}

func NewContour(img image.Image) *Contour {
	bounds := img.Bounds()
	mask := make([][]bool, bounds.Dy())
	for i := range mask {
		mask[i] = make([]bool, bounds.Dx())
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if isContour(img, x, y) {
				mask[y-bounds.Min.Y][x-bounds.Min.X] = true
			}
		}
	}

	return &Contour{Mask: mask}
}

func (c *Contour) Render(img image.Image) image.Image {
	bounds := img.Bounds()
	out := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if c.Mask[y-bounds.Min.Y][x-bounds.Min.X] {
				out.SetNRGBA(x, y, color.NRGBA{R: 0, G: 0, B: 0, A: 255})
			} else {
				out.SetNRGBA(x, y, color.NRGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}
	return out
}

func isContour(img image.Image, x, y int) bool {
	bounds := img.Bounds()
	current := img.At(x, y)
	neighbors := [4][2]int{
		{x, y - 1}, // up
		{x, y + 1}, // down
		{x - 1, y}, // left
		{x + 1, y}, // right
	}
	for _, n := range neighbors {
		nx, ny := n[0], n[1]
		if nx < bounds.Min.X || nx >= bounds.Max.X || ny < bounds.Min.Y || ny >= bounds.Max.Y {
			return true
		}
		if img.At(nx, ny) != current {
			return true
		}
	}
	return false
}
