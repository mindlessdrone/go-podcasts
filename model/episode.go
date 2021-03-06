package model

import "time"

// Episode domain object
type Episode struct {
	Title       string
	Description string
	Published   time.Time
	GUID        string
	Played      bool
	Length      int
	URL         string
}

func NewEpisode(title, description string, published time.Time, guid string,
	length int, url string) Episode {

	return Episode{
		title, description, published, guid, false, length, url,
	}
}

/*
func (episode Episode) Title() string {
	return episode.title
}

func (episode Episode) Description() string {
	return episode.description
}

func (episode Episode) Published() *time.Time {
	return &episode.published
}

func (episode Episode) GUID() string {
	return episode.guid
}

func (episode Episode) Played() bool {
	return episode.played
}

func (episode *Episode) SetPlayed() {
	episode.played = true
}

func (episode *Episode) SetUnplayed() {
	episode.played = false
}

func (episode Episode) Length() int {
	return episode.length
}

func (episode Episode) URL() string {
	return episode.url
}
*/
