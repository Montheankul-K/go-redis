package main

import (
	"github.com/Montheankul-K/go-redis/handlers"
	"github.com/Montheankul-K/go-redis/repositories"
	"github.com/Montheankul-K/go-redis/services"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepository := repositories.NewProductRepositoryDB(db)
	// productRepository := repositories.NewProductRepositoryRedis(db, redisClient)
	productService := services.NewCatalogService(productRepository)
	// productService := services.NewCatalogServiceRedis(productRepository, redisClient)
	// productHandler := handlers.NewCatalogHandler(productService)
	productHandler := handlers.NewCatalogHandler(productService)

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Microsecond * 10)
		return c.SendString("Hello, World!")
	})
	app.Get("/products", productHandler.GetProducts)

	app.Listen(":8000")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/redis-demo")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // redis server
	})
}
