package action

type LocationInfoResp struct {
	AncestorOrigins map[int]string `json:"ancestorOrigins"`
	Href            string         `json:"href"`
	Origin          string         `json:"origin"`
	Protocol        string         `json:"protocol"`
	Host            string         `json:"host"`
	Hostname        string         `json:"hostname"`
	Port            string         `json:"port"`
	Pathname        string         `json:"pathname"`
	Search          string         `json:"search"`
	Hash            string         `json:"hash"`
}
