package appl

import (
	"io/ioutil"
	"net/http"
)

type GoFeedRetriever struct{}

func (GoFeedRetriever) GrabData(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	feedData := string(data)
	return &feedData, nil
}
