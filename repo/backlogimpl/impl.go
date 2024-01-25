package backlogimpl

import "gorm.io/gorm"

type BiRepositoryImpl struct {
	db *gorm.DB
}

func NewBiRepository(db *gorm.DB) BiRepository {
	return &BiRepositoryImpl{
		db: db,
	}
}
