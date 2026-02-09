package services

import (
	"gocashier.db/internal/models"
	"gocashier.db/internal/repository"
)

type categoryService struct {
	categoryRepo repository.CategoryRepo
}

type CategoryService interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	UpdateById(id int, Category *models.Category) error
	DeleteById(id int) error
	GetById(id int) (*models.Category, error)
}

func NewcategoryService(categoryRepo repository.CategoryRepo) CategoryService {
	return categoryService{categoryRepo: categoryRepo}
}

func (c categoryService) Create(category *models.Category) error {
	return c.categoryRepo.Create(category)
}

func (c categoryService) GetAll() ([]models.Category, error) {
	return c.categoryRepo.GetAll()
}

func (c categoryService) UpdateById(id int, Category *models.Category) error {
	return c.categoryRepo.UpdateById(id, Category)
}

func (c categoryService) DeleteById(id int) error {
	return c.categoryRepo.DeleteById(id)
}

func (c categoryService) GetById(id int) (*models.Category, error) {
	return c.categoryRepo.GetById(id)
}
