package appl

import (
	"io/ioutil"
	"net/http"
)

type HTTPFeedRetriever struct{}

func (HTTPFeedRetriever) GrabData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	feedData := string(data)
	return feedData, nil
}
