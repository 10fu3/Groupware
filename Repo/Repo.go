package Repo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
import "../Entity"

type Repository struct {
	db *gorm.DB
}

func SetupDatabase(env Entity.Env) (*Repository, error) {

	config := Entity.DbSettings{
		Host:     env.DB_Host,
		Port:     env.DB_Port,
		User:     env.DB_User,
		Password: env.DB_Password,
		DbName:   env.DB_Name,
	}

	return DbConnect(config)
}

func DbConnect(settings Entity.DbSettings) (*Repository, error) {
	db, err := gorm.Open("mysql", settings.GetDbStrings())

	if err != nil {
		return nil, err
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// マイグレーション
	db.AutoMigrate(&Entity.User{})
	db.AutoMigrate(&Entity.Group{})

	return &Repository{db: db}, nil
}

func (r *Repository) doDB(do func() error) error {
	if r.db == nil {
		return fmt.Errorf("database not connected")
	}
	return do()
}
