package repository

import "eko-car/infrastructure/database"

type HealthRepository interface {
	DatabaseHealth() (status bool, err error)
}

type healthRepository struct {
	database *database.Database
}

func NewHealthFeature(db *database.Database) HealthRepository {
	return &healthRepository{
		database: db,
	}
}
