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
	GetUsersById(ctx context.Context, id int) (domain.Users, error)
	DeleteUser(ctx context.Context, id int) error
	// EditUser(ctx context.Context, id int) error
}
