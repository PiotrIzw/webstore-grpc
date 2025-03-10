package middleware

import (
	"context"
	"github.com/PiotrIzw/webstore-grcp/internal/service"
	"github.com/PiotrIzw/webstore-grcp/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a gRPC interceptor for authentication and authorization.
func AuthInterceptor(rolesService *service.RolesService) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		if info.FullMethod == "/account.AccountService/CreateAccount" || info.FullMethod == "/account.AccountService/Login" {
			return handler(ctx, req)
		}

		// Extract the JWT token from the metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		tokens := md.Get("authorization")
		if len(tokens) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		// Validate the token
		userID, err := auth.ValidateToken(tokens[0])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		// Add the user ID to the context
		ctx = context.WithValue(ctx, "user_id", userID)

		// Authorize the request (if required)
		//if err := service.Authorize(ctx, rolesService, "accounts:read"); err != nil {
		//	return nil, err
		//}

		// Call the handler
		return handler(ctx, req)
	}
}
