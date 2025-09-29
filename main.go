package main

import (
	"bufio"
	"context"
	"fmt"
	appsvc "go-ddd/application/user"
	domain "go-ddd/domain/user"
	infra "go-ddd/infra/inmemory/user"
	"os"
	"strings"
)

func main() {
    repo := infra.NewRepository()
    dom := domain.New(repo)
    app := appsvc.New(repo, dom)

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input your name :")
		fmt.Print(">")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println(input)
		if input == "" {
			continue
		}

		if err := app.Register(ctx, appsvc.UserRegisterCommand{Name: input}); err != nil {
			fmt.Printf("failed to register user: %v\n", err)
		} else {
			fmt.Println("-------------------------")
			fmt.Println("user created:")
			fmt.Println("-------------------------")
			fmt.Println("user name:")
			fmt.Println("- " + input)
			fmt.Println("-------------------------")
		}

		fmt.Println("Continue? (y/n)")
		fmt.Print(">")

		ans, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(ans)) == "n" {
			break
		}
	}
}