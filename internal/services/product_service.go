package services

import (
	"gocashier.db/internal/models"
	"gocashier.db/internal/repository"
)

type productService struct {
	productRepo repository.ProductRepo
}

type ProductService interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	UpdateById(id int, product *models.Product) error
	DeleteById(id int) error
	GetById(id int) (*models.Product, error)
	GetDetailProductById(id int) (*models.ProductDetail, error)
}

func NewProductService(productRepo repository.ProductRepo) ProductService {
	return productService{
		productRepo: productRepo,
	}
}

func (p productService) Create(product *models.Product) error {
	return p.productRepo.Create(product)
}

func (p productService) GetAll() ([]models.Product, error) {
	return p.productRepo.GetAll()
}

func (p productService) UpdateById(id int, product *models.Product) error {
	return p.productRepo.UpdateById(id, product)
}

func (p productService) DeleteById(id int) error {
	return p.productRepo.DeleteById(id)
}

func (p productService) GetById(id int) (*models.Product, error) {
	return p.productRepo.GetById(id)
}

func (p productService) GetDetailProductById(id int) (*models.ProductDetail, error) {
	return p.productRepo.GetDetailProductById(id)
}
