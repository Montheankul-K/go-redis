package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Montheankul-K/go-redis/services"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"time"
)

type catalogHandlerRedis struct {
	catalogService services.CatalogService
	redisClient    *redis.Client
}

func NewCatalogHandlerRedis(catalogService services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{catalogService: catalogService, redisClient: redisClient}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {
	key := "handler::GetProducts"

	// redis get
	if responseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}

	// service
	products, err := h.catalogService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// redis set
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("database")
	return c.JSON(response)
}
