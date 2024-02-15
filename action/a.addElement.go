package colleaction

type Mark string

type AddElementOpts struct {
	NewNode   string            `json:"newNode"`
	Id        string            `json:"id"`
	Prepend   bool              `json:"prepend"` // not append
	Params    map[string]string `json:"params"`
	InnerHTML string            `json:"innerHTML"`
	Mark      Mark              `json:"mark"`
}
