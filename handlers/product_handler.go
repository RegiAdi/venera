package handlers

import (
	"github.com/RegiAdi/venera/kernel"
	"github.com/RegiAdi/venera/models"
	"github.com/RegiAdi/venera/responses"
	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	GetProduct(id string) (models.Product, error)
	GetProducts() ([]models.Product, error)
	CreateProduct(product models.Product) (responses.ProductResponse, error)
	UpdateProduct(id string, product models.Product) error
	DeleteProduct(id string) error
}

type ProductHandler struct {
	productService ProductService
}

func NewProductHandler(productService ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (handler *ProductHandler) GetProductHandler(c *fiber.Ctx) error {
	product, err := handler.productService.GetProduct(c.Params("id"))
	if err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusNotFound,
			Status:     kernel.StatusFailed,
			Message:    "Product not found",
			Data:       nil,
		})
	}

	return responses.SendResponse(c, responses.BaseResponse{
		StatusCode: kernel.StatusOK,
		Status:     kernel.StatusSuccess,
		Message:    "Product retrieved successfully",
		Data:       product,
	})
}

func (handler *ProductHandler) GetProductsHandler(c *fiber.Ctx) error {
	products, err := handler.productService.GetProducts()
	if err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusNotFound,
			Status:     kernel.StatusFailed,
			Message:    "Products not found",
			Data:       nil,
		})
	}

	return responses.SendResponse(c, responses.BaseResponse{
		StatusCode: kernel.StatusOK,
		Status:     kernel.StatusSuccess,
		Message:    "Products retrieved successfully",
		Data:       products,
	})
}

func (handler *ProductHandler) CreateProductHandler(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusBadRequest,
			Status:     kernel.StatusFailed,
			Message:    "Failed to parse request",
			Data:       nil,
		})
	}

	productResponse, err := handler.productService.CreateProduct(product)
	if err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusBadRequest,
			Status:     kernel.StatusFailed,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return responses.SendResponse(c, responses.BaseResponse{
		StatusCode: kernel.StatusCreated,
		Status:     kernel.StatusSuccess,
		Message:    "Product created successfully",
		Data:       productResponse,
	})
}

func (handler *ProductHandler) UpdateProductHandler(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusBadRequest,
			Status:     kernel.StatusFailed,
			Message:    "Failed to parse request",
			Data:       nil,
		})
	}

	if err := handler.productService.UpdateProduct(c.Params("id"), product); err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusBadRequest,
			Status:     kernel.StatusFailed,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return responses.SendResponse(c, responses.BaseResponse{
		StatusCode: kernel.StatusOK,
		Status:     kernel.StatusSuccess,
		Message:    "Product updated successfully",
		Data:       nil,
	})
}

func (handler *ProductHandler) DeleteProductHandler(c *fiber.Ctx) error {
	if err := handler.productService.DeleteProduct(c.Params("id")); err != nil {
		return responses.SendResponse(c, responses.BaseResponse{
			StatusCode: kernel.StatusBadRequest,
			Status:     kernel.StatusFailed,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return responses.SendResponse(c, responses.BaseResponse{
		StatusCode: kernel.StatusOK,
		Status:     kernel.StatusSuccess,
		Message:    "Product deleted successfully",
		Data:       nil,
	})
}
