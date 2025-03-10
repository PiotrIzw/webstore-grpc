package repository

import (
	"database/sql"
	"github.com/PiotrIzw/webstore-grcp/internal/orders"
)

type OrdersRepository struct {
	db *sql.DB
}

func NewOrdersRepository(db *sql.DB) *OrdersRepository {
	return &OrdersRepository{db: db}
}

func (r *OrdersRepository) CreateOrder(order *orders.Order) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var orderID int
	query := `INSERT INTO orders (user_id, total, status) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(query, order.UserID, order.Total, order.Status).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, item := range order.Items {
		query := `INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`
		_, err := tx.Exec(query, orderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return orderID, tx.Commit()
}
