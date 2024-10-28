package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type HashService struct{}

// CheckPasswordHash implements Service.
func (h *HashService) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword implements Service.
func (h *HashService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func NewService() Service {
	return &HashService{}
}
