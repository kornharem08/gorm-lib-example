package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kornharem08/app/internal/handler/userhandler"
	sqlwrap "github.com/kornharem08/gorm"
)

func main() {
	// Get database connection string from environment variable
	// dsn := os.Getenv("DB_CONNECTION_STRING")
	// if dsn == "" {
	// 	fmt.Println("Error: DB_CONNECTION_STRING is not set")
	// 	return
	// }
	// Connect to database
	conn, err := sqlwrap.New("sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=master&encrypt=true&trustServerCertificate=true")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the database!")

	handler := userhandler.NewHandler(conn)
	r := gin.Default()
	r.GET("/users", handler.Find)
	r.GET("/users/:employeeId", handler.FindByEmployeeId)
	r.Run() // listen and serve on 0.0.0.0:8080

	// // Start a new session (transaction)
	// session, err := conn.NewSession()
	// if err != nil {
	// 	log.Fatalf("Failed to create session: %v", err)
	// }

	// // Create a new user
	// user := User{
	// 	Name:      "John Doe",
	// 	Email:     "john@example.com",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	// if err := session.Create(ctx, &user); err != nil {
	// 	session.Rollback()
	// 	log.Fatalf("Failed to create user: %v", err)
	// }

	// fmt.Printf("User created: ID=%d, Name=%s\n", user.ID, user.Name)

	// // Query for users
	// var users []User
	// db := conn.Database()

	// if err := db.Table(&User{}).Find(ctx, &users); err != nil {
	// 	log.Fatalf("Failed to query users: %v", err)
	// }

	// fmt.Println("Users in database:")
	// for _, u := range users {
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	// }

	// // Commit the transaction
	// if err := session.Commit(); err != nil {
	// 	log.Fatalf("Failed to commit transaction: %v", err)
	// }

	// fmt.Println("Database operations completed successfully!")
}
