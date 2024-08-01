package action

import (
	"github.com/vanng822/css"
)

type StyleString string

func (str StyleString) String() string {
	return string(str)
}

func (str StyleString) Parse() map[string]string {
	m := map[string]string{}
	p := css.ParseBlock(str.String())

	for _, rule := range p {
		m[rule.Property] = rule.Value.ParsedText()
	}
	return m
}

type SetStyleRequest struct {
	Styles []StyleRecord `json:"styles"`
}

type StyleRecord struct {
	K    string `json:"k"`
	V    any    `json:"v"`
	Must bool   `json:"important,omitempty"`
}
