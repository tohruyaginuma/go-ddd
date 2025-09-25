package userrepo

import (
	"context"
	"errors"
	"fmt"

	domain "go-ddd/domain/user"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


var (
	ErrSaveUser = errors.New("failed to save user")
	ErrFindUser = errors.New("failed to find user")
)

type Repository struct {
	pool *pgxpool.Pool
}

func New (pool *pgxpool.Pool) domain.Repository {
    return &Repository{pool: pool}
}

func (ur *Repository) Save(ctx context.Context, u *domain.User) (*domain.User, error) {
	const query = `
		MERGE INTO users Using (SELECT $1::uuid AS id, $2::text AS name) AS data
		ON users.id = data.id
		WHEN MATCHED THEN UPDATE SET name = data.name
		WHEN NOT MATCHED THEN INSERT (id, name) VALUES(data.id, data.name);
	`
	_, err := ur.pool.Exec(ctx, query, u.ID().String(), u.Name().String())

    if err != nil {
        return nil, fmt.Errorf("%w: %v", ErrSaveUser, err)
    }

    return u, nil
}

func (ur *Repository) Find(ctx context.Context, un domain.UserName) (*domain.User, error) {
    const query = "SELECT id, name FROM users WHERE name = $1 LIMIT 1"
    
	var idStr, nameStr string
    err := ur.pool.QueryRow(ctx, query, un.String()).Scan(&idStr, &nameStr)
    if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
        return nil, fmt.Errorf("%w: %v", ErrFindUser, err)
    }

	uID, err := domain.ParseUserID(idStr)
	if err != nil {
        return nil, fmt.Errorf("%w: parse user id: %v", ErrFindUser, err)
    }

	uName, err := domain.NewUserName(nameStr)

	if err != nil {
        return nil, fmt.Errorf("%w: invalid user name: %v", ErrFindUser, err)
    }

    u, err := domain.NewUserWithID(uID, uName)

	if err != nil {
        return nil, fmt.Errorf("%w: rehydrate user: %v", ErrFindUser, err)
    }

    return u, nil
}