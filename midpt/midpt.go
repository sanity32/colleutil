package midpt

import (
	"log"
	"math/rand"
)

func Centered() *Point {
	return &Point{.5, .5}
}

type Point [2]float64

func New(xy ...float64) *Point {
	var x, y float64 = .5, .5
	if len(xy) > 0 {
		x = xy[0]
	}
	if len(xy) > 1 {
		y = xy[1]
	}
	return &Point{x, y}
}

func (m *Point) SpreadFactor(spX, spY float64) *Point {
	m[0] += (rand.Float64() - .5) * spX
	m[1] += (rand.Float64() - .5) * spY
	return m
}

func (m *Point) SafeLimit(lx0, ly0, lx1, ly1 float64) *Point {
	if lx0 > lx1 {
		log.Fatalf("safe limit x0 (%v) > x1 (%v)", lx0, lx1)
	}
	if ly0 > ly1 {
		log.Fatalf("safe limit y0 (%v) > y1 (%v)", ly0, ly1)
	}

	if m[0] < lx0 {
		m[0] = lx0
	}
	if m[0] > lx1 {
		m[0] = lx1
	}
	if m[1] < ly0 {
		m[1] = ly0
	}
	if m[1] > ly1 {
		m[1] = ly1
	}
	return m
}
