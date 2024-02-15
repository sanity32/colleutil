package colleselector

var parentSeps = traitMap[bool]{
	false: {"->"},
	true:  {"<-"},
}

type parent struct {
	Selector string
	Rootwise bool
}

func parseParent(sel string) (base string, result trait[parent]) {
	base = sel
	if found, idx, rootwise, token := parentSeps.findLastOccurance(base); found {
		result.Active = true
		result.Data.Selector = sel[:idx]
		result.Data.Rootwise = rootwise
		base = sel[idx+len(token):]
	}
	return
}
