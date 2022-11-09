package golang

import (
	"sort"
)

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) <= 1 {
		return 0
	}

	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})

	lines := 1
	prev := NewPoint(stockPrices[1]).calcSlope(NewPoint(stockPrices[0]))
	for i := 2; i < len(stockPrices); i++ {
		curr := NewPoint(stockPrices[i]).calcSlope(NewPoint(stockPrices[i-1]))
		if curr.Equal(prev) {
			continue
		}

		lines++
		prev = curr
	}

	return lines
}

type Point struct {
	X int
	Y int
}

type Slope struct {
	Y int
	X int
}

func (s2 Slope) Equal(s1 Slope) bool {
	return s1.Y*s2.X == s1.X*s2.Y
}

func NewPoint(p []int) Point {
	return Point{X: p[0], Y: p[1]}
}

func (p2 Point) calcSlope(p1 Point) Slope {
	y := (p2.Y - p1.Y)
	x := (p2.X - p1.X)

	return Slope{Y: y, X: x}
}
