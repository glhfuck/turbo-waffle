package domain

import "time"

type Link struct {
	Id           int       `json:"-"               db:"link_id"`
	OwnerId      int       `json:"owner_id"               db:"owner_id"`
	OriginalURL  string    `json:"original_URL"    db:"original_url"   binding:"required"`
	CreationDate time.Time `json:"creation_date"   db:"creation_date"`
	UpdateDate   time.Time `json:"update_date"     db:"update_date"`
	VisitsCount  int       `json:"visits_count"    db:"visits_count"`
}
