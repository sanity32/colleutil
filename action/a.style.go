package action

type SetStyleRequest struct {
	Styles []StyleRecord `json:"styles"`
}

type StyleRecord struct {
	K    string `json:"k"`
	V    any    `json:"v"`
	Must bool   `json:"important,omitempty"`
}
