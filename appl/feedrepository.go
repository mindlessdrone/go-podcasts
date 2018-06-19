package appl

import "github.com/mindlessdrone/go-podcasts/model"

type FeedRepository interface {
	Query(id int) (*model.Feed, error)
	QueryAll() ([]model.Feed, error)
	Add(feed *model.Feed) error
	Update(id int, feed *model.Feed) error
	SetEpisodePlayed(id int, guid string, played bool) error
}

type DefaultRepository struct{}

func (repo DefaultRepository) Query(id int) (*model.Feed, error) {
	return nil, nil
}

func (repo DefaultRepository) QueryAll() ([]model.Feed, error) {
	return nil, nil
}

func (repo DefaultRepository) Add(feed *model.Feed) error {
	return nil
}

func (repo DefaultRepository) Update(id int, feed *model.Feed) error {
	return nil
}

func (repo DefaultRepository) SetEpisodePlayed(id int, guid string, played bool) error {
	return nil
}
