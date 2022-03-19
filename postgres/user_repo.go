package postgres

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"

	"github.com/voodoostack/fitstackapi"
)

type UserRepo struct {
	DB *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (ur *UserRepo) Create(ctx context.Context, user fitstackapi.User) (fitstackapi.User, error) {
	tx, err := ur.DB.Pool.Begin(ctx)
	if err != nil {
		return fitstackapi.User{}, fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	user, err = createUser(ctx, tx, user)
	if err != nil {
		return fitstackapi.User{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return fitstackapi.User{}, fmt.Errorf("error commiting: %v", err)

	}
	return user, nil
}

func createUser(ctx context.Context, tx pgx.Tx, user fitstackapi.User) (fitstackapi.User, error) {
	query := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING *;`

	u := fitstackapi.User{}

	if err := pgxscan.Get(ctx, tx, &u, query, user.Email, user.Username, user.Password); err != nil {
		return fitstackapi.User{}, fmt.Errorf("error insert: %v", err)
	}
	return u, nil
}

func (ur *UserRepo) GetByUsername(ctx context.Context, username string) (fitstackapi.User, error) {
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1;`

	u := fitstackapi.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, username); err != nil {
		if pgxscan.NotFound(err) {
			return fitstackapi.User{}, fitstackapi.ErrNotFound
		}

		return fitstackapi.User{}, fmt.Errorf("error select: %v", err)
	}
	return u, nil
}

func (ur *UserRepo) GetByEmail(ctx context.Context, email string) (fitstackapi.User, error) {
	query := `SELECT * FROM users WHERE email = $1 LIMIT 1;`
	u := fitstackapi.User{}

	if err := pgxscan.Get(ctx, ur.DB.Pool, &u, query, email); err != nil {
		if pgxscan.NotFound(err) {
			return fitstackapi.User{}, fitstackapi.ErrNotFound
		}
		return fitstackapi.User{}, fmt.Errorf("error select: %v", err)
	}
	return u, nil
}
