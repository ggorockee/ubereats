package service

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	restaurantDto "ubereats/app/domain/restaurant/dto"
	restaurantRepo "ubereats/app/domain/restaurant/repository"
	restaurantRes "ubereats/app/domain/restaurant/response"

	"gorm.io/gorm"
)

type CategoryService interface {
	CreateCategory(inputParam *restaurantDto.CreateCategoryInput) (*restaurantRes.CreateCategoryOutput, error)
}

type categoryService struct {
	dbConn       *gorm.DB
	categoryRepo restaurantRepo.CategoryRepository
}

// CreateCategory implements CategoryService.
func (s *categoryService) CreateCategory(inputParam *restaurantDto.CreateCategoryInput) (*restaurantRes.CreateCategoryOutput, error) {
	var category *entity.Category
	err := s.dbConn.Transaction(func(tx *gorm.DB) error {
		var err error
		category, err = s.categoryRepo.CreateCategory(inputParam)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &restaurantRes.CreateCategoryOutput{
			CoreResponse: common.CoreResponse{
				Message: err.Error(),
			},
		}, err
	}

	return &restaurantRes.CreateCategoryOutput{
		CoreResponse: common.CoreResponse{
			Ok:   true,
			Data: category,
		},
	}, nil
}

func NewCategoryService(
	categoryRepo restaurantRepo.CategoryRepository,
	dbConn *gorm.DB,
) CategoryService {
	return &categoryService{
		dbConn:       dbConn,
		categoryRepo: categoryRepo,
	}
}
