package model

import (
	"time"
)

// Feed domain object
type Feed struct {
	title                string
	description          string
	iconURL              string
	feedURL              string
	author               string
	updated, feedUpdated time.Time
	episodes             []Episode
}

// Title: returns title of feed
func (feed Feed) Title() string {
	return feed.title
}

// returns description of feed
func (feed Feed) Description() string {
	return feed.description
}

func (feed Feed) Icon() string {
	return feed.iconURL
}

func (feed Feed) URL() string {
	return feed.feedURL
}

func (feed Feed) Author() string {
	return feed.author
}

func (feed Feed) GetUpdated() *time.Time {
	return &feed.updated
}

func (feed *Feed) SetUpdated(newTime time.Time) {
	feed.updated = newTime
}

func (feed *Feed) SetFeedUpdated(newTime time.Time) {
	feed.feedUpdated = newTime
}

func (feed Feed) Episodes() []Episode {
	return feed.episodes
}

func (feed *Feed) AddEpisodes(episodes ...Episode) {
	feed.episodes = append(episodes, feed.episodes...)
}

func (feed Feed) UnplayedEpisodes() []Episode {
	episodes := make([]Episode, 0)
	for _, episode := range feed.episodes {
		if !episode.played {
			episodes = append(episodes, episode)
		}
	}
	return episodes
}

func (feed Feed) NumEpisodesUnplayed() int {
	return len(feed.UnplayedEpisodes())
}
