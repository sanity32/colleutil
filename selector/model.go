package colleselector

import "encoding/json"

type Model struct {
	Selector         string
	Visibility       Trait[bool]
	Contains         Trait[Contains]
	Parent           *Model
	ParentIsRootwise bool
}

func (m Model) JSON() string {
	j, _ := json.MarshalIndent(m, "", "	")
	return string(j)
}

type ParentlessModel struct {
	Selector   string
	Parent     Trait[Parent]
	Visibility Trait[bool]
	Contains   Trait[Contains]
}

func (r ParentlessModel) HasActiveTraits() bool {
	return r.Parent.Active || r.Visibility.Active || r.Contains.Active
}
