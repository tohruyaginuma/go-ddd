package main

import (
	"context"
	"fmt"
	"log"
	"os"

	application "go-ddd/application/user"
	domain "go-ddd/domain/user"
	userrepo "go-ddd/infra/inmemory/user"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
      log.Fatal("Error loading .env file")
    }
}

func main () {
  ctx := context.Background()

  pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

  repo := userrepo.New(pool)
  domSrv := domain.New(repo)
  appSrv := application.New(repo, domSrv)

  usr, err := appSrv.CreateUser(ctx, "Toru")

  if err != nil {
	log.Fatalln("user create failed", err)
  } 

  fmt.Println(usr.Name())
}