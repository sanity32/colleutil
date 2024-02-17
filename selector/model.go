package selector

import "encoding/json"

type Model struct {
	Selector         string
	Visibility       trait[bool]
	Contains         trait[contains]
	Parent           *Model
	ParentIsRootwise bool
}

func (m Model) JSON() string {
	j, _ := json.MarshalIndent(m, "", "	")
	return string(j)
}
