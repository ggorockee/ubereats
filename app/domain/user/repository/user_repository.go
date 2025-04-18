package repository

import (
	"fmt"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	userDto "ubereats/app/domain/user/dto"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateAccount(c *fiber.Ctx, inputParam *userDto.CreateAccountIn) (*entity.User, error)
	FineOne(key, value string) (*entity.User, error)
	hashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type userRepository struct {
	dbConn *gorm.DB
}

// CheckPasswordHash implements UserRepository.
func (r *userRepository) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

// hashPassword implements UserRepository.
func (r *userRepository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err

}

// CreateAccount implements UserRepository.
func (r *userRepository) CreateAccount(c *fiber.Ctx, inputParam *userDto.CreateAccountIn) (*entity.User, error) {
	// email check
	email := inputParam.Email
	password := inputParam.Password

	_, err := r.FineOne("email", email)

	if err == nil {
		return nil, fmt.Errorf("이미 가입된 이메일 주소임 ㅋㅋ")
	}

	hashedPassword, err := r.hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	var user entity.User
	if err := common.DecodeStructure(inputParam, &user); err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	if err := common.ValidateStruct(&user); err != nil {
		return nil, err
	}

	err = r.dbConn.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

// FineOne implements UserRepository.
func (r *userRepository) FineOne(key, value string) (*entity.User, error) {
	var obj entity.User
	query := fmt.Sprintf("%s = ?", key)
	err := r.dbConn.Where(query, value).First(&obj).Error
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userRepository{
		dbConn: dbConn,
	}
}
