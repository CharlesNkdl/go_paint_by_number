package image_processing

import "image"

type Dimension struct {
	w int
	h int
}

func NewDimensionFromImg(img image.NRGBA) Dimension {
	return Dimension{
		w: img.Bounds().Dx(),
		h: img.Bounds().Dy(),
	}
}
