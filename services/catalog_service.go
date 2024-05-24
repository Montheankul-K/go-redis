package services

import "github.com/Montheankul-K/go-redis/repositories"

type catalogService struct {
	productRepository repositories.ProductRepository
}

func NewCatalogService(productRepository repositories.ProductRepository) CatalogService {
	return catalogService{productRepository: productRepository}
}

func (s catalogService) GetProducts() (products []Product, err error) {
	productsDB, err := s.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, product := range productsDB {
		products = append(products, Product{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
		})
	}

	return products, nil
}
