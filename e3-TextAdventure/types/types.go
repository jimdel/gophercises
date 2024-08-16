package types

type Story struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Chapter `json:"options"`
}

type Chapter struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type CYOAGameConfig map[string]Story
