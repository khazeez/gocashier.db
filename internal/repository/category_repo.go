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
	query := `
		INSERT INTO category (category_name, description)
		VALUES ($1, $2)
		RETURNING id, category_name, description;
	`

	return p.db.QueryRow(
		query,
		category.Name,
		category.Description,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
}

func (p *categoryRepo) GetAll() ([]models.Category, error) {
	query := `SELECT id, category_name, description, created_at FROM category;`
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
			&category.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error select category: %w", err)
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil

}

func (p *categoryRepo) UpdateById(id int, category *models.Category) error {
	query := `
		UPDATE category
		SET category_name = $1,
		    description   = $2
		WHERE id = $3
		RETURNING id, category_name, description;
	`

	err := p.db.QueryRow(
		query,
		category.Name,
		category.Description,
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)

	if err == sql.ErrNoRows {
		return fmt.Errorf("category not found")
	}
	if err != nil {
		return fmt.Errorf("error update category: %w", err)
	}

	return nil
}

func (p *categoryRepo) DeleteById(id int) error {
	query := `DELETE FROM category WHERE id=$1;`
	result, err := p.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error delete category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (p *categoryRepo) GetById(id int) (*models.Category, error) {
	query := `SELECT id, category_name, description FROM category WHERE id=$1;`
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
