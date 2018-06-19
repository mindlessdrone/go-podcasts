package appl

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mindlessdrone/go-podcasts/model"
)

type FeedRetrieverStub struct{}

func (FeedRetrieverStub) GrabData(url string) (string, error) {
	testTable := make(map[string]string)

	testTable["http://thisoneisinvalid.org/rss"] = "test_data/invalid.txt"

	fileName, present := testTable[url]
	if present {
		file, _ := os.Open(fileName)
		defer file.Close()

		data, _ := ioutil.ReadAll(file)
		return string(data), nil
	} else {
		return "", errors.New("ERROR")
	}
}

type CallTracker struct {
	callCount map[string]int
}

func (mock CallTracker) Called(funcName string) int {
	return mock.callCount[funcName]
}

type SuccessSaveMock struct {
	CallTracker
	DefaultRepository
}

func (mock *SuccessSaveMock) Add(feed *model.Feed) error {
	mock.callCount["Add"]++
	return nil
}

func TestAddFeed(t *testing.T) {
	t.Run("FeedNotExist", func(t *testing.T) {
		feedServices := NewFeedServices(&FeedRetrieverStub{}, &DefaultRepository{})

		err := feedServices.AddFeed("http://thisdoesnotexist.com/rss")
		if err == nil {
			t.Error("FeedServices did not return an error on nonexistent url")
		}
	})

	t.Run("FeedInvalid", func(t *testing.T) {
		feedServices := NewFeedServices(&FeedRetrieverStub{}, &DefaultRepository{})

		err := feedServices.AddFeed("http://thisoneisinvalid.org/rss")
		if err == nil {
			t.Error("FeedServices did not return an error on invalid feed")
		}
	})
}
