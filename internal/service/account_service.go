package service

import (
	"context"
	"errors"
	"github.com/PiotrIzw/webstore-grcp/internal/account"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
	"github.com/PiotrIzw/webstore-grcp/pkg/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type AccountService struct {
	repo         repository.AccountRepository
	rolesService *RolesService
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(repo repository.AccountRepository, rolesService *RolesService) *AccountService {
	return &AccountService{
		repo:         repo,
		rolesService: rolesService, // Inject RolesService
	}
}

// CreateAccount creates a new account.
func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the account model
	acc := &account.Account{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
		Status:         "ACTIVE", // Default status
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Save the account to the database
	err = s.repo.CreateAccount(acc)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{AccountId: acc.ID}, nil
}

func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	// Authorize the request using the injected RolesService
	if err := Authorize(ctx, s.rolesService, "accounts:read"); err != nil {
		return nil, err
	}

	// Proceed with the method logic
	acc, err := s.repo.GetAccount(req.AccountId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get account: %v", err)
	}
	if acc == nil {
		return nil, status.Errorf(codes.NotFound, "account not found")
	}

	return &pb.GetAccountResponse{
		Username: acc.Username,
		Email:    acc.Email,
		Status:   acc.Status,
	}, nil
}

// UpdateAccount updates an existing account.
func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	// Retrieve the existing account
	acc, err := s.repo.GetAccount(req.AccountId)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, errors.New("account not found")
	}

	// Update the account fields
	acc.Username = req.Username
	acc.Email = req.Email
	//acc.Status = req.Status

	// Save the updated account
	err = s.repo.UpdateAccount(acc)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAccountResponse{AccountId: req.AccountId, Success: true}, nil
}

// DeleteAccount deletes an account by ID.
func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	err := s.repo.DeleteAccount(req.AccountId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccountResponse{Success: true}, nil
}

func (s *AccountService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Fetch the account from the database
	acc, err := s.repo.GetAccountByUsername(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch account: %v", err)
	}
	if acc == nil {
		return nil, status.Errorf(codes.NotFound, "account not found")
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(acc.HashedPassword), []byte(req.Password)); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	// Generate a JWT token
	token, err := auth.GenerateToken(acc.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	// Return the token
	return &pb.LoginResponse{Token: token}, nil
}
