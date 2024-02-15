package colleaction

type TimeoutVisualizerOpts struct {
	Id          string `json:"id"`
	CooldownMs  int    `json:"cooldown"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	TimeoutMs   int    `json:"timeout"`
	AccentColor string `json:"accentColor"` //accent-color: coral;
}
