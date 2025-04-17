package repository

import "gorm.io/gorm"

type UserRepository interface{}

type userRepository struct {
	dbConn *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userRepository{
		dbConn: dbConn,
	}
}
