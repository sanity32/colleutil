package selector

func Parse[T ~string](selector T) *Model {
	result := Model{}
	r := parseToParentless(string(selector))
	result.Selector = r.Selector
	result.Contains = r.Contains
	result.Visibility = r.Visibility
	if p := r.Parent; p.Active {
		result.Parent = Parse(p.Data.Selector)
		result.ParentIsRootwise = p.Data.Rootwise
	}
	return &result
}
