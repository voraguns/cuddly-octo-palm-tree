package pgserver

import (
	"cuddly-octo-palm-tree/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseConnection returns the db connection instance
func NewDatabaseConnection(conf *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conf.DBConfig.Dsn), &gorm.Config{
		Logger:      conf.Logger.DBLogger,
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(conf.DBConfig.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(conf.DBConfig.MaxOpenConns)
	conf.Logger.ContextLogger.WithField("type", "setup:db").Info("successful SQL connection")
	return db, nil
}
