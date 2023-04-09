package services

import (
	"latihan-jwt/models"
	"latihan-jwt/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProductByID(produntID uint) (*models.Product, error)
	UpdateProduct(input *models.Product) (*models.Product, error)
	DeleteProduct(ID uint) error
	GetProducts() ([]models.Product, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.productRepository.Create(product)
}

func (s *productService) GetProductByID(produntID uint) (*models.Product, error) {
	return s.productRepository.FindByID(produntID)
}

func (s *productService) UpdateProduct(input *models.Product) (*models.Product, error) {
	product, err := s.productRepository.FindByID(input.ID)
	if err != nil {
		return product, err
	}

	product.Description = input.Description
	product.Title = input.Title

	updatedProduct, err := s.productRepository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *productService) DeleteProduct(ID uint) error {
	err := s.productRepository.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *productService) GetProducts() ([]models.Product, error) {
	photos, err := s.productRepository.FindAll()
	if err != nil {
		return photos, err
	}

	return photos, nil
}
