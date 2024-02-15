package rect

const (
	DEFAULT_FACTOR_X = .5
	DEFAULT_FACTOR_Y = .5
	DEFAULT_SAFE_X   = .1
	DEFAULT_SAFE_Y   = .1
)

type MidPoint struct {
	Factor         [2]float64
	SafeBorder     [2]float64
	PreserveBorder bool
}

func (o MidPoint) fX() float64 {
	return valOrDef(o.Factor[0], DEFAULT_FACTOR_X)
}

func (o MidPoint) fY() float64 {
	return valOrDef(o.Factor[1], DEFAULT_FACTOR_Y)
}

func (o MidPoint) sX() float64 {
	return valOrDef(o.SafeBorder[0], DEFAULT_SAFE_X)
}

func (o MidPoint) sY() float64 {
	return valOrDef(o.SafeBorder[1], DEFAULT_SAFE_Y)
}

func (o MidPoint) GetFxFySxSy() (fX, fY, sX, sY float64) {
	fX = o.fX()
	fY = o.fY()
	sX = o.sX()
	sY = o.sY()
	return
}

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
	fX, fY, sX, sY := opts.GetFxFySxSy()

	oX := w * fX
	oY := h * fY
	safeTermX := (0.5 - fX) * (1 - sX) * w
	safeTermY := (0.5 - fY) * (1 - sY) * h
	x = int(l + oX + safeTermX)
	y = int(t + oY + safeTermY)
	return
}
