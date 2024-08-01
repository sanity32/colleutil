package rect

import (
	"log"
	"testing"
)

func Test1(t *testing.T) {
	r := Rect{
		Top:    0,
		Bottom: 300,
		Right:  300,
		Left:   0,
		X:      0,
		Y:      0,
		Width:  300,
		Height: 300,
		Ok:     true,
	}
	m := MidPoint{Factor: [2]float64{.9, .1}}
	x, y := r.FindMidPoint(m)
	log.Println(x, y)
}
