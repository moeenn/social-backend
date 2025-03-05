// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	ID              uuid.UUID
	Content         string
	PostID          uuid.UUID
	LikesCount      int32
	CreatedByID     uuid.UUID
	ParentCommentID pgtype.UUID
	HierarchyID     uuid.UUID
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
	DeletedAt       pgtype.Timestamp
}

type CommentLike struct {
	CommentID uuid.UUID
	UserID    uuid.UUID
}

type Post struct {
	ID            uuid.UUID
	Title         string
	Content       string
	CreatedByID   uuid.UUID
	CommentsCount int32
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
	DeletedAt     pgtype.Timestamp
}

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Name      string
	Role      string
	CreatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}
