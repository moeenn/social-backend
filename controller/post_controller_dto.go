package controller

import (
	"errors"
	"sandbox/db/models"
)

type CreateUpdatePostRequestBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (b *CreateUpdatePostRequestBody) Validate() error {
	if b.Title == "" {
		return errors.New("invalid title")
	}

	if b.Body == "" {
		return errors.New("invalid body")
	}

	return nil
}

type PostResponse struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Body         string  `json:"body"`
	CreatedById  string  `json:"createdById"`
	CommentCount int32   `json:"commentCount"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    *string `json:"updatedAt,omitempty"`
}

func PostResponseFromModel(model *models.Post) *PostResponse {
	var updatedAt *string
	if model.UpdatedAt.Valid {
		val := model.UpdatedAt.Time.String()
		updatedAt = &val
	}

	return &PostResponse{
		ID:           model.ID.String(),
		Title:        model.Title,
		Body:         model.Content,
		CreatedById:  model.CreatedByID.String(),
		CommentCount: model.CommentsCount,
		CreatedAt:    model.CreatedAt.Time.String(),
		UpdatedAt:    updatedAt,
	}
}
