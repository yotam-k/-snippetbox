package models

import (
	"errors"
	"time"
)

// Provides a datastore-agnostic way to reflect no matching records
var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID 		int
	Title 	string
	Content string
	Created time.Time
	Expires time.Time
}