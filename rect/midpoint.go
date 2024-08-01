package rect

import (
	"log"
	"math/rand"
)

// const (
// 	DEFAULT_FACTOR_X = .5
// 	DEFAULT_FACTOR_Y = .5
// 	DEFAULT_SAFE_X   = .1
// 	DEFAULT_SAFE_Y   = .1
// )

func NewMidPoint(xy ...float64) MidPoint {
	var x, y float64 = .5, .5
	if len(xy) > 0 {
		x = xy[0]
	}
	if len(xy) > 1 {
		y = xy[1]
	}
	return MidPoint{Factor: [2]float64{x, y}}
}

type MidPoint struct {
	Factor [2]float64 // X,Y [0,1)
}

// Bigger, wider spread
// x=.5 with spX=.2 x -> [.4,.6)
// x=.1		spX=.4 ->	x=[0,.3)
// x=.1		spX=.4	limX=.1	->	x=[.1,.3)
func (m *MidPoint) SpreadFactor(spX, spY float64) *MidPoint {
	rxx := (rand.Float64() - .5)
	ryx := (rand.Float64() - .5)
	rx := rxx * (spX)
	ry := ryx * (spY)
	m.Factor[0] += rx
	m.Factor[1] += ry
	return m
}

func (m *MidPoint) SafeLimit(lx0, ly0, lx1, ly1 float64) *MidPoint {
	if lx0 > lx1 {
		log.Fatalf("safe limit x0 (%v) > x1 (%v)", lx0, lx1)
	}
	if ly0 > ly1 {
		log.Fatalf("safe limit y0 (%v) > y1 (%v)", ly0, ly1)
	}

	if m.Factor[0] < lx0 {
		m.Factor[0] = lx0
	}
	if m.Factor[0] > lx1 {
		m.Factor[0] = lx1
	}
	if m.Factor[1] < ly0 {
		m.Factor[1] = ly0
	}
	if m.Factor[1] > ly1 {
		m.Factor[1] = ly1
	}
	return m
}

// func (o MidPoint) fX() float64 {
// 	return valOrDef(o.Factor[0], DEFAULT_FACTOR_X)
// }

// func (o MidPoint) fY() float64 {
// 	return valOrDef(o.Factor[1], DEFAULT_FACTOR_Y)
// }

// func (o MidPoint) sX() float64 {
// 	return valOrDef(o.SafeBorder[0], DEFAULT_SAFE_X)
// }

// func (o MidPoint) sY() float64 {
// 	return valOrDef(o.SafeBorder[1], DEFAULT_SAFE_Y)
// }

// func (o MidPoint) GetFxFySxSy() (fX, fY, sX, sY float64) {
// 	fX = o.fX()
// 	fY = o.fY()
// 	sX = o.sX()
// 	sY = o.sY()
// 	return
// }

func (rect Rect) GetWHLT(preserveBorder bool) (w, h, l, t float64) {
	w = float64(rect.Width)
	h = float64(rect.Height)
	l = float64(rect.Left)
	t = float64(rect.Top)
	if preserveBorder {
		w += float64(-rect.BorderLeft - rect.BorderRight)
		h += float64(-rect.BorderTop - rect.BorderBottom)
		l += float64(rect.BorderLeft)
		t += float64(rect.BorderTop)
	}
	return
}

func (rect Rect) FindMidPoint(opts MidPoint) (x, y int) {
	w, h, l, t := rect.GetWHLT(true)
	// fX, fY, sX, sY := opts.GetFxFySxSy()

	oX := w * opts.Factor[0]
	oY := h * opts.Factor[1]
	// safeTermX := (0.5 - fX) * (1 - sX) * w
	// safeTermY := (0.5 - fY) * (1 - sY) * h
	// log.Println("l + oX + safeTermX", l, oX, safeTermX)
	x = int(l + oX)
	y = int(t + oY)
	return
}

// func (rect Rect) FindMidPoint(opts MidPoint) (x, y int) {
// 	w, h, l, t := rect.GetWHLT(true)
// 	fX, fY, sX, sY := opts.GetFxFySxSy()

// 	oX := w * fX
// 	oY := h * fY
// 	safeTermX := (0.5 - fX) * (1 - sX) * w
// 	safeTermY := (0.5 - fY) * (1 - sY) * h
// 	log.Println("l + oX + safeTermX", l, oX, safeTermX)
// 	x = int(l + oX + safeTermX)
// 	y = int(t + oY + safeTermY)
// 	return
// }
