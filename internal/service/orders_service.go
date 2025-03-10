package service

import (
	"context"
	"github.com/PiotrIzw/webstore-grcp/internal/orders"
	"github.com/PiotrIzw/webstore-grcp/internal/pb"
	"github.com/PiotrIzw/webstore-grcp/internal/repository"
	"strconv"
)

type OrdersService struct {
	repo *repository.OrdersRepository
	pb.UnimplementedOrdersServiceServer
}

func NewOrdersService(repo *repository.OrdersRepository) *OrdersService {
	return &OrdersService{repo: repo}
}

func (s *OrdersService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var total float64
	for _, item := range req.Items {
		total += item.Price * float64(item.Quantity)
	}

	var orderItems []orders.OrderItem

	for _, item := range req.Items {
		total += item.Price * float64(item.Quantity)
		orderItems = append(orderItems, orders.OrderItem{
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity),
			Price:     item.Price,
		})
	}

	order := &orders.Order{
		UserID: req.UserId,
		Total:  total,
		Status: "PENDING",
		Items:  orderItems,
	}

	orderId, err := s.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{OrderId: strconv.Itoa(orderId), Total: total}, nil
}
