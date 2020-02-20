package spotify

import (
	"encoding/json"
	"regexp"
)

type Artist struct {
	Name string `json:name`
}

type Album struct{}

type Track struct {
	Name    string   `json:name`
	Artists []Artist `json:artists`
	Album   Album    `json:album`
}

type Item struct {
	Track Track `json:track`
}

type T struct {
	Items []Item `json:items`
}

type Spotify struct {
	Tracks T `json:tracks`
}

type S struct{}

func New() *S {
	return &S{}
}

func (s *S) Parse(body []byte) Spotify {

	str := string(body)
	exp := `Spotify.Entity = (?P<json>{.*})`
	data := regexp.MustCompile(exp)

	match := data.FindStringSubmatch(str)
	val := match[1]

	var valU Spotify
	json.Unmarshal([]byte(val), &valU)

	return valU

}
