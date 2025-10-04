package user

import (
	"context"
	"database/sql"
	"fmt"

	domain "go-ddd/domain/user"
)

type UserFactory struct {
	db *sql.DB
}

func NewUserFactory(db *sql.DB) *UserFactory {
	return &UserFactory{db: db}
}

func (f *UserFactory) Create(ctx context.Context, name domain.UserName) (*domain.User, error) {
	var seqID int64
	query := "SELECT NEXTVAL('user_seq')"
	err := f.db.QueryRowContext(ctx, query).Scan(&seqID)

	if err != nil {
		return nil, fmt.Errorf("failed to get next user sequence: %w", err)
	}

	id, err := domain.NewUserID(fmt.Sprintf("%d", seqID))
	
	if err != nil {
		return nil, err
	}
	
	return domain.NewUserWithID(id, name), nil
}