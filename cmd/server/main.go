package main

import (
	"github.com/PiotrIzw/webstore-grcp/config/database"
	"github.com/PiotrIzw/webstore-grcp/internal/middleware"
	"github.com/PiotrIzw/webstore-grcp/internal/middleware/authorizer"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
	"github.com/PiotrIzw/webstore-grcp/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	database.ConnectDB()

	db := database.DB

	rolesRepo := repository.NewRolesRepository(db)

	authorizerUtil := authorizer.NewAuthorizer(rolesRepo)
	rolesService := service.NewRolesService(rolesRepo, authorizerUtil)

	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo, authorizerUtil)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := middleware.AuthInterceptor()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor),
		grpc.StreamInterceptor(middleware.StreamAuthInterceptor()))
	pb.RegisterAccountServiceServer(grpcServer, accountService)

	preferencesRepo := repository.NewPreferencesRepository(db)
	preferencesService := service.NewPreferencesService(preferencesRepo, authorizerUtil)
	pb.RegisterPreferencesServiceServer(grpcServer, preferencesService)

	pb.RegisterRolesServiceServer(grpcServer, rolesService)

	ordersRepo := repository.NewOrdersRepository(db)
	ordersService := service.NewOrdersService(ordersRepo, authorizerUtil)
	pb.RegisterOrdersServiceServer(grpcServer, ordersService)

	fileRepo := repository.NewFileRepository(db)
	fileService := service.NewFileService(fileRepo, authorizerUtil)
	pb.RegisterFileServiceServer(grpcServer, fileService)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
