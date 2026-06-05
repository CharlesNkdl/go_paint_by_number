package calculation

import "math"

type CentroidRegion struct {
	ColorIndex int
	CenterX    int
	CenterY    int
}

func FindAllRegions(indexMap [][]int, pixLimit int) []CentroidRegion {
	rows := len(indexMap)
	if rows == 0 {
		return nil
	}
	cols := len(indexMap[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	var regions []CentroidRegion
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if visited[y][x] {
				continue
			}
			colorIdx := indexMap[y][x]
			pixels := floodFill(indexMap, visited, x, y, colorIdx, cols, rows)
			if len(pixels) < pixLimit {
				continue
			}
			cx, cy := bestLabelPosition(pixels, indexMap, colorIdx, cols, rows)
			regions = append(regions, CentroidRegion{
				ColorIndex: colorIdx,
				CenterX:    cx,
				CenterY:    cy,
			})
		}
	}

	return regions
}

func floodFill(indexMap [][]int, visited [][]bool, startX, startY, targetIdx, cols, rows int) [][2]int {
	var pixels [][2]int
	stack := [][2]int{{startX, startY}}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := p[0], p[1]
		if x < 0 || x >= cols || y < 0 || y >= rows {
			continue
		}
		if visited[y][x] || indexMap[y][x] != targetIdx {
			continue
		}
		visited[y][x] = true
		pixels = append(pixels, [2]int{x, y})
		stack = append(stack,
			[2]int{x + 1, y},
			[2]int{x - 1, y},
			[2]int{x, y + 1},
			[2]int{x, y - 1},
		)
	}

	return pixels
}

func centroid(pixels [][2]int) (int, int) {
	sx, sy := 0, 0
	for _, p := range pixels {
		sx += p[0]
		sy += p[1]
	}
	return sx / len(pixels), sy / len(pixels)
}

// centroidInRegion is the evolution of centroid since, if the zone is large, it sometimes put
// the label into another region
// in the end sometimes the number was too close anyway so another method is our go too
func centroidInRegion(pixels [][2]int, indexMap [][]int, targetIdx int) (int, int) {
	cx, cy := centroid(pixels)
	if indexMap[cy][cx] == targetIdx {
		return cx, cy
	}
	minDist := math.MaxFloat64
	bestX, bestY := pixels[0][0], pixels[0][1]

	for _, p := range pixels {
		dx := float64(p[0] - cx)
		dy := float64(p[1] - cy)
		dist := dx*dx + dy*dy
		if dist < minDist {
			minDist = dist
			bestX, bestY = p[0], p[1]
		}
	}

	return bestX, bestY
}

// bestLabelPosition returns the point inside the region
// that is furthest from any border pixel since its hard on calculation we use a subsample
func bestLabelPosition(pixels [][2]int, indexMap [][]int, targetIdx int, cols, rows int) (int, int) {
	candidates := pixels
	if len(pixels) > 300 {
		step := len(pixels) / 300
		candidates = make([][2]int, 0, 300)
		for i := 0; i < len(pixels); i += step {
			candidates = append(candidates, pixels[i])
		}
	}

	maxDist := -1
	bestX, bestY := pixels[0][0], pixels[0][1]

	for _, p := range candidates {
		dist := minDistToBorder(p[0], p[1], indexMap, targetIdx, cols, rows)
		if dist > maxDist {
			maxDist = dist
			bestX, bestY = p[0], p[1]
		}
	}

	return bestX, bestY
}

func minDistToBorder(x, y int, indexMap [][]int, targetIdx, cols, rows int) int {
	type point struct{ x, y, dist int }

	visited := make(map[[2]int]bool)
	queue := []point{{x, y, 0}}
	visited[[2]int{x, y}] = true

	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nx, ny := cur.x+d[0], cur.y+d[1]

			if nx < 0 || nx >= cols || ny < 0 || ny >= rows {
				return cur.dist
			}
			if indexMap[ny][nx] != targetIdx {
				return cur.dist
			}

			k := [2]int{nx, ny}
			if !visited[k] {
				visited[k] = true
				queue = append(queue, point{nx, ny, cur.dist + 1})
			}
		}
	}

	return 0
}
