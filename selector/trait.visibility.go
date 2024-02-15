package colleselector

import "strings"

var traitV = traitMap[bool]{
	true:  {":visible"},
	false: {":notVisible", ":invisible"},
}

func parseTraitVisibility[T ~string](selector T) (stripped string, result trait[bool]) {
	sel := string(selector)
	stripped = sel
	if found, _, k, sep := traitV.findLastOccurance(sel); found {
		a := strings.SplitN(sel, sep, 2)
		stripped = a[0] + a[1]
		result.Active = true
		result.Data = k
	}
	return
}
