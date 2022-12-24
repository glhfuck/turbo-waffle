package domain

import "time"

type Link struct {
	ShortPath    string
	OriginalURL  string
	CreationDate time.Time
	UpdateDate   time.Time
	VisitsCount  int64
}
