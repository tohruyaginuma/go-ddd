package user

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type IUserRepository interface {
	Save(u User) (*User, error)
	Find(un UserName) (*User, error)
}

type userRepogitry struct {}

func NewUserRepository () IUserRepository {
	return userRepogitry{}
}

func (ur userRepogitry) Save(u User) (*User, error) {
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
    
	defer conn.Close(context.Background())
    
    cmd := "INSERT INTO Users (id, name) VALUES ($1, $2)"
    
    _, err = conn.Exec(context.Background(), cmd, u.getId(), u.getName())

    if err != nil {
        log.Printf("Exists check failed: %v", err)
    }
    

    fmt.Printf("Created successfully new User named %s %s\n", u.getId(), u.getName())

	return nil, nil
}

// public void Save(User user)
// {
// 	using (var connection = new SqlConnection(connectionString))
// 	using (var command = connection.CreateCommand())
// 	{
// 		connection.Open();
// 		command.CommandText = @"
// MERGE INTO users
// USING (
// SELECT @id AS id, @name AS name
// ) AS data
// ON users.id = data.id
// WHEN MATCHED THEN
// UPDATE SET name = data.name
// WHEN NOT MATCHED THEN
// INSERT (id, name)
// VALUES (data.id, data.name);
// ";
// 		command.Parameters.Add(new SqlParameter("@id", user.Id.Value));
// 		command.Parameters.Add(new SqlParameter("@name", user.Name.Value));
// 		command.ExecuteNonQuery();
// 	}
// }

func (ur userRepogitry) Find(un UserName) *User {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
    
	defer conn.Close(context.Background())
    
    cmd := "SELECT 1 FROM Users WHERE name = $1"

    var exists int
    err = conn.QueryRow(context.Background(), cmd, user.getName()).Scan(&exists)
    if err == pgx.ErrNoRows {
        return false
    }
    if err != nil {
        log.Printf("Exists check failed: %v", err)
        return nil
    }
    return true
}

// public User Find(UserName userName)
// {
// 	using (var connection = new SqlConnection(connectionString))
// 	using (var command = connection.CreateCommand())
// 	{
// 		connection.Open();
// 		command.CommandText = "SELECT * FROM users WHERE name = @name";
// 		command.Parameters.Add(new SqlParameter("@name", userName.Value));
// 		using (var reader = command.ExecuteReader())
// 		{
// 			if (reader.Read())
// 			{
// 				var id = reader["id"] as string;
// 				var name = reader["name"] as string;

// 				return new User(
// 					new UserId(id),
// 					new UserName(name)
// 				);
// 			}
// 			else
// 			{
// 				return null;
// 			}
// 		}
// 	}
// }