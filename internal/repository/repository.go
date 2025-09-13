package repository

import (
	"fmt"

	"github.com/precioussilence/go-gin-starter/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	User UserRepository
}

func NewRepositories(cfg *config.Config) *Repositories {
	db := initDB(cfg)
	return &Repositories{
		User: NewUserRepository(db),
	}
}

func initDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to init db: " + err.Error())
	}
	return db
}
