package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sandbox/db/constraints"
	"sandbox/db/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	CONN_STRING = "postgresql://devuser:devpass@localhost:5432/dev"
)

func run(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, CONN_STRING)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		return err
	}

	db := models.New(conn)
	input := models.UserCreateParams{
		ID:       uuid.New(),
		Email:    "customer@site.com",
		Password: pgtype.Text{Valid: false, String: ""},
		Role:     "CUSTOMER",
	}

	createdUser, err := db.UserCreate(ctx, input)
	err = constraints.ProcessConstraintError(err, constraints.UserConstraints)
	if err != nil {
		return err
	}

	// users, err := db.UserList(ctx, models.UserListParams{
	// 	Limit:  10,
	// 	Offset: 0,
	// })
	// if err != nil {
	// 	return err
	// }

	// id, err := uuid.Parse("8864902a-184b-479a-a8a6-dac0710e0d06")
	// if err != nil {
	// 	return err
	// }

	// deletedUser, err := db.UserDelete(ctx, id)
	// if err != nil {
	// 	return err
	// }

	encoded, err := json.Marshal(createdUser)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", encoded)
	return nil
}
func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
