package grpc

import (
	"context"
	client "yandexlms/pkg/api/order"
)

type OrderService struct {
	client.UnimplementedOrderServiceServer
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *client.CreateOrderRequest) (*client.CreateOrderResponse, error) {
	return &client.CreateOrderResponse{}, nil
}
