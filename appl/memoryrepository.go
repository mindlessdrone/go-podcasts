package appl

import "github.com/mindlessdrone/go-podcasts/model"

type MemoryRepository struct {
	DefaultRepository
	feeds map[int]*model.Feed
}
