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
	UpdateById(product *models.Product) error
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

	query := `INSERT INTO product (id, category_id, product_name, price, stock) VALUES ($1, $2, $3, $4, $5);`
	_, err := p.db.Exec(query, product.ID, product.CategoryId, product.Name, product.Price, product.Stock)
	if err != nil {
		return fmt.Errorf("error create product: %w", err)
	}
	return nil

}

func (p *productRepo) GetAll() ([]models.Product, error) {
	query := `SELECT * FROM product;`
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

func (p *productRepo) UpdateById(product *models.Product) error {
	query := `UPDATE (category_id, product_name, price, stock) SET (category_id=$1, product_name=$2, price=$3, stock=$4) WHERE id=$5`
	_, err := p.db.Exec(query, product.CategoryId, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return fmt.Errorf("Error update product: %w", err)
	}

	return nil
}

func (p *productRepo) DeleteById(id int) error {
	query := `DELETE FROM product WHERE id=$1;`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error delete product: %w", err)
	}

	return nil
}

func (p *productRepo) GetById(id int) (*models.Product, error) {
	query := `SELECT * FROM product WHERE id=$1;`
	row := p.db.QueryRow(query, id)
	var product models.Product
	err := row.Scan(
		&product.ID,
		&product.CategoryId,
		&product.Name,
		&product.Price,
		&product.Stock,
	)

	if err != nil {
		return nil, fmt.Errorf("Error get product: %w", err)
	}

	return &product, nil
}

func (p *productRepo) GetDetailProductById(id int) (*models.ProductDetail, error) {
	query := `SELECT p.id, p.product_name, p.price, p.stock, p.created_at c.category_name, c.description FROM product p INNER JOIN category c ON p.category_id = c.id WHERE p.id=$1`
	row := p.db.QueryRow(query, id)
	var product models.ProductDetail
	err := row.Scan(
&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.Category.Name,
		&product.Category.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("Error get detail product: %w", err)
	}

	return  &product, nil

}
