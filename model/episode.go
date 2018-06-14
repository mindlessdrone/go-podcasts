package model

import "time"

// Episode domain object
type Episode struct {
	title       string
	description string
	published   time.Time
	guid        string
	played      bool
	length      int
}
