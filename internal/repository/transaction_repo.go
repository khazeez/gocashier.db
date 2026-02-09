package repository

import (
	"database/sql"
	"fmt"

	"gocashier.db/internal/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT product_name, price, stock FROM product WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE product SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transaction (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, sub_total) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}


func (t *TransactionRepository) GetReportToday() (*models.TransactionReport, error) {

	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0),
			COUNT(id)
		FROM "transaction"
		WHERE created_at >= CURRENT_DATE
		  AND created_at < CURRENT_DATE + INTERVAL '1 day'
	`

	var report models.TransactionReport
	row := t.db.QueryRow(query)
	if err := row.Scan(&report.TotalRevenue, &report.TotalTransaction); err != nil {
		return nil, err
	}

	query = `
		SELECT 
			p.product_name,
			SUM(td.quantity) AS total_sold
		FROM transaction_details td
		JOIN "transaction" t ON t.id = td.transaction_id
		JOIN product p ON p.id = td.product_id
		WHERE t.created_at >= CURRENT_DATE
		  AND t.created_at < CURRENT_DATE + INTERVAL '1 day'
		GROUP BY p.id, p.product_name
		ORDER BY total_sold DESC
		LIMIT 1
	`

	row = t.db.QueryRow(query)
	err := row.Scan(
		&report.BestSellingProduct.Name,
		&report.BestSellingProduct.QuantitySelled,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &report, nil
}
