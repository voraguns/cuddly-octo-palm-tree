package backlogimpl

import (
	"context"
	"cuddly-octo-palm-tree/domain"
)

type BiRepository interface {
	GetUsers(ctx context.Context) ([]*domain.Users, error)
	GetProducts(ctx context.Context) ([]*domain.Products, error)
	GetOrderItems(ctx context.Context) ([]*domain.OrderItems, error)
	GetOrders(ctx context.Context) ([]*domain.Orders, error)
}
