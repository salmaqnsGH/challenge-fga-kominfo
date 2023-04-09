package services

import (
	"errors"
	"fmt"
	"latihan-jwt/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repositories.ProductRepositoryMock{Mock: mock.Mock{}}
var productServiceMock = NewProductService(productRepo)

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	// Configure mock to return nil for product and an error for error
	productRepo.Mock.On("FindByID", uint(1)).Return(nil, errors.New("product not found"))

	product, err := productServiceMock.GetProductByID(uint(1))
	fmt.Println("product", product)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}
