package appl

import (
	"strconv"

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
	feed, err := feedParser.ParseString(feedData)
	if err != nil {
		return err
	}

	newFeed := model.NewFeed(
		feed.Title,
		feed.Description,
		feed.Image.URL,
		url,
		feed.Author.Name,
		*feed.UpdatedParsed,
	)

	episodes := itemsToEpisodes(feed.Items)
	newFeed.AddEpisodes(episodes...)
	err = feedServices.feedRepository.Add(&newFeed)
	if err != nil {
		return err
	}
	return nil
}

func (feedServices FeedServices) AllFeeds() ([]*model.Feed, error) {
	feeds, err := feedServices.feedRepository.QueryAll()
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

func (feedServices FeedServices) AllFeedIDS() ([]int, error) {
	ids, err := feedServices.feedRepository.ItemIDs()
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (feedServices FeedServices) RefreshPodcast(id int) error {
	return nil
}

func itemsToEpisodes(items []*gofeed.Item) []model.Episode {
	episodes := make([]model.Episode, 0)

	for _, item := range items {
		length, _ := strconv.Atoi(item.Enclosures[0].Length)
		episodes = append(episodes, model.NewEpisode(
			item.Title,
			item.Description,
			*item.PublishedParsed,
			item.GUID,
			length,
			item.Enclosures[0].URL,
		))
	}

	return episodes
}
