package model

import (
	"time"
)

// Feed domain object
type Feed struct {
	title                string
	description          string
	iconURL              string
	feedURL              string
	author               string
	updated, feedUpdated time.Time
	episodes             []Episode
}
