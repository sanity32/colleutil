package colleselector

import "strings"

var traitV = TraitMap[bool]{
	true:  {":visible"},
	false: {":notVisible", ":invisible"},
}

func parseTraitVisibility(sel string) (stripped string, result Trait[bool]) {
	stripped = sel
	if found, _, k, sep := traitV.findLastOccurance(sel); found {
		a := strings.SplitN(sel, sep, 2)
		stripped = a[0] + a[1]
		result.Active = true
		result.Data = k
	}
	return
}
