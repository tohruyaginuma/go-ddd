package circle

import "context"

type Repository interface {
	Save(ctx context.Context, circle *Circle) error
	FindByID(ctx context.Context, id *CircleID) (Circle, error)
	FindByName(ctx context.Context, name *CircleName) (Circle, error)
}