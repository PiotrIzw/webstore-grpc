package authorizer

import (
	"context"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Authorizer struct {
	repo *repository.RolesRepository
}

func NewAuthorizer(repo *repository.RolesRepository) *Authorizer {
	return &Authorizer{repo: repo}
}

func (a *Authorizer) Authorize(ctx context.Context, permission string) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "user ID not found in context")
	}

	allowed, err := a.repo.CheckPermission(userID, permission)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to check permission: %v", err)
	}
	if !allowed {
		return status.Errorf(codes.PermissionDenied, "user does not have permission: %s", permission)
	}

	return nil
}
