package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Montheankul-K/go-redis/repositories"
	"github.com/go-redis/redis/v8"
	"time"
)

type catalogServiceRedis struct {
	productRepository repositories.ProductRepository
	redisClient       *redis.Client
}

func NewCatalogServiceRedis(productRepository repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{
		productRepository: productRepository,
		redisClient:       redisClient,
	}
}

func (s catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "service::GetProducts"

	// redis get
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	// repository
	productDB, err := s.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, product := range productDB {
		products = append(products, Product{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
		})
	}

	// redis set
	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, data, time.Second*10)
	}

	fmt.Println("database")
	return products, nil
}
