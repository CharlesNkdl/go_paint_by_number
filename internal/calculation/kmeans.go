package calculation

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
)

type KMeans struct {
	K         int
	IterLimit int
	H         Helper
	Centroids []color.RGBA
}

func NewKMeans(k int, iterLimit int) *KMeans {
	return &KMeans{
		K:         k,
		IterLimit: iterLimit,
		H:         Helper{},
	}
}

func (km *KMeans) initCentroids(pixels []color.RGBA) {
	// to init centroids, we take one
	fmt.Println(len(pixels), cap(pixels), pixels[1])
	// here i make a slice containing rgba strucs, with a len of 0  and a cap of the number of k means peak
	km.Centroids = make([]color.RGBA, 0, km.K)
	// here i append , a randomly selected pixels to be the first centroid randomly
	randomPixel := pixels[rand.Intn(len(pixels))]
	km.Centroids = append(km.Centroids, randomPixel)
	// for each of the number of centroids, we will calculate the distance using distance sq
	nbCentroids := len(km.Centroids)
	for nbCentroids < km.K {
		// we make a slice, with unknow capacity, of float 64, with the len of pixels
		distances := make([]float64, len(pixels))
		total := 0.0
		for i, p := range pixels {
			minDist := math.MaxFloat64
			for _, c := range km.Centroids {
				distToCentroid := km.H.DistanceSq(p, c)
				if distToCentroid < minDist {
					minDist = distToCentroid
				}
			}
			distances[i] = minDist
			total += minDist
		}
		threshold := rand.Float64() * total
		cumul := 0.0
		for i, d := range distances {
			cumul += d
			if cumul >= threshold {
				km.Centroids = append(km.Centroids, pixels[i])
				nbCentroids++
				break
			}
		}
	}
}

func (km *KMeans) assign(pixels []color.RGBA) []int {
	assignments := make([]int, len(pixels))
	for i, p := range pixels {
		minDist := math.MaxFloat64
		minIdx := 0
		for j, c := range km.Centroids {
			dist := km.H.DistanceSq(p, c)
			if dist < minDist {
				minDist = dist
				minIdx = j
			}
		}
		assignments[i] = minIdx
	}
	return assignments
}

func (km *KMeans) update(pixels []color.RGBA, assignments []int) {
	sums := make([][3]float64, km.K)
	counts := make([]int, km.K)

	for i, p := range pixels {
		idx := assignments[i]
		sums[idx][0] += float64(p.R)
		sums[idx][1] += float64(p.G)
		sums[idx][2] += float64(p.B)
		counts[idx]++
	}

	for i := range km.Centroids {
		if counts[i] == 0 {
			// if no , centroids is useless so just re randomize it
			km.Centroids[i] = pixels[rand.Intn(len(pixels))]
			continue
		}
		km.Centroids[i] = color.RGBA{
			R: uint8(sums[i][0] / float64(counts[i])),
			G: uint8(sums[i][1] / float64(counts[i])),
			B: uint8(sums[i][2] / float64(counts[i])),
			A: 255,
		}
	}
}

func (km *KMeans) Fit(pixels []color.RGBA) {
	// we do the calculation on only one third of the image too reduce calculation
	sample := km.H.Subsample(pixels, 0.30)

	km.initCentroids(sample)

	for i := 0; i < km.IterLimit; i++ {
		assignments := km.assign(sample)
		km.update(sample, assignments)
	}
}

func (km *KMeans) Quantize(img image.Image) image.Image {
	bounds := img.Bounds()
	out := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			p := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}

			minDist := math.MaxFloat64
			nearest := km.Centroids[0]
			for _, c := range km.Centroids {
				if d := km.H.DistanceSq(p, c); d < minDist {
					minDist = d
					nearest = c
				}
			}

			out.SetNRGBA(x, y, color.NRGBA{
				R: nearest.R,
				G: nearest.G,
				B: nearest.B,
				A: p.A,
			})
		}
	}
	return out
}
