package model

import "time"

// Episode domain object
type Episode struct {
	title       string
	description string
	published   time.Time
	guid        string
	played      bool
	length      int
	url         string
}

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
