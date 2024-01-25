package backlogimpl

import (
	"context"
	"cuddly-octo-palm-tree/domain"
	"cuddly-octo-palm-tree/infra/db/pgserver/models"

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

func (r *BiRepositoryImpl) GetUserById(ctx context.Context, id int) (domain.Users, error) {
	var user domain.Users
	if err := r.db.First(&user, id).Error; err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

func (r *BiRepositoryImpl) GetOrderWithProductName(ctx context.Context, id int) (domain.OrderWithProduct, error) {
	var orderWithProductname domain.OrderWithProduct
	var order domain.OrderItems
	if err := r.db.First(&order).Joins("").Where("").Error; err != nil {
		return domain.OrderWithProduct{}, err
	}
	return orderWithProductname, nil
}

func (r *BiRepositoryImpl) DeleteUser(ctx context.Context, id int) error {
	var user models.Users

	if err := r.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *BiRepositoryImpl) EditUser(ctx context.Context, id int) error {
	var user models.Users
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
