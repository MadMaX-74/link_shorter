package auth

import (
	"errors"
	"go_dev/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, name, password string) (string, error) {
	existsUser, _ := service.UserRepository.FindByEmail(email)
	if existsUser != nil {
		return "", errors.New(UserAlreadyExists)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Name:     name,
		Email:    email,
		Password: string(hashPassword),
	}
	createdUser, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return createdUser.Email, nil
}
func (service *AuthService) Login(email, password string) (string, error) {
	existsUser, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if existsUser == nil {
		return "", errors.New(ErrUserNotFound)
	}
	err = bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(password))
	if err != nil {
		return "", err
	}
	return existsUser.Email, nil
}
