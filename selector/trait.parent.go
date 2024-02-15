package colleselector

var parentSeps = TraitMap[bool]{
	false: {"->"},
	true:  {"<-"},
}

func parseParent(sel string) (base string, result Trait[Parent]) {
	base = sel
	if found, idx, rootwise, token := parentSeps.findLastOccurance(base); found {
		result.Active = true
		result.Data.Selector = sel[:idx]
		result.Data.Rootwise = rootwise
		base = sel[idx+len(token):]
	}
	return
}
