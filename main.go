package main

import (
	"os"

	"github.com/Adonis2115/trend-backtest/database"
	"github.com/Adonis2115/trend-backtest/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()
	app := generateApp()
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	stocks := app.Group("/stocks")
	stocks.Get("/", handlers.Stocks)
	return app
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}
	err = database.InitDB()
	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
