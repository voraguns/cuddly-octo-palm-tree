package backlogimpl

import (
	"context"
	"cuddly-octo-palm-tree/config"
	"cuddly-octo-palm-tree/domain"
	bi_repo "cuddly-octo-palm-tree/repo/backlogimpl"

	log "github.com/sirupsen/logrus"
)

type BiServiceImpl struct {
	biRepo bi_repo.BiRepository
	// producer *kafka.Producer
	logger *log.Entry
}

func NewBiService(conf *config.Config, repo bi_repo.BiRepository) BiService {
	return &BiServiceImpl{
		biRepo: repo,
		// producer: producer,
		logger: conf.Logger.ContextLogger.WithFields(log.Fields{
			"type": "service:Bi",
		}),
	}
}

func (s *BiServiceImpl) GetUsers(ctx context.Context) ([]*domain.Users, error) {
	users, err := s.biRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *BiServiceImpl) GetOrders(ctx context.Context) ([]*domain.Orders, error) {
	order, err := s.biRepo.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *BiServiceImpl) GetOrderItems(ctx context.Context) ([]*domain.OrderItems, error) {
	orderItem, err := s.biRepo.GetOrderItems(ctx)
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (s *BiServiceImpl) GetProducts(ctx context.Context) ([]*domain.Products, error) {
	products, err := s.biRepo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *BiServiceImpl) GetUsersById(ctx context.Context, id int) (domain.Users, error) {
	user, err := s.biRepo.GetUserById(ctx, id)
	if err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

func (s *BiServiceImpl) DeleteUser(ctx context.Context, id int) error {
	err := s.biRepo.DeleteUser(ctx, id)
	return err
}
