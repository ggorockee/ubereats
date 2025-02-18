package repository

import (
	"errors"
	"fmt"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"

	userDto "ubereats/app/domain/user/dto"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(context ...*fiber.Ctx) (*[]entity.User, error)
	CreateAccount(input *userDto.CreateAccount, context ...*fiber.Ctx) (*entity.User, error)

	GetFindById(id int, context ...*fiber.Ctx) (*entity.User, error)
	hashPassword(password string) (string, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// GetFindById implements UserRepository.
func (r *userRepository) GetFindById(id int, context ...*fiber.Ctx) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser implements UserRepository.
// func (r *userRepository) UpdateUser(input *userDto.UpdateUser, id int, context ...*fiber.Ctx) (*entity.User, error) {
// 	user, err := r.GetFindById(id, context...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := common.DecodeStructure(input, user); err != nil {
// 		return nil, err
// 	}

// 	if err := r.db.Save(user).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// CreateAccount implements UserRepository.
func (r *userRepository) CreateAccount(input *userDto.CreateAccount, context ...*fiber.Ctx) (*entity.User, error) {
	// 이메일로 기존 사용자 조회
	var existingUser entity.User
	err := r.db.Where("email = ?", input.Email).First(&existingUser).Error
	if err == nil {
		// 계정이 이미 존재함
		return nil, fmt.Errorf("user with email %s already exists", input.Email)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// DB 오류
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}

	// 계정이 존재하지 않음 -> 새 사용자 생성
	var user entity.User
	if err := common.DecodeStructure(input, &user); err != nil {
		return nil, fmt.Errorf("failed to decode input to user: %w", err)
	}

	// 비밀번호 해싱
	hashedPassword, err := r.hashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	// DB에 저장
	if err := r.db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

// GetAll implements UserRepository.
func (r *userRepository) GetAllUser(context ...*fiber.Ctx) (*[]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func NewUserRepository(d *gorm.DB) UserRepository {
	return &userRepository{db: d}
}
