package auth

import (
	"errors"
	"time"
	"vnpay-demo/src/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID uint64, roles []model.Role, userFullName string, timeExp int) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ParseToken(tokenString string) (*Claims, error)
}

type JwtService struct {
	secretKey string
}

func NewService(secretKey string) Service {
	return &JwtService{secretKey: secretKey}
}

type Claims struct {
	UserID       uint64       `json:"user_id"`
	UserFullName string       `json:"user_full_name"`
	Role         []model.Role `json:"role"`
	jwt.RegisteredClaims
}

func (s *JwtService) GenerateToken(userID uint64, roles []model.Role, userFullName string, timeExp int) (string, error) {
	claims := &Claims{
		UserID:       userID,
		UserFullName: userFullName,
		Role:         roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(timeExp) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return token, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func (s *JwtService) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
