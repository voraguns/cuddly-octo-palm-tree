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
	GetUserById(ctx context.Context, id int) (domain.Users, error)
	GetOrderWithProductName(ctx context.Context, id int) (domain.OrderWithProduct, error)
	DeleteUser(ctx context.Context, id int) error
	EditUser(ctx context.Context, id int) error
}
