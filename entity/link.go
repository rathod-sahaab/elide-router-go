package entity

import "time"

type Link struct {
	ID        ID
	Slug      string
	Target    string
	CreatorID ID
	GroupID   ID
	ProjectID ID
	CreatedAt time.Time
	UpdatedAt time.Time
}
