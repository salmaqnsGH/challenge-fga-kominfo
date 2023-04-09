package services

import (
	"errors"
	"latihan-jwt/models"
	"latihan-jwt/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductServiceGetOneProduct(t *testing.T) {
	var productRepo = &repositories.MockProductRepository{Mock: mock.Mock{}}
	var productServiceMock = NewProductService(productRepo)

	expectedProduct := &models.Product{
		GORMModel: models.GORMModel{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:      1,
		Title:       "Produk 1",
		Description: "Deskripsi Produk 1",
		User:        &models.User{},
	}

	productRepo.Mock.On("FindByID", uint(1)).Return(expectedProduct, nil)

	resultProduct, err := productServiceMock.GetProductByID(uint(1))

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, resultProduct)
	assert.Equal(t, expectedProduct.GORMModel.ID, resultProduct.GORMModel.ID, "result has to be '1'")
	assert.Equal(t, expectedProduct.UserID, resultProduct.UserID, "result has to be '1'")
	assert.Equal(t, expectedProduct.Title, resultProduct.Title, "result has to be 'Produk 1'")
	assert.Equal(t, expectedProduct.Description, resultProduct.Description, "result has to be 'Deskripsi Produk 1'")

	productRepo.Mock.AssertExpectations(t)
}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	var productRepo = &repositories.MockProductRepository{Mock: mock.Mock{}}
	var productServiceMock = NewProductService(productRepo)

	productRepo.Mock.On("FindByID", uint(2)).Return(nil, errors.New("record not found"))

	resultProduct, err := productServiceMock.GetProductByID(uint(2))
	assert.Nil(t, resultProduct)
	assert.Error(t, err)
	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error(), "error response has to be 'record not found'[=]")

	productRepo.Mock.AssertExpectations(t)
}

func TestFindAllProductFound(t *testing.T) {
	var productRepo = &repositories.MockProductRepository{Mock: mock.Mock{}}
	var productServiceMock = NewProductService(productRepo)

	expectedProducts := []models.Product{
		{GORMModel: models.GORMModel{ID: 1}, Title: "Product 1", Description: "Product 1 description"},
		{GORMModel: models.GORMModel{ID: 2}, Title: "Product 2", Description: "Product 2 description"},
	}

	productRepo.Mock.On("FindAll").Return(expectedProducts, nil)

	products, err := productServiceMock.GetProducts()

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	productRepo.Mock.AssertExpectations(t)
}

func TestFindAllProductNotFound(t *testing.T) {
	var productRepo = &repositories.MockProductRepository{Mock: mock.Mock{}}
	var productServiceMock = NewProductService(productRepo)

	expectedError := errors.New("products not found")
	expectedProducts := []models.Product{}

	productRepo.Mock.On("FindAll").Return([]models.Product{}, expectedError)

	result, err := productServiceMock.GetProducts()

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.NotNil(t, err)
	assert.Equal(t, expectedProducts, result)

	productRepo.Mock.AssertExpectations(t)
}
