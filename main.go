package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/virhanali/go-rest-api-fiber/book"
	"github.com/virhanali/go-rest-api-fiber/database"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx){
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("api/v1/book/:id", book.GetBooks)
	app.Post("api/v1/book", book.NewBooks)
	app.Delete("api/v1/book/:id", book.DeleteBooks)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database succesfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
	 
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	app.Get("/", helloWorld)

	setupRoutes(app)

	app.Listen(3000)
}