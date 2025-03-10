package service

import (
	"context"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
)

type RolesService struct {
	repo *repository.RolesRepository
	pb.UnimplementedRolesServiceServer
}

func NewRolesService(repo *repository.RolesRepository) *RolesService {
	return &RolesService{repo: repo}
}

func (s *RolesService) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.AssignRoleResponse, error) {
	err := s.repo.AssignRole(req.UserId, req.RoleName)
	if err != nil {
		return nil, err
	}
	return &pb.AssignRoleResponse{Success: true}, nil
}

func (s *RolesService) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	allowed, err := s.repo.CheckPermission(req.UserId, req.Permission)
	if err != nil {
		return nil, err
	}
	return &pb.CheckPermissionResponse{Allowed: allowed}, nil
}
