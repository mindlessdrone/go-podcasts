package appl

import (
	"github.com/mindlessdrone/go-podcasts/model"
	"github.com/mmcdole/gofeed"
)

type FeedServices struct {
	feedRetriever  FeedRetriever
	feedRepository FeedRepository
}

func NewFeedServices(feedRetriever FeedRetriever, feedRepository FeedRepository) FeedServices {
	return FeedServices{
		feedRetriever, feedRepository,
	}
}

func (feedServices FeedServices) AddFeed(url string) error {

	// try to grab feed data
	feedData, err := feedServices.feedRetriever.GrabData(url)
	if err != nil {
		return err
	}

	feedParser := gofeed.NewParser()

	// try to parse the data
	_, err = feedParser.ParseString(feedData)
	if err != nil {
		return err
	}

	return nil
}

func addEpisodesToFeed(feed *model.Feed, items []gofeed.Item) {

}
