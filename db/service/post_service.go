package service

import (
	"context"
	"errors"
	"sandbox/db/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PostService struct {
	db *models.Queries
}

func NewPostService(db *models.Queries) *PostService {
	return &PostService{
		db: db,
	}
}

func (s *PostService) CreatePost(ctx context.Context, args models.PostCreateParams) (*models.Post, error) {
	newPost, err := s.db.PostCreate(ctx, args)
	if err != nil {
		return nil, err
	}
	return &newPost, nil
}

func (s *PostService) UpdatePost(ctx context.Context, args models.PostUpdateParams) (*models.Post, error) {
	updatedPost, err := s.db.PostUpdate(ctx, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &updatedPost, nil
}

func (s *PostService) DeletePost(ctx context.Context, postId uuid.UUID) error {
	if err := s.db.PostDelete(ctx, postId); err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("post not found")
		}
		return err
	}

	return nil
}

// TODO: implement post list method.
