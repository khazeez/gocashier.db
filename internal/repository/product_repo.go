package repository

import (
	"database/sql"
	"fmt"

	"gocashier.db/internal/models"
)

type productRepo struct {
	db *sql.DB
}

type ProductRepo interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	UpdateById(id int, product *models.Product) error
	DeleteById(id int) error
	GetById(id int) (*models.Product, error)
	GetDetailProductById(id int) (*models.ProductDetail, error)
}

func NewProductRepository(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(product *models.Product) error {

	query := `INSERT INTO product (category_id, product_name, price, stock) VALUES ($1, $2, $3, $4) RETURNING category_id, product_name, price, stock, created_at;`

	err := p.db.QueryRow(
		query,
		product.CategoryId,
		product.Name,
		product.Price,
		product.Stock,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("error create product: %w", err)
	}
	return nil
}

func (p *productRepo) GetAll() ([]models.Product, error) {
	query := `SELECT id, category_id, product_name, price, stock, created_at FROM product;`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error select product: %w", err)
	}

	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.CategoryId,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error select product: %w", err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, err

}

func (p *productRepo) UpdateById(id int, product *models.Product) error {
	query := `
		UPDATE product
		SET category_id = $1,
		    product_name = $2,
		    price        = $3,
		    stock        = $4
		WHERE id = $5
		RETURNING id, category_id, product_name, price, stock, created_at;
	`

	err := p.db.QueryRow(
		query,
		product.CategoryId,
		product.Name,
		product.Price,
		product.Stock,
		id,
	).Scan(
		&product.ID,
		&product.CategoryId,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return fmt.Errorf("product not found")
	}
	if err != nil {
		return fmt.Errorf("error update product: %w", err)
	}

	return nil
}

func (p *productRepo) DeleteById(id int) error {
	query := `DELETE FROM product WHERE id=$1;`
	result, err := p.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error delete product: %w", err)
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

func (p *productRepo) GetById(id int) (*models.Product, error) {
	query := `SELECT id, category_id, product_name, price, stock, created_at FROM product WHERE id=$1;`
	row := p.db.QueryRow(query, id)
	var product models.Product
	err := row.Scan(
		&product.ID,
		&product.CategoryId,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("Error get product: %w", err)
	}

	return &product, nil
}

func (p *productRepo) GetDetailProductById(id int) (*models.ProductDetail, error) {
	query := `
SELECT
    p.id,
    p.product_name,
    p.price,
    p.stock,
    p.created_at,
    c.id,
    c.category_name,
    c.description,
    c.created_at
FROM product p
INNER JOIN category c ON p.category_id = c.id
WHERE p.id = $1
`

	row := p.db.QueryRow(query, id)

	var product models.ProductDetail

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.Category.ID,
		&product.Category.Name,
		&product.Category.Description,
		&product.Category.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error get detail product: %w", err)
	}

	return &product, nil
}
