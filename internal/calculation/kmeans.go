package calculation

import (
	"fmt"
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
	km.Centroids = make([]color.RGBA, 0, km.K)
	km.Centroids = append(km.Centroids, pixels[rand.Intn(len(pixels))])
	nbCentroids := len(km.Centroids)
	for nbCentroids < km.K {
		distances := make([]float64, len(pixels))
		total := 0.0
		for i, p := range pixels {
			minDist := math.MaxFloat64
			for _, c := range km.Centroids {
				if d := km.H.DistanceSq(p, c); d < minDist {
					minDist = d
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

//func (km *KMeans) assign(pixels []color.RGBA) []int {}

func (km *KMeans) update(pixels []color.RGBA, assignments []int) {
	//
}

func (km *KMeans) Fit(pixels []color.RGBA) {
	km.initCentroids(pixels)
}

// func (km *KMeans) Quantize(img image.Image) image.Image {}
