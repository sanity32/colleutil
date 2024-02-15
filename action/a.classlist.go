package colleaction

type UpdateClasslistOpts struct {
	Action string `json:"action"` //"remove", "toggle" or "add" (DEFAULT)
	Key    string `json:"key"`
}
