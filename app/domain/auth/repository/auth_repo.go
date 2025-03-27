package repository

import (
	"fmt"
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
	authDto "ubereats/app/domain/auth/dto"
	"ubereats/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(c *fiber.Ctx, inputParam *authDto.LoginInput) (string, error)
	SignUp(inputParam *authDto.SignUpInput) (*entity.User, error)
	FindByOne(key, value string) (*entity.User, error)
	hashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}

type authRepository struct {
	dbConn *gorm.DB
	cfg    *config.Config
}

// hashPassword implements AuthRepository.
func (r *authRepository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// FindByOne implements AuthRepository.
func (r *authRepository) FindByOne(key string, value string) (*entity.User, error) {
	var output entity.User

	query := fmt.Sprintf("%s = ?", key)
	err := r.dbConn.Where(query, value).First(&output).Error
	if err != nil {
		return nil, err
	}

	return &output, nil
}

// Login implements AuthRepository.
func (r *authRepository) Login(c *fiber.Ctx, inputParam *authDto.LoginInput) (string, error) {
	email := inputParam.Email
	password := inputParam.Password

	existingUser, _ := r.FindByOne("email", email)
	if existingUser == nil {
		return "", fmt.Errorf("user가 존재하지 않습니당")
	}

	if !r.CheckPasswordHash(password, existingUser.Password) {
		return "", fmt.Errorf("패스워드가 틀렸어용")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = existingUser.ID
	t, err := token.SignedString([]byte(r.cfg.Secret.Jwt))
	if err != nil {
		return "", fmt.Errorf("token signedString 에러 %w", err)
	}

	c.Set("Authorization", "Bearer "+t)
	return t, nil

}

func (r *authRepository) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// SignUp implements AuthRepository.
func (r *authRepository) SignUp(inputParam *authDto.SignUpInput) (*entity.User, error) {
	email := inputParam.Email
	password := inputParam.Password
	role := inputParam.Role

	// 회원이 있는지 확인
	checkUser, _ := r.FindByOne("email", email)
	if checkUser != nil {
		return nil, fmt.Errorf("해당 %s은 존재합니당", email)
	}

	hashedPassword, err := r.hashPassword(password)
	if err != nil {
		return nil, err
	}

	createUser := entity.User{
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	if err := common.ValidateStruct(&createUser); err != nil {
		return nil, err
	}

	if err := r.dbConn.Create(&createUser).Error; err != nil {
		return nil, fmt.Errorf("계정 생~성 실패!! %w", err)
	}

	return &createUser, nil

}

func NewAuthRepository(
	dbConn *gorm.DB,
	cfg *config.Config,
) AuthRepository {
	return &authRepository{
		dbConn: dbConn,
		cfg:    cfg,
	}
}
