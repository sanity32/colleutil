package rect

type Rect struct {
	Ok           bool `json:"ok,omitempty"`
	Top          int  `json:"top"`
	Right        int  `json:"right"`
	Bottom       int  `json:"bottom"`
	Left         int  `json:"left"`
	Width        int  `json:"width"`
	Height       int  `json:"height"`
	X            int  `json:"x"`
	Y            int  `json:"y"`
	BorderBottom int  `json:"borderBottom"`
	BorderLeft   int  `json:"borderLeft"`
	BorderRight  int  `json:"borderRight"`
	BorderTop    int  `json:"borderTop"`
}

// Not sure about adding borders
func (r *Rect) SumWith(add Rect) {
	r.Top += add.Top + add.BorderTop
	r.Bottom += add.Bottom + add.BorderBottom
	r.Left += add.Left + add.BorderLeft
	r.Right += add.Right + add.BorderRight
	r.X = r.Left
	r.Y = r.Top
}

func (r Rect) IsVoid() bool {
	return r.X == 0 && r.Y == 0 && r.Width == 0 && r.Height == 0
}
