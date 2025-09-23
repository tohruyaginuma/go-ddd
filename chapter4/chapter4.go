package chapter4

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type UserId string
type UserName string

func init() {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
}

func (id UserId) String() string {
    return string(id)
}

func (name UserName) String() string {
    return string(name)
}

func NewUserId () UserId{
    return UserId(uuid.NewString())
}

func NewUserName (s string) (UserName, error){
    s = strings.TrimSpace(s)

    if s == "" {
        return "", errors.New("username cannot be empty")
    }

    if utf8.RuneCountInString(s) < 3 {
        return "", errors.New("user name must be at least 3 characters")
    }

    return UserName(s), nil
}

type User struct {
	id	UserId
	name UserName
}

func NewUser(name UserName) (User, error){
    return User{
        id: NewUserId(),
        name: name,
    }, nil
}

func (u User) ID() UserId {
    return u.id
}

func (u User) Name() UserName {
    return u.name
}

type UserService struct {}

func (us UserService) Exists(user User) bool {
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
    
	defer conn.Close(context.Background())
    
    cmd := "SELECT 1 FROM Users WHERE id = $1"

    var exists int
    err = conn.QueryRow(context.Background(), cmd, user.ID()).Scan(&exists)
    if err == pgx.ErrNoRows {
        return false
    }
    if err != nil {
        log.Printf("Exists check failed: %v", err)
        return false
    }
    return true
}

func NewUserService() UserService {
    return UserService{}
}

func CreateUser(userName string)  {
    un, err := NewUserName(userName)

    if err != nil {
        panic(err)
    }
    
    user, err := NewUser(un)
    if err != nil {
        panic(err)
    }

    var userService = NewUserService()

    if userService.Exists(user) {
        panic(fmt.Sprintf("%s already exists.", userName))
    }

    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
    
	defer conn.Close(context.Background())
    
    cmd := "INSERT INTO Users (id, name) VALUES ($1, $2)"
    
    _, err = conn.Exec(context.Background(), cmd, user.ID(), user.Name())

    if err != nil {
        log.Printf("Exists check failed: %v", err)
    }
    

    fmt.Printf("Created successfully new User named %s %s\n", user.name, user.id)
}

func Execute() {
    CreateUser("toru")
}