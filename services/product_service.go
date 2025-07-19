package services

import (
	"github.com/RegiAdi/venera/models"
	"github.com/RegiAdi/venera/responses"
)

type ProductRepository interface {
	GetProductByID(id string) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product models.Product) (responses.ProductResponse, error)
	UpdateProduct(id string, product models.Product) error
	DeleteProduct(id string) error
}

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (service *ProductService) GetProduct(id string) (models.Product, error) {
	return service.productRepository.GetProductByID(id)
}

func (service *ProductService) GetProducts() ([]models.Product, error) {
	return service.productRepository.GetAllProducts()
}

func (service *ProductService) CreateProduct(product models.Product) (responses.ProductResponse, error) {
	return service.productRepository.CreateProduct(product)
}

func (service *ProductService) UpdateProduct(id string, product models.Product) error {
	return service.productRepository.UpdateProduct(id, product)
}

func (service *ProductService) DeleteProduct(id string) error {
	return service.productRepository.DeleteProduct(id)
}
