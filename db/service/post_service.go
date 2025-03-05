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

type ListPostsResult struct {
	Posts      []models.PostsListRow
	TotalCount int64
}

func (s *PostService) ListPosts(ctx context.Context, params models.PostsListParams) (*ListPostsResult, error) {
	postsCount, err := s.db.PostsCount(ctx)
	if err != nil {
		return nil, err
	}

	posts, err := s.db.PostsList(ctx, params)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &ListPostsResult{Posts: []models.PostsListRow{}, TotalCount: 0}, nil
		}
		return nil, err
	}

	result := ListPostsResult{
		Posts:      posts,
		TotalCount: postsCount,
	}

	return &result, nil
}

func (s *PostService) GetPostById(ctx context.Context, postId uuid.UUID) (*models.PostByIdRow, error) {
	post, err := s.db.PostById(ctx, postId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}
