package user

import "context"

type Service struct {
    repo Repository
}

func New(repo Repository) Service {
    return Service{
        repo: repo,
    }
}

func (us Service) Exists(ctx context.Context, un UserName) (bool, error) {
    user, err := us.repo.Find(ctx, un);

    if err != nil {
        return false, err
    }

    return user != nil, nil
}