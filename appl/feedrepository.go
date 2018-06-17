package appl

import "github.com/mindlessdrone/go-podcasts/model"

type FeedRepository interface {
	Query(id int) model.Feed
	QueryAll() []model.Feed
	Add(feed model.Feed) error
	Update(id int, fields map[string]interface{}) error
	AddEpisode(id int, episode model.Episode) error
	ToggleEpisodePlayed(id int, guid string) error
}
