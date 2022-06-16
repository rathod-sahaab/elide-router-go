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

func NewLink(slug string, target string, creatorID ID, groupID ID, projectID ID) (*Link, error) {
	link := Link{
		ID:        NewID(),
		Slug:      slug,
		Target:    target,
		CreatorID: creatorID,
		GroupID:   groupID,
		ProjectID: projectID,
		CreatedAt: time.Now(),
	}

	err := link.Validate()

	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (link *Link) Validate() error {
	if link.Slug == "" {
		return ErrInvalidEntity
	}

	if link.Target == "" {
		return ErrInvalidEntity
	}

	return nil
}
