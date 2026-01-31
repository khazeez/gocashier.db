package repository

import (
	"database/sql"
	"fmt"

	"gocashier.db/internal/models"
)

type categoryRepo struct {
	db *sql.DB
}

type CategoryRepo interface {
	Create(category *models.Category) error
	GetAll() ([]models.Category, error)
	UpdateById(id int, Category *models.Category) error
	DeleteById(id int) error
	GetById(id int) (*models.Category, error)
}

func NewcategoryRepository(db *sql.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (p *categoryRepo) Create(category *models.Category) error {

	query := `INSERT INTO category (id, category_name, description) VALUES ($1, $2, $3);`
	_, err := p.db.Exec(query, category.ID, category.Name, category.Description)
	if err != nil {
		return fmt.Errorf("error create category: %w", err)
	}
	return nil

}

func (p *categoryRepo) GetAll() ([]models.Category, error) {
	query := `SELECT * FROM category;`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error select category: %w", err)
	}

	defer rows.Close()

	categories := []models.Category{}
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("error select category: %w", err)
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, err

}

func (p *categoryRepo) UpdateById(id int, category *models.Category) error {
	query := `UPDATE (category_name, description) SET (category_name=$1, description=$2) WHERE id=$3`
	_, err := p.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return fmt.Errorf("Error update category: %w", err)
	}

	return nil
}

func (p *categoryRepo) DeleteById(id int) error {
	query := `DELETE FROM category WHERE id=$1;`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error delete category: %w", err)
	}

	return nil
}

func (p *categoryRepo) GetById(id int) (*models.Category, error) {
	query := `SELECT * FROM category WHERE id=$1;`
	row := p.db.QueryRow(query, id)
	var category models.Category
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)

	if err != nil {
		return nil, fmt.Errorf("Error get category: %w", err)
	}

	return &category, nil
}
