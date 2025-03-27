package repository

// import (
// 	"errors"
// 	"fmt"
// 	"ubereats/app/core/entity"
// 	"ubereats/app/core/helper/common"
// 	"ubereats/config"

// 	userDto "ubereats/app/domain/user/dto"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// type UserRepository interface {
// 	GetAllUser(context ...*fiber.Ctx) (*[]entity.User, error)
// 	CreateAccount(input *userDto.CreateAccount, context ...*fiber.Ctx) (*entity.User, error)

// 	GetFindById(id int, context ...*fiber.Ctx) (*entity.User, error)
// 	hashPassword(password string) (string, error)
// 	Login(input *userDto.Login, context ...*fiber.Ctx) (string, error)
// 	CheckPasswordHash(password, hash string) bool
// 	Me(c *fiber.Ctx) (*entity.User, error)
// }

// type userRepository struct {
// 	db  *gorm.DB
// 	cfg *config.Config
// }

// func (r *userRepository) Me(c *fiber.Ctx) (*entity.User, error) {
// 	if c == nil {
// 		return nil, errors.New("fiber context is required but nil provided")
// 	}

// 	// request_user 가져오기
// 	userRaw := c.Locals("request_user")
// 	if userRaw == nil {
// 		return nil, errors.New("no user found in request context")
// 	}

// 	// entity.User로 캐스팅
// 	user, ok := userRaw.(entity.User)
// 	if !ok {
// 		// anonymous 문자열인지 확인 (AuthMiddleware와 연계)
// 		if anon, isString := userRaw.(string); isString && anon == "anonymous" {
// 			return nil, errors.New("user is not authenticated")
// 		}
// 		return nil, errors.New("failed to cast request_user to entity.User")
// 	}

// 	return &user, nil
// }

// func (r *userRepository) CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func (r *userRepository) Login(input *userDto.Login, context ...*fiber.Ctx) (string, error) {
// 	email := input.Email
// 	password := input.Password

// 	var existingUser entity.User
// 	if err := r.db.Where("email = ?", email).First(&existingUser).Error; err != nil {
// 		return "", err
// 	}

// 	if !r.CheckPasswordHash(password, existingUser.Password) {
// 		return "", errors.New("password is incorrect")
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["user_id"] = existingUser.ID
// 	t, err := token.SignedString([]byte(r.cfg.Secret.Jwt))
// 	if err != nil {
// 		return "", err
// 	}

// 	if len(context) != 1 {
// 		return "", errors.New("fiber context not parsing")
// 	}

// 	c := context[0]
// 	c.Set("Authorization", "Bearer "+t)
// 	return t, nil

// }

// func (r *userRepository) hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// // GetFindById implements UserRepository.
// func (r *userRepository) GetFindById(id int, context ...*fiber.Ctx) (*entity.User, error) {
// 	var user entity.User
// 	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// // UpdateUser implements UserRepository.
// // func (r *userRepository) UpdateUser(input *userDto.UpdateUser, id int, context ...*fiber.Ctx) (*entity.User, error) {
// // 	user, err := r.GetFindById(id, context...)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	if err := common.DecodeStructure(input, user); err != nil {
// // 		return nil, err
// // 	}

// // 	if err := r.db.Save(user).Error; err != nil {
// // 		return nil, err
// // 	}
// // 	return user, nil
// // }

// // CreateAccount implements UserRepository.
// func (r *userRepository) CreateAccount(input *userDto.CreateAccount, context ...*fiber.Ctx) (*entity.User, error) {
// 	// 이메일로 기존 사용자 조회
// 	var existingUser entity.User
// 	err := r.db.Where("email = ?", input.Email).First(&existingUser).Error
// 	if err == nil {
// 		// 계정이 이미 존재함
// 		return nil, fmt.Errorf("user with email %s already exists", input.Email)
// 	}
// 	if !errors.Is(err, gorm.ErrRecordNotFound) {
// 		// DB 오류
// 		return nil, fmt.Errorf("failed to check user existence: %w", err)
// 	}

// 	// 계정이 존재하지 않음 -> 새 사용자 생성
// 	var user entity.User
// 	if err := common.DecodeStructure(input, &user); err != nil {
// 		return nil, fmt.Errorf("failed to decode input to user: %w", err)
// 	}

// 	// 비밀번호 해싱

// 	hashedPassword, err := r.hashPassword(input.Password)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to hash password: %w", err)
// 	}

// 	user.Password = hashedPassword

// 	// DB에 저장
// 	if err := r.db.Create(&user).Error; err != nil {
// 		return nil, fmt.Errorf("failed to create user: %w", err)
// 	}

// 	return &user, nil
// }

// // GetAll implements UserRepository.
// func (r *userRepository) GetAllUser(context ...*fiber.Ctx) (*[]entity.User, error) {
// 	var users []entity.User
// 	if err := r.db.Find(&users).Error; err != nil {
// 		return nil, err
// 	}
// 	return &users, nil
// }

// func NewUserRepository(d *gorm.DB, c *config.Config) UserRepository {
// 	return &userRepository{db: d, cfg: c}
// }
