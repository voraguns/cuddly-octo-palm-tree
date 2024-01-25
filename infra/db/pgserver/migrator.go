package pgserver

import (
	"cuddly-octo-palm-tree/infra/db/pgserver/models"

	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

// NewMigrator is the factory of Migrator
func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{
		db: db,
	}
}

// Migrate method migrates db schemas
func (m *Migrator) Migrate() error {

	return m.db.AutoMigrate(
		&models.Users{},
		&models.Orders{},
		&models.Products{},
	)
}
