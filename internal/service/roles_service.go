package service

import (
	"context"
	"github.com/PiotrIzw/webstore-grcp/internal/middleware/authorizer"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
)

type RolesService struct {
	repo       *repository.RolesRepository
	authorizer *authorizer.Authorizer
	pb.UnimplementedRolesServiceServer
}

func NewRolesService(repo *repository.RolesRepository, authorizer *authorizer.Authorizer) *RolesService {
	return &RolesService{repo: repo, authorizer: authorizer}
}

func (s *RolesService) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.AssignRoleResponse, error) {

	if err := s.authorizer.Authorize(ctx, "roles:write"); err != nil {
		return nil, err
	}

	err := s.repo.AssignRole(req.UserId, req.RoleName)
	if err != nil {
		return nil, err
	}
	return &pb.AssignRoleResponse{Success: true}, nil
}
