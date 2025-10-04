package user_test

import (
	"context"
	appsvc "go-ddd/application/user"
	domain "go-ddd/domain/user"
	inmem "go-ddd/infra/inmemory/user"
	"testing"
)

const (
	char2 = "12" 
	char3 = "123"
	char20 = "12345678901234567890" 
)

func newApp(t *testing.T) (*appsvc.Service, *inmem.Repository, *inmem.UserFactory) {
	t.Helper()

	repo := inmem.NewRepository()
	dom := domain.New(repo)
	factory := inmem.NewUserFactory()
	app := appsvc.New(repo, dom, factory)

	return app, repo, factory
}

func TestRegister_SuccessMinUserName(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	app, repo, _ := newApp(t)
	
	// minimum length
	name:= char3
	err := app.Register(ctx, appsvc.UserRegisterCommand{Name: name})

	if err != nil {
		t.Fatalf("Register returned error: %v", err)
	}

	un, err := domain.NewUserName(name)

	if err != nil {
		t.Fatalf("NewUserName returned error: %v", err)
	}

	u, err := repo.FindByName(ctx, un)

	if err != nil {
		t.Fatalf("FIndByName error: %v", err)
	}

	if u == nil {
		t.Fatalf("expected user to be saved, got nil")
	}

}

func TestRegister_SuccessMaxUserName(t *testing.T){
	t.Parallel()
	ctx := context.Background()
	name := char20 

	app, repo, _ := newApp(t)

	if err := app.Register(ctx, appsvc.UserRegisterCommand{Name: name}); err != nil {
		t.Fatalf("Register returned error: %v", err)
	}

	un, err := domain.NewUserName(name)

	if err != nil {
		t.Fatalf("NewUserName returned error: %v", err)
	}

	u, err := repo.FindByName(ctx, un)

	if err != nil {
		t.Fatalf("FindByName error: %v", err)
	}

	if u == nil {
		t.Fatalf("expected user to be saved, got nil")
	}
}

func TestRegister_SuccessMinUserName_DuplicateCase(t *testing.T){
	t.Parallel()
	ctx := context.Background()
	app, repo, _ := newApp(t)

	name := char3
	if err := app.Register(ctx, appsvc.UserRegisterCommand{Name: name}); err != nil {
		t.Fatalf("Register returned error: %v", err)
	}
	un, _ := domain.NewUserName(name)
	u, err := repo.FindByName(ctx, un)

	if err != nil {
		t.Fatalf("FindByName error: %v", err)
	}

	if u == nil {
		t.Fatalf("expected user to be saved, got nil")
	}
}

func TestRegister_InvalidUserNameLengthMin(t *testing.T){
	t.Parallel()
	ctx := context.Background()
	name := char2

	app, _, _ := newApp(t)
	if err := app.Register(ctx, appsvc.UserRegisterCommand{Name: name}); err == nil {
		t.Fatalf("Register returned error: %v", err)
	}
}

func TestRegister_InvalidUserNameLengthMax(t *testing.T){
	t.Parallel()
	ctx := context.Background()
	name := char20 + "X"

	app, _, _:= newApp(t)
	if err := app.Register(ctx, appsvc.UserRegisterCommand{Name: name}); err == nil {
		t.Fatalf("expected error for too long name, got nil")
	}
}

func TestRegister_AlreadyExists(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	existing := "test-user"

	// 既存ユーザー作成
	un, err := domain.NewUserName(existing)
	if err != nil {
		t.Fatalf("NewUserName returned error: %v", err)
	}
	app, repo, factory := newApp(t)

	u, err := factory.Create(ctx, un)
	if err := repo.Save(ctx, u); err != nil {
		t.Fatalf("seed Save error: %v", err)
	}
	if err != nil {
		t.Fatalf("UserFactory.Create returned error: %v", err)
	}

	// 同じ名前で登録 → エラーを期待
	if err := app.Register(ctx, appsvc.UserRegisterCommand{Name: existing}); err == nil {
		t.Fatalf("expected error for already existing user, got nil \n%v", err)
	}
}