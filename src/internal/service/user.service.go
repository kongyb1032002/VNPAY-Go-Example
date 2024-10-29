package service

import (
	"errors"
	"os"
	"time"
	"vnpay-demo/src/internal/model"
	"vnpay-demo/src/internal/repository"
	"vnpay-demo/src/internal/response"
	"vnpay-demo/src/pkg/auth"
	"vnpay-demo/src/pkg/hash"
)

type UserService interface {
	SignUp(user *model.User) error
	SignIn(username, password string) (*response.SignInResponse, error)
	ChangePassword(userID uint64, oldPassword, newPassword string) error

	List(conditions map[string]interface{}) (*[]model.User, error)
	Total(conditions map[string]interface{}) (uint64, error)
	GetByID(userID uint64) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint64) error
	RefreshToken(refreshToken string, userId uint64) (*response.SignInResponse, error)

	GetByIDs(ids []uint64) (*[]model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	hashService    hash.Service
	authService    auth.Service
}

// Total implements UserService.
func (s *userService) Total(conditions map[string]interface{}) (uint64, error) {
	return s.userRepository.Total(conditions)
}

// Delete implements UserService.
func (s *userService) Delete(id uint64) error {
	return s.userRepository.Delete(id)
}

func NewUserService(r repository.UserRepository, h hash.Service, a auth.Service) UserService {
	return &userService{
		userRepository: r,
		hashService:    h,
		authService:    a,
	}
}

func (s *userService) List(conditions map[string]interface{}) (*[]model.User, error) {
	return s.userRepository.List(conditions)
}

func (s *userService) SignUp(user *model.User) error {
	hashedPassword, err := s.hashService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.HashedPassword = hashedPassword
	return s.userRepository.Create(user)
}

func (s *userService) SignIn(username, password string) (*response.SignInResponse, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username")
	}
	if !s.hashService.CheckPasswordHash(password, user.HashedPassword) {
		user.AccessFailedCount += 1
		return nil, errors.New("invalid password")
	}
	token, err := s.authService.GenerateToken(user.ID, user.Roles, username, 24)
	if err != nil {
		return nil, err
	}
	// update key secret refresh token
	refreshToken := keySecretRefreshToken(token)
	user.RefreshToken = refreshToken
	user.LockoutEnd = time.Now()
	s.Update(user)

	return &response.SignInResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) RefreshToken(refreshToken string, userId uint64) (*response.SignInResponse, error) {
	user, err := s.userRepository.FindByID(userId)
	if err != nil {
		return nil, errors.New("invalid user")
	}
	if user.RefreshToken != refreshToken {
		return nil, errors.New("invalid refreshToken")
	}

	newToken, err := s.authService.GenerateToken(user.ID, user.Roles, user.Username, 168)
	if err != nil {
		return nil, err
	}

	user.RefreshToken = keySecretRefreshToken(newToken)
	s.Update(user)

	return &response.SignInResponse{
		Token:        newToken,
		RefreshToken: user.RefreshToken,
	}, nil
}

// GetByID implements UserService.
func (s *userService) GetByID(userID uint64) (*model.User, error) {
	return s.userRepository.FindByID(userID)
}

func (s *userService) Update(user *model.User) error {
	// Optionally hash the password if it is updated
	if user.Password != "" {
		hashedPassword, err := s.hashService.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.HashedPassword = hashedPassword
	}
	return s.userRepository.Update(user)
}

func (s *userService) ChangePassword(userID uint64, oldPassword, newPassword string) error {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return err
	}
	if !s.hashService.CheckPasswordHash(oldPassword, user.HashedPassword) {
		return errors.New("incorrect old password")
	}
	hashedPassword, err := s.hashService.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user.HashedPassword = hashedPassword
	return s.userRepository.Update(user)
}

func (s *userService) GetByIDs(ids []uint64) (*[]model.User, error) {
	return s.userRepository.GetByIDs(ids)
}

func keySecretRefreshToken(token string) string {
	refreshToken := os.Getenv("KEY_SECRET") + token
	return refreshToken
}
