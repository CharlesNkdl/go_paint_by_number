package calculation

import (
	"image"
	"image/color"
)

func MorphClose(img image.Image, radius int) image.Image {
	return erode(dilate(img, radius), radius)
}

func dilate(img image.Image, radius int) image.Image {
	bounds := img.Bounds()
	out := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			out.SetNRGBA(x, y, dominantColor(img, x, y, radius, bounds))
		}
	}
	return out
}

func erode(img image.Image, radius int) image.Image {
	bounds := img.Bounds()
	out := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			out.SetNRGBA(x, y, minorityColor(img, x, y, radius, bounds))
		}
	}
	return out
}

func dominantColor(img image.Image, cx, cy, radius int, bounds image.Rectangle) color.NRGBA {
	counts := make(map[color.RGBA]int)

	for dy := -radius; dy <= radius; dy++ {
		for dx := -radius; dx <= radius; dx++ {
			x, y := cx+dx, cy+dy
			if x < bounds.Min.X || x >= bounds.Max.X || y < bounds.Min.Y || y >= bounds.Max.Y {
				continue
			}
			r, g, b, a := img.At(x, y).RGBA()
			c := color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)}
			counts[c]++
		}
	}

	var dominant color.RGBA
	max := 0
	for c, n := range counts {
		if n > max {
			max = n
			dominant = c
		}
	}
	return color.NRGBA{R: dominant.R, G: dominant.G, B: dominant.B, A: dominant.A}
}

func minorityColor(img image.Image, cx, cy, radius int, bounds image.Rectangle) color.NRGBA {
	r, g, b, a := img.At(cx, cy).RGBA()
	current := color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)}

	counts := make(map[color.RGBA]int)
	total := 0

	for dy := -radius; dy <= radius; dy++ {
		for dx := -radius; dx <= radius; dx++ {
			x, y := cx+dx, cy+dy
			if x < bounds.Min.X || x >= bounds.Max.X || y < bounds.Min.Y || y >= bounds.Max.Y {
				continue
			}
			rv, gv, bv, av := img.At(x, y).RGBA()
			c := color.RGBA{R: uint8(rv >> 8), G: uint8(gv >> 8), B: uint8(bv >> 8), A: uint8(av >> 8)}
			counts[c]++
			total++
		}
	}

	if counts[current]*2 < total {
		var dominant color.RGBA
		maxColor := 0
		for c, n := range counts {
			if n > maxColor {
				maxColor = n
				dominant = c
			}
		}
		return color.NRGBA{R: dominant.R, G: dominant.G, B: dominant.B, A: dominant.A}
	}

	return color.NRGBA{R: current.R, G: current.G, B: current.B, A: current.A}
}
