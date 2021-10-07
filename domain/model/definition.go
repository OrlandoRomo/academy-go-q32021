package model

// Struct to unmarshal the Urban dictionary response of definitions based on a term
// API For more infomation: https://rapidapi.com/community/api/urban-dictionary/
type List struct {
	Definitions []*Definition `json:"list"`
}

type Definition struct {
	Definition  string   `json:"definition,omitempty"`
	Permalink   string   `json:"permalink,omitempty"`
	ThumbsUp    int64    `json:"thumbs_up,omitempty"`
	SoundUrls   []string `json:"sound_urls,omitempty"`
	Author      string   `json:"author,omitempty"`
	Word        string   `json:"word,omitempty"`
	Defid       int64    `json:"defid,omitempty"`
	CurrentVote string   `json:"current_vote,omitempty"`
	WrittenOn   string   `json:"written_on,omitempty"`
	Example     string   `json:"example,omitempty"`
	ThumbsDown  int64    `json:"thumbs_down,omitempty"`
}
