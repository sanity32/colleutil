package rect

import "errors"

var ErrRectIsVoid = errors.New("rect is void")

func IsRateValue(v float64) bool {
	return v <= 1 && v > 0
}

func valOrDef(val float64, def float64) float64 {
	if val == 0 {
		return val
	}
	return def
}
