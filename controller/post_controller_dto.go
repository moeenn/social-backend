package controller

import (
	"errors"
	"sandbox/db/models"
	"sandbox/db/service"
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
	Content      string  `json:"body"`
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
		Content:      model.Content,
		CreatedById:  model.CreatedByID.String(),
		CommentCount: model.CommentsCount,
		CreatedAt:    model.CreatedAt.Time.String(),
		UpdatedAt:    updatedAt,
	}
}

type PostWithCreatedByResponse struct {
	PostResponse
	CreatedByName string `json:"createdByName"`
}

type PostsListResponse struct {
	Posts      []*PostWithCreatedByResponse `json:"posts"`
	TotalCount int64                        `json:"totalCount"`
}

func ListPostsResponseFromListPostsResult(result *service.ListPostsResult) *PostsListResponse {
	posts := []*PostWithCreatedByResponse{}
	for _, post := range result.Posts {
		var updatedAt *string
		if post.UpdatedAt.Valid {
			val := post.UpdatedAt.Time.String()
			updatedAt = &val
		}

		posts = append(posts, &PostWithCreatedByResponse{
			PostResponse: PostResponse{
				ID:           post.ID.String(),
				Title:        post.Title,
				Content:      post.Content,
				CreatedById:  post.CreatedByID.String(),
				CommentCount: post.CommentsCount,
				CreatedAt:    post.CreatedAt.Time.String(),
				UpdatedAt:    updatedAt,
			},
			CreatedByName: post.CreatedByName,
		})
	}

	return &PostsListResponse{
		Posts:      posts,
		TotalCount: result.TotalCount,
	}
}

func PostWithCreatedByResponseFromPostById(post *models.PostByIdRow) *PostWithCreatedByResponse {
	var updatedAt *string
	if post.UpdatedAt.Valid {
		val := post.UpdatedAt.Time.String()
		updatedAt = &val
	}

	return &PostWithCreatedByResponse{
		PostResponse: PostResponse{
			ID:           post.ID.String(),
			Title:        post.Title,
			Content:      post.Content,
			CreatedById:  post.CreatedByID.String(),
			CommentCount: post.CommentsCount,
			CreatedAt:    post.CreatedAt.Time.String(),
			UpdatedAt:    updatedAt,
		},
		CreatedByName: post.CreatedByName,
	}
}
