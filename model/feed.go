package model

import (
	"time"
)

// Feed domain object
type Feed struct {
	Title                string
	Description          string
	IconURL              string
	FeedURL              string
	Author               string
	Updated, FeedUpdated time.Time
	Episodes             []Episode
}

func NewFeed(title, description, iconURL, feedURL, author string, feedUpdated time.Time) Feed {
	return Feed{
		title,
		description,
		iconURL,
		feedURL,
		author,
		time.Now(),
		feedUpdated,
		make([]Episode, 0),
	}
}

/*// Title: returns title of feed
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

func (feed Feed) GetFeedUpdated() *time.Time {
	return &feed.feedUpdated
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
*/
func (feed *Feed) AddEpisodes(episodes ...Episode) {
	feed.Episodes = append(episodes, feed.Episodes...)
}

func (feed Feed) UnplayedEpisodes() []Episode {
	episodes := make([]Episode, 0)
	for _, episode := range feed.Episodes {
		if !episode.Played {
			episodes = append(episodes, episode)
		}
	}
	return episodes
}

func (feed Feed) NumEpisodesUnplayed() int {
	return len(feed.UnplayedEpisodes())
}
