package main

import (
	"github.com/PiotrIzw/webstore-grcp/config/database"
	"github.com/PiotrIzw/webstore-grcp/internal/middleware"
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
	rolesService := service.NewRolesService(rolesRepo)

	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo, rolesService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := middleware.AuthInterceptor(rolesService)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	pb.RegisterAccountServiceServer(grpcServer, accountService)

	preferencesRepo := repository.NewPreferencesRepository(db)
	preferencesService := service.NewPreferencesService(preferencesRepo)
	pb.RegisterPreferencesServiceServer(grpcServer, preferencesService)

	pb.RegisterRolesServiceServer(grpcServer, rolesService)

	ordersRepo := repository.NewOrdersRepository(db)
	ordersService := service.NewOrdersService(ordersRepo)
	pb.RegisterOrdersServiceServer(grpcServer, ordersService)

	fileRepo := repository.NewFileRepository(db)
	fileService := service.NewFileService(fileRepo)
	pb.RegisterFileServiceServer(grpcServer, fileService)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
