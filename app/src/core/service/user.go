package service

import (
	"context"
	"github.com/alexeymyakinin/ruck/app/src/api/http/schema"
	"github.com/alexeymyakinin/ruck/app/src/core/model"
	"github.com/alexeymyakinin/ruck/app/src/core/repo"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	log  echo.Logger
	repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository, log echo.Logger) *UserService {
	return &UserService{
		log:  log,
		repo: repo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *schema.UserCreateRequest) (*model.User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		us.log.Error(err)
		return nil, err
	}

	res, err := us.repo.Insert(ctx, &model.User{Email: user.Email, Username: user.Username, Password: string(pwd)})
	if err != nil {
		us.log.Error(err)
		return nil, err
	}

	return res, nil
}

func (us *UserService) GetUser(ctx context.Context, userID uint64) (*model.User, error) {
	res, err := us.repo.SelectWhereId(ctx, userID)
	if err != nil {
		us.log.Error(err)
		return nil, err
	}

	return res, nil
}

func (us *UserService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	res, err := us.repo.SelectWhereUsername(ctx, username)
	if err != nil {
		us.log.Error(err)
		return nil, err
	}
	return res, nil
}
