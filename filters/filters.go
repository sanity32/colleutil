package collefilters

import "fmt"

type Filters struct {
	IsVisible     bool   `json:"isVisible,omitempty"`
	IsInvisible   bool   `json:"isInvisible,omitempty"`
	Contains      string `json:"contains,omitempty"`
	ContainsCS    bool   `json:"containsCS,omitempty"` // case sensitive
	NotContains   string `json:"notContains,omitempty"`
	NotContainsCS bool   `json:"notContainsCS,omitempty"` // case sensitive
	Has           string `json:"has,omitempty"`
	HasNot        string `json:"hasNot,omitempty"`
}

func (f Filters) ToString() string {
	var visibility string
	var contains string
	var containsCs string
	var notContains string
	var notContainsCs string
	if f.IsInvisible {
		visibility += "☺"
	}
	if f.IsVisible {
		visibility += "☻"
	}
	if !f.IsVisible && !f.IsInvisible {
		visibility += "☼"
	}
	if f.Contains != "" {
		switch f.ContainsCS {
		case true:
			containsCs = "CONTAINS"
		case false:
			containsCs = "Contains"
		}
		contains = fmt.Sprintf(`[%v="%v"]`, containsCs, f.Contains)
	}
	if f.NotContains != "" {
		switch f.NotContainsCS {
		case true:
			containsCs = "NOT_CONTAINS"
		case false:
			containsCs = "NotContains"
		}
		notContains = fmt.Sprintf(`[%v="%v"]`, notContainsCs, f.NotContains)
	}
	if visibility == "" && contains == "" && notContains == "" {
		return ""
	}
	return "[" + visibility + contains + notContains + "]"
}

func FilterOnlyVisible() Filters {
	return Filters{
		IsVisible:   true,
		IsInvisible: false,
	}
}

func FilterOnlyInisible() Filters {
	return Filters{
		IsVisible:   false,
		IsInvisible: true,
	}
}
