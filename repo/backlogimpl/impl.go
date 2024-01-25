package backlogimpl

import (
	"context"
	"cuddly-octo-palm-tree/domain"

	"gorm.io/gorm"
)

type BiRepositoryImpl struct {
	db *gorm.DB
}

func NewBiRepository(db *gorm.DB) BiRepository {
	return &BiRepositoryImpl{
		db: db,
	}
}

func (r *BiRepositoryImpl) GetUsers(ctx context.Context) ([]*domain.Users, error) {
	var users []*domain.Users
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *BiRepositoryImpl) GetProducts(ctx context.Context) ([]*domain.Products, error) {
	var products []*domain.Products
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *BiRepositoryImpl) GetOrderItems(ctx context.Context) ([]*domain.OrderItems, error) {
	var orderItem []*domain.OrderItems
	if err := r.db.Find(&orderItem).Error; err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (r *BiRepositoryImpl) GetOrders(ctx context.Context) ([]*domain.Orders, error) {
	var orders []*domain.Orders
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
