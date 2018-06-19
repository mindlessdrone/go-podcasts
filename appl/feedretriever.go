package appl

// interface for getting feed data
type FeedRetriever interface {
	GrabData(url string) (string, error)
}
