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


func (t *TransactionRepository) GetReportToday() (error, *models.TransactionReport) {
  query := `SELECT SUM(total_amount) AS total_revenue, COUNT(id) AS total_transaction from transaction WHERE created_at >= CURRENT_DATE AND created_at < CURRENT_DATE + INTERVAL '1 day'`
  rows, err := t.db.Query(query)
  if err != nil {
	return fmt.Errorf(err.Error()), nil
  }

  var product models.TransactionReport
  err = rows.Scan(
	&product.TotalRevenue,
	&product.TotalTransaction,
  )


	query = `SELECT id from transaction WHERE created_at >= CURRENT_DATE AND created_at < CURRENT_DATE + INTERVAL '1 day'`
	rows, err = t.db.Query(query)
	if err != nil {
		return fmt.Errorf(err.Error()), nil
	}

	var items []int
	for rows.Next() {
		var item int
		err = rows.Scan(&item)
		if err != nil {
			return fmt.Errorf(err.Error()), nil
		}

		items = append(items, item)
	}

	freq:= make(map[int]int)
	for _, val := range items {
		freq[val]++
	}

	maxCount := 0
	mostFrequent := 0
	for key, count := range freq {
		if count > maxCount {
			maxCount = count
			mostFrequent = key
		}
	}

	query = `SELECT product_name from product WHERE id=$1`
	var product2 models.Product
	row := t.db.QueryRow(query, mostFrequent)
	err = row.Scan(
		&product2.Name,
	)

	fmt.Println("Elemen terbanyak:", mostFrequent, "dengan jumlah:", maxCount)


	return nil, &models.TransactionReport{
			TotalRevenue: product.TotalRevenue,
			TotalTransaction: product.TotalTransaction,
			BestSellingProduct: models.BestSellingProduct{Name: product2.Name, QuantitySelled: maxCount},
	}

}
