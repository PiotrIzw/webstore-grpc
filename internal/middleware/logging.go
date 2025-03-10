package middleware

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Received request: %v", req)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return resp, err
}
