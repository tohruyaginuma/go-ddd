package main

import (
	"context"
	appsvc "go-ddd/application/user"
	domain "go-ddd/domain/user"
	inmemory "go-ddd/infra/inmemory/user"
	router "go-ddd/interfaces/http"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func startup() *appsvc.Service {
	repo := inmemory.NewRepository()
    domSvc := domain.New(repo)
	factory := inmemory.NewUserFactory()
    return appsvc.New(repo, domSvc, factory)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := startup()
	uc := router.NewUserHandler(app)

	e := echo.New()
	router.RegisterRoute(e, uc)

	go func () {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server start error : %v", err)
		}
	} ()

	<-ctx.Done()

	log.Println("shutting down...")

	shutownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := e.Shutdown(shutownCtx); err != nil {
		log.Fatalf("server shutdown error: %v", err)
	}

	log.Println("server stopped gracefully")
}