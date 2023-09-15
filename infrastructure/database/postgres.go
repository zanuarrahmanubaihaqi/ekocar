package database

import (
	"eko-car/infrastructure/shared/constant"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Dialect  string
	Host     string
	Name     string
	Username string
	Password string
	Port     string
}

type Database struct {
	*sqlx.DB
}

func LoadDatabase(config DatabaseConfig) (database *Database, err error) {

	datasource := fmt.Sprintf("%s:%s@(%s:3306)/%s?parseTime=true",
		// config.Dialect,
		config.Username,
		config.Password,
		config.Host,
		config.Name)
	db, err := sqlx.Connect(config.Dialect, datasource)
	if err != nil {
		err = fmt.Errorf(constant.ErrConnectToDB, err)
		return
	}

	database = &Database{
		db,
	}

	return
}
