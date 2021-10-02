package main

import (
	"fmt"
	"log"
	"os"
	"task/api/routes"
	"task/zapLog"
	"time"

	"task/package/auth"
	"task/package/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
)

func main() {

	//use godotenv if not using docker
	//godotenv.Load(".config")

	db, err := MysqlConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}

	usersRepository := users.NewRepoMySQL(db)
	usersService := users.NewService(usersRepository)
	authService := auth.NewService(usersRepository)

	// uploadService := upload.NewService()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome!"))
	})
	
	routes.UserRouter(app.Group("/api/v1/users"), usersService)
	routes.AuthRouter(app.Group("/api/v1/auth"), authService)
	// routes.UploadRouter(app.Group("/api/v1/upload"), uploadService)


	log.Fatal(app.Listen(":"+os.Getenv("APP_PORT")))
}

func MysqlConnection() (*sqlx.DB, error) {
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_NAME"))
	client, err := sqlx.Open("mysql", ds)
	if err != nil {
		zapLog.Error("error database : "+err.Error())
		return nil, err
	}
	
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 30)
	client.SetConnMaxIdleTime(time.Minute * 1)
	client.SetMaxOpenConns(100)
	client.SetMaxIdleConns(10)

	return client, nil

}
