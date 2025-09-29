package main

import (
	"bufio"
	"context"
	"fmt"
	appsvc "go-ddd/application/user"
	domain "go-ddd/domain/user"
	infra "go-ddd/infra/inmemory/user"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func startup() *appsvc.Service {
	repo := infra.NewRepository()
    domSvc := domain.New(repo)
    return appsvc.New(repo, domSvc)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := startup()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input your name :")
		fmt.Print("> ")

		input, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("\nbye。")

				return 
			}

			log.Printf("read error: %v", err)
			continue
		}

		input = strings.TrimSpace(input)
		
		if input == "" {
			continue
		}

		if err := app.Register(ctx, appsvc.UserRegisterCommand{Name: input}); err != nil {
			log.Printf("failed to register user: %v", err)
		} else {
			fmt.Println("-------------------------")
			fmt.Println("user created:")
			fmt.Println("-------------------------")
			fmt.Println("user name:")
			fmt.Println("- " + input)
			fmt.Println("-------------------------")
		}

		fmt.Println("Continue? (y/n)")
		fmt.Print("> ")

		ans, err := reader.ReadString('\n')

		if err != nil {           
			if err == io.EOF {    
				fmt.Println("\nbye.")
				return
			}
			log.Printf("read error: %v", err) 
			continue
		}
		if strings.ToLower(strings.TrimSpace(ans)) == "n" {
			break
		}
	}
}