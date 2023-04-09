package services

import (
	"latihan-jwt/models"
	"latihan-jwt/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProductByID(produntID uint) (*models.Product, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.productRepository.CreateProduct(product)
}

func (s *productService) GetProductByID(produntID uint) (*models.Product, error) {
	return s.productRepository.GetProductByID(produntID)
}
