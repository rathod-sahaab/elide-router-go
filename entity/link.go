package entity

import (
	"time"

	"github.com/go-playground/validator"
)

type Link struct {
	ID        ID
	Slug      string `validate:"required,slug"`
	Target    string `validate:"required,url"`
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
	validate := validator.New()

	validate.RegisterValidation("slug", slugValidator)

	return validate.Struct(link)
}

func slugValidator(fl validator.FieldLevel) bool {
	slug := fl.Field().String()

	// TODO: implement better slug validation, using regexp
	if slug == "" || slug == "api" {
		return false
	}
	return true
}
