package backlogimpl

import (
	"context"
	"cuddly-octo-palm-tree/domain"
)

type BiService interface {
	GetUsers(ctx context.Context) ([]*domain.Users, error)
	GetOrders(ctx context.Context) ([]*domain.Orders, error)
	GetOrderItems(ctx context.Context) ([]*domain.OrderItems, error)
	GetProducts(ctx context.Context) ([]*domain.Products, error)
}
