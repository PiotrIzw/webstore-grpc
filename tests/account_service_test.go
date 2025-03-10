package tests

import (
	"context"
	"errors"
	"github.com/PiotrIzw/webstore-grcp/internal/account"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockAccountRepository struct{}

func (m *MockAccountRepository) CreateAccount(acc *account.Account) error {
	acc.ID = "123"
	return nil
}

func (m *MockAccountRepository) GetAccount(id string) (*account.Account, error) {
	if id == "123" {
		return &account.Account{
			ID:       "123",
			Username: "testuser",
			Email:    "test@example.com",
			Status:   "ACTIVE",
		}, nil
	}
	return nil, errors.New("account not found")
}

func (m *MockAccountRepository) UpdateAccount(acc *account.Account) error {
	if acc.ID == "123" {
		return nil
	}
	return errors.New("account not found")
}

func (m *MockAccountRepository) DeleteAccount(id string) error {
	if id == "123" {
		return nil
	}
	return errors.New("account not found")
}

func TestCreateAccount(t *testing.T) {
	repo := &MockAccountRepository{}
	service := service.NewAccountService(repo)

	req := &pb.CreateAccountRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	res, err := service.CreateAccount(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "123", res.AccountId)
}
