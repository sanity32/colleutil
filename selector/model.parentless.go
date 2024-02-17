package selector

type parentlessModel struct {
	Selector   string
	Parent     trait[parent]
	Visibility trait[bool]
	Contains   trait[contains]
}

func (r parentlessModel) HasActiveTraits() bool {
	return r.Parent.Active || r.Visibility.Active || r.Contains.Active
}

func parseToParentless(sel string) parentlessModel {
	sel, parent := parseParent(sel)
	sel, v := parseTraitVisibility(sel)
	sel, contains := parseTraitContains(sel)
	return parentlessModel{
		Selector:   sel,
		Parent:     parent,
		Visibility: v,
		Contains:   contains,
	}
}
