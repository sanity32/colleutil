package rect

import "github.com/sanity32/colleutil/midpt"

func (rect Rect) whlt(preserveBorder bool) (w, h, l, t float64) {
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

func (rect Rect) FindMidPoint(opts *midpt.Point) (x, y int) {
	if opts == nil {
		opts = midpt.Centered()
	}
	w, h, l, t := rect.whlt(true)
	oX := w * opts[0]
	oY := h * opts[1]
	x = int(l + oX)
	y = int(t + oY)
	return
}
