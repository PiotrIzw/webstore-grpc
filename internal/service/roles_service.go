package service

import (
	"context"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func Authorize(ctx context.Context, rolesService *RolesService, permission string) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "user ID not found in context")
	}

	allowed, err := rolesService.repo.CheckPermission(userID, permission)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to check permission: %v", err)
	}
	if !allowed {
		return status.Errorf(codes.PermissionDenied, "user does not have permission: %s", permission)
	}

	return nil
}
