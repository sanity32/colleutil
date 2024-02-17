package selector

type Model struct {
	Selector   string
	Parent     trait[parent]
	Visibility trait[bool]
	Contains   trait[contains]
}

func (r Model) HasActiveTraits() bool {
	return r.Parent.Active || r.Visibility.Active || r.Contains.Active
}

func Parse[T ~string](selector T) Model {
	sel := string(selector)
	sel, parent := parseParent(sel)
	sel, v := parseTraitVisibility(sel)
	sel, contains := parseTraitContains(sel)
	return Model{
		Selector:   sel,
		Parent:     parent,
		Visibility: v,
		Contains:   contains,
	}
}
