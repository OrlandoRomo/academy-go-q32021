package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	urbanDictionaryURL = "https://mashape-community-urban-dictionary.p.rapidapi.com/define"
	rapidapiKeyName    = "x-rapidapi-key"
)

type UrbanDictionary struct {
	ApiURL  string
	Headers map[string]string
}

func NewUrbanDictionary(apiKey string) *UrbanDictionary {
	return &UrbanDictionary{
		ApiURL: urbanDictionaryURL,
		Headers: map[string]string{
			rapidapiKeyName: apiKey,
		},
	}
}

func (u *UrbanDictionary) GetDefinitionsByTerm(term string) ([]*List, error) {
	var list ListDefinition
	url := fmt.Sprintf("%s?term=%s", u.ApiURL, term)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(rapidapiKeyName, u.Headers[rapidapiKeyName])

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return list.List, nil
}

type ListDefinition struct {
	List []*List `json:"list"`
}

// Struct to unmarshal the Urban dictionary response of definitions based on a term
// API For more infomation: https://rapidapi.com/community/api/urban-dictionary/
type List struct {
	Definition  string   `json:"definition"`
	Permalink   string   `json:"permalink"`
	ThumbsUp    int64    `json:"thumbs_up"`
	SoundUrls   []string `json:"sound_urls"`
	Author      string   `json:"author"`
	Word        string   `json:"word"`
	Defid       int64    `json:"defid"`
	CurrentVote string   `json:"current_vote"`
	WrittenOn   string   `json:"written_on"`
	Example     string   `json:"example"`
	ThumbsDown  int64    `json:"thumbs_down"`
}
