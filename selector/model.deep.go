package selector

import "encoding/json"

type DeepModel struct {
	Selector         string
	Visibility       trait[bool]
	Contains         trait[contains]
	Parent           *DeepModel
	ParentIsRootwise bool
}

func (m DeepModel) JSON() string {
	j, _ := json.MarshalIndent(m, "", "	")
	return string(j)
}

func Deep[T ~string](selector T) *DeepModel {
	result := DeepModel{}
	r := Parse(string(selector))
	result.Selector = r.Selector
	result.Contains = r.Contains
	result.Visibility = r.Visibility
	if p := r.Parent; p.Active {
		result.Parent = Deep(p.Data.Selector)
		result.ParentIsRootwise = p.Data.Rootwise
	}
	return &result
}
