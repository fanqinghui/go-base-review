package one

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(b Point) float64{
	return math.Hypot(p.X-b.X,p.Y-b.Y)
}

