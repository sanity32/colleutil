package selector

import "strings"

var traitContains = traitMap[bool]{
	false: {":contains("},
	true:  {":containsCs("},
}

type contains struct {
	Contains      string
	CaseSensitive bool
}

func parseTraitContains(sel string) (stripped string, result trait[contains]) {
	stripped = sel
	if found, cs, sep := traitContains.findAny(sel); found {
		a := strings.SplitN(sel, sep, 2)
		strippedBefore, rest := a[0], a[1]
		a = strings.SplitN(rest, ")", 2)
		content, strippedAfter := a[0], a[1]
		stripped = strippedBefore + strippedAfter

		result.Active = true
		result.Data.Contains = content
		result.Data.CaseSensitive = cs
	}
	return
}
