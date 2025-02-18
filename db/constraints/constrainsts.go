package constraints

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	UserConstraints = map[string]string{
		"email_unique": "email already exists",
	}
)

func ProcessConstraintError(err error, constraints map[string]string) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		constraint := pgErr.ConstraintName
		found, ok := constraints[constraint]
		if !ok {
			return err
		}
		return errors.New(found)
	}

	return nil
}
