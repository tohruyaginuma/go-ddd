package user_test

import (
	"context"
	"testing"

	application "go-ddd/application/user"
	domain "go-ddd/domain/user"
	inmemory "go-ddd/internal/testrepo/user"
)

const userName = "test"

func TestCreateUser(t *testing.T) {
	t.Parallel()
	
	ctx := context.Background()
	repo := inmemory.New()

	domainSvc := domain.New(repo)
	appSvc := application.New(repo, domainSvc)

	u, err :=appSvc.CreateUser(ctx, userName)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if u == nil {
		t.Fatal("CreateUser returned nil user")
	}

	found, err := repo.Find(ctx, u.Name())
	if err != nil {
		t.Fatalf("repo.Find error: %v", err)
	}
	if found == nil {
		t.Fatalf("user not found in repo after CreateUser")
	}

	if got, want := found.Name().String(), userName; got != want {
		t.Errorf("saved user name mismatch: got %q, want %q", got, want)
	}
}