package repository

import (
	"database/sql"
	"github.com/PiotrIzw/webstore-grcp/internal/account"
	"time"
)

type AccountRepository interface {
	CreateAccount(acc *account.Account) error
	GetAccount(id string) (*account.Account, error)
	UpdateAccount(acc *account.Account) error
	DeleteAccount(id string) error
}

type AccountRepositoryImpl struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{db: db}
}

// CreateAccount inserts a new account into the database.
func (r *AccountRepositoryImpl) CreateAccount(acc *account.Account) error {
	query := `INSERT INTO accounts (username, email, hashed_password, status, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRow(query, acc.Username, acc.Email, acc.HashedPassword, acc.Status, time.Now(), time.Now()).Scan(&acc.ID)
}

// GetAccount retrieves an account by ID.
func (r *AccountRepositoryImpl) GetAccount(id string) (*account.Account, error) {
	var acc account.Account
	query := `SELECT id, username, email, status, created_at, updated_at FROM accounts WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&acc.ID, &acc.Username, &acc.Email, &acc.Status, &acc.CreatedAt, &acc.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Account not found
		}
		return nil, err
	}
	return &acc, nil
}

// UpdateAccount updates an existing account.
func (r *AccountRepositoryImpl) UpdateAccount(acc *account.Account) error {
	query := `UPDATE accounts
              SET username = $1, email = $2, status = $3, updated_at = $4
              WHERE id = $5`
	_, err := r.db.Exec(query, acc.Username, acc.Email, acc.Status, time.Now(), acc.ID)
	return err
}

// DeleteAccount deletes an account by ID.
func (r *AccountRepositoryImpl) DeleteAccount(id string) error {
	query := `DELETE FROM accounts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
