package calculation

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type LabelingHelper struct{}

func (l LabelingHelper) Render(contourImg image.Image, regions []CentroidRegion) image.Image {
	bounds := contourImg.Bounds()
	out := image.NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := contourImg.At(x, y).RGBA()
			out.SetNRGBA(x, y, color.NRGBA{
				R: uint8(r >> 8), G: uint8(g >> 8),
				B: uint8(b >> 8), A: uint8(a >> 8),
			})
		}
	}
	for _, r := range regions {
		l.drawLabel(out, r.CenterX, r.CenterY, r.ColorIndex)
	}
	return out
}

func (l LabelingHelper) drawLabel(img *image.NRGBA, x, y, index int) {
	label := fmt.Sprintf("%d", index)
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.NRGBA{R: 0, G: 0, B: 0, A: 255}),
		Face: basicfont.Face7x13,
		Dot:  fixed.P(x-3*len(label), y+5),
	}
	d.DrawString(label)
}
