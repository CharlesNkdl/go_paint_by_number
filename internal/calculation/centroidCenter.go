package calculation

type CentroidCenter struct {
	Index int
	X     int
	Y     int
}

// FindCenters is made to find the center of the centroid to put the label
// but it didnt work since i changed to floodfill but keep it for ref
func FindCenters(indexMap [][]int, numColors int) []CentroidCenter {
	sumX := make([]int, numColors+1)
	sumY := make([]int, numColors+1)
	counts := make([]int, numColors+1)

	for y, row := range indexMap {
		for x, idx := range row {
			sumX[idx] += x
			sumY[idx] += y
			counts[idx]++
		}
	}

	centers := make([]CentroidCenter, 0, numColors)
	for i := 1; i <= numColors; i++ {
		if counts[i] == 0 {
			continue
		}
		centers = append(centers, CentroidCenter{
			Index: i,
			X:     sumX[i] / counts[i],
			Y:     sumY[i] / counts[i],
		})
	}
	return centers
}
