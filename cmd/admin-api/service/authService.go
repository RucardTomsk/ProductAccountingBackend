package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"productAccounting-v1/cmd/admin-api/api/model"
	"productAccounting-v1/cmd/admin-api/storage/dao"
	"productAccounting-v1/internal/domain/base"
	"productAccounting-v1/internal/domain/entity"
	"productAccounting-v1/internal/domain/enum"

	"time"
)

const (
	salt       = "nsfgnstg45s5fbnsfdg"
	signingKey = "qwerqwerGS#jjsS"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserGuid string `json:"userGuid"`
	UserRole string `json:"userRole"`
}

type AuthService struct {
	storage *dao.AuthStorage
}

func NewAuthService(
	storage *dao.AuthStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) Register(request *model.RegisterRequest) (*uuid.UUID, *base.ServiceError) {
	user := entity.User{
		Email:    request.Email,
		Password: encryptString(request.Password),
		Role:     enum.ParseRoles(request.Role),
	}

	if err := s.storage.CreateUser(&user); err != nil {
		return nil, base.NewNeo4jWriteError(err)
	}

	return &user.ID, nil
}

func (s *AuthService) Login(request *model.AuthRequest) (*model.Token, *base.ServiceError) {
	user, err := s.storage.GetUser(request.Email, encryptString(request.Password))
	if err != nil {
		return nil, base.NewNeo4jReadError(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID.String(),
		user.Role.String(),
	})

	valueToken, err := token.SignedString([]byte(signingKey))
	return &model.Token{Value: valueToken}, nil
}

func encryptString(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
