package colleselector

func Parse(sel string) *Model {
	result := Model{}
	r := ParseLayer(sel)
	result.Selector = r.Selector
	result.Contains = r.Contains
	result.Visibility = r.Visibility
	if p := r.Parent; p.Active {
		result.Parent = Parse(p.Data.Selector)
		result.ParentIsRootwise = p.Data.Rootwise
	}
	return &result
}

func ParseLayer(sel string) ParentlessModel {
	sel, parent := parseParent(sel)
	sel, v := parseTraitVisibility(sel)
	sel, contains := parseTraitContains(sel)
	return ParentlessModel{
		Selector:   sel,
		Parent:     parent,
		Visibility: v,
		Contains:   contains,
	}
}
