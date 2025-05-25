package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func connectDB() *pgxpool.Pool {
	connString := "postgres://postgres:password@localhost:5432/crud_db?sslmode=disable"

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return dbpool
}

func createUser(db *pgxpool.Pool, name, email string) {
	sql := `INSERT INTO users(name, email) VALUES($1, $2)`
	_, err := db.Exec(context.Background(), sql, name, email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating user: %v\n", err)
		return
	}
	fmt.Println("User created successfully")
}

func readUsers(db *pgxpool.Pool) {
	rows, err := db.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	fmt.Printf("Users: %+v\n", users)
}

func updateUser(db *pgxpool.Pool, id int, newName string) {
	sql := `UPDATE users SET name = $1 WHERE id = $2`
	_, err := db.Exec(context.Background(), sql, newName, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating user: %v\n", err)
		return
	}
	fmt.Println("User updated successfully")
}

func deleteUser(db *pgxpool.Pool, id int) {
	sql := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(context.Background(), sql, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting user: %v\n", err)
		return
	}
	fmt.Println("User deleted successfully")
}

func main() {
	db := connectDB()
	defer db.Close()

	// Example operations
	createUser(db, "Alice", "alice@example.com")
	readUsers(db)
	updateUser(db, 1, "Alice Smith")
	readUsers(db)
	deleteUser(db, 1)
	readUsers(db)
}
