
package service

import (
	"ai_admin_project/internal/model"
	"ai_admin_project/internal/repository"
	"ai_admin_project/internal/request"
	"ai_admin_project/pkg/auth"
	"ai_admin_project/pkg/utils"
	"errors"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Register(req request.RegisterRequest) (*model.User, error) {
	// Check if user exists
	_, err := s.Repo.FindByUsername(req.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(req request.LoginRequest) (string, error) {
	user, err := s.Repo.FindByUsername(req.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid username or password")
	}

	return auth.GenerateToken(user.ID)
}
